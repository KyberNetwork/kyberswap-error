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
}

type restTransformFunc func(rootCause error, entities ...string) *errors.RestAPIError

type restTransformer struct {
	mapping map[string]restTransformFunc
}

var restTransformerInstance *restTransformer

func initRestTransformerInstance() {
	if restTransformerInstance == nil {
		restTransformerInstance = &restTransformer{}
		restTransformerInstance.mapping = map[string]restTransformFunc{}
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeRequired, errors.NewRestAPIErrRequired)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeNotAcceptedValue, errors.NewRestAPIErrNotAcceptedValue)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeOutOfRange, errors.NewRestAPIErrOutOfRange)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeInvalidFormat, errors.NewRestAPIErrInvalidFormat)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeUnauthenticated, errors.NewRestAPIErrUnauthenticated)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeNotFound, errors.NewRestAPIErrNotFound)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeDuplicate, errors.NewRestAPIErrDuplicate)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeAlreadyExists, errors.NewRestAPIErrAlreadyExits)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeUnknown, errors.NewRestAPIErrInternal)
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
		return apiErrForTag(validationErr.Tag(), err, validationErr.Field())
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

// apiErrForTag return RestAPIError which corresponds to the validation tag
func apiErrForTag(tag string, err error, fields ...string) *errors.RestAPIError {
	switch tag {
	case "required":
		return errors.NewRestAPIErrRequired(err, fields...)
	case "oneof":
		return errors.NewRestAPIErrNotAcceptedValue(err, fields...)
	case "min", "max":
		return errors.NewRestAPIErrOutOfRange(err, fields...)
	default:
		return errors.NewRestAPIErrInternal(err)
	}
}
