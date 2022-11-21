package transformers

import (
	"encoding/json"
	errs "errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"

	c "github.com/KyberNetwork/kyberswap-error/pkg/constants"
	"github.com/KyberNetwork/kyberswap-error/pkg/errors"
)

type IRestTransformer interface {
	DomainErrToRestAPIErr(domainErr *errors.DomainError) *errors.RestAPIError
	ValidationErrToRestAPIErr(err error) *errors.RestAPIError
	RegisterTransformFunc(domainErrCode string, transformFunc restTransformFunc)
	RegisterValidationTag(tag string, function restTransformFunc)
}

type restTransformFunc func(rootCause error, entities ...string) *errors.RestAPIError

type restTransformer struct {
	mapping       map[string]restTransformFunc
	validationErr map[string]restTransformFunc
}

var restTransformerInstance *restTransformer

func initRestTransformerInstance() {
	if restTransformerInstance == nil {
		restTransformerInstance = &restTransformer{
			mapping:       make(map[string]restTransformFunc),
			validationErr: make(map[string]restTransformFunc),
		}

		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeRequired, errors.NewRestAPIErrRequired)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeNotAcceptedValue, errors.NewRestAPIErrNotAcceptedValue)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeOutOfRange, errors.NewRestAPIErrOutOfRange)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeInvalidFormat, errors.NewRestAPIErrInvalidFormat)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeInvalid, errors.NewRestAPIErrInvalid)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeUnauthenticated, errors.NewRestAPIErrUnauthenticated)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeUnauthorized, errors.NewRestAPIErrUnauthorized)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeNotFound, errors.NewRestAPIErrNotFound)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeDuplicate, errors.NewRestAPIErrDuplicate)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeAlreadyExists, errors.NewRestAPIErrAlreadyExits)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeUnknown, errors.NewRestAPIErrInternal)

		restTransformerInstance.RegisterValidationTag("required", errors.NewRestAPIErrRequired)
		restTransformerInstance.RegisterValidationTag("oneof", errors.NewRestAPIErrNotAcceptedValue)
		restTransformerInstance.RegisterValidationTag("min", errors.NewRestAPIErrOutOfRange)
		restTransformerInstance.RegisterValidationTag("max", errors.NewRestAPIErrOutOfRange)
		restTransformerInstance.RegisterValidationTag("numeric", errors.NewRestAPIErrInvalidFormat)
		restTransformerInstance.RegisterValidationTag("unique", errors.NewRestAPIErrDuplicate)
		restTransformerInstance.RegisterValidationTag("hexadecimal", errors.NewRestAPIErrInvalidFormat)
	}
}

func RestTransformerInstance() IRestTransformer {
	return restTransformerInstance
}

// ValidationErrToRestAPIErr transforms ValidationError to RestAPIError
// this function will be used when bind JSON request to DTO in gin framework
func (t *restTransformer) ValidationErrToRestAPIErr(err error) *errors.RestAPIError {
	var validationErrs validator.ValidationErrors
	var unmarshalTypeErr *json.UnmarshalTypeError
	var jsonSynTaxErr *json.SyntaxError
	var numErr *strconv.NumError
	if errs.As(err, &validationErrs) {
		validationErr := validationErrs[0]
		return t.apiErrForTag(validationErr.Tag(), err, validationErr.Field())
	}
	if errs.As(err, &unmarshalTypeErr) {
		field := unmarshalTypeErr.Field
		fieldArr := strings.Split(field, ".")
		return errors.NewRestAPIErrInvalidFormat(err, fieldArr[len(fieldArr)-1])
	}
	if errs.As(err, &jsonSynTaxErr) {
		return errors.NewRestAPIErrInvalidFormat(err)
	}
	if errs.As(err, &numErr) {
		return errors.NewRestAPIErrInvalidFormat(err)
	}
	return errors.NewRestAPIErrInternal(err)
}

// DomainErrToRestAPIErr transforms DomainError to RestAPIError
func (t *restTransformer) DomainErrToRestAPIErr(domainErr *errors.DomainError) *errors.RestAPIError {
	f := t.mapping[domainErr.Code]
	if f == nil {
		return errors.NewRestAPIErrInternal(fmt.Errorf("can not transform error, DomainError: %v", domainErr))
	}
	return f(domainErr, domainErr.ErrorEntities...)
}

// RegisterTransformFunc is used to add new function to transform DomainError to RestAPIError
// if the domainErrCode is already registered, the old transform function will be overridden
func (t *restTransformer) RegisterTransformFunc(domainErrCode string, function restTransformFunc) {
	t.mapping[domainErrCode] = function
}

// RegisterValidationTag is used to define new validation tag and respective API error
// if the validation tag is already registered, the old respective API error will be overridden
func (t *restTransformer) RegisterValidationTag(tag string, function restTransformFunc) {
	t.validationErr[tag] = function
}

// apiErrForTag return RestAPIError which corresponds to the validation tag
func (t *restTransformer) apiErrForTag(tag string, err error, fields ...string) *errors.RestAPIError {
	f := t.validationErr[tag]
	if f == nil {
		return errors.NewRestAPIErrInternal(err)
	}
	return f(err, fields...)
}
