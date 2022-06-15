package transformers

import (
	"fmt"

	c "github.com/KyberNetwork/kyberswap-error/pkg/constants"
	"github.com/KyberNetwork/kyberswap-error/pkg/errors"
)

type IRestTransformer interface {
	DomainErrToRestAPIErr(domainErr *errors.DomainError) *errors.RestAPIError
	RegisterTransformFunc(domainErrCode int, transformFunc restTransformFunc)
}

type restTransformFunc func(rootCause error, entities ...string) *errors.RestAPIError

type restTransformer struct {
	mapping map[int]restTransformFunc
}

var restTransformerInstance *restTransformer

func InitRestTransformerInstance() {
	if restTransformerInstance == nil {
		restTransformerInstance = &restTransformer{}
		restTransformerInstance.mapping = map[int]restTransformFunc{}
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeRequired, errors.NewRestAPIErrRequired)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeNotAcceptedValue, errors.NewRestAPIErrNotAcceptedValue)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeOutOfRange, errors.NewRestAPIErrOutOfRange)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeInvalidFormat, errors.NewRestAPIErrInvalidFormat)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeUnauthenticated, errors.NewRestAPIErrUnauthenticated)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeNotFound, errors.NewRestAPIErrNotFound)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeDuplicate, errors.NewRestAPIErrDuplicate)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeAlreadyExists, errors.NewRestAPIErrAlreadyExits)
		restTransformerInstance.RegisterTransformFunc(c.DomainErrCodeInternal, errors.NewRestAPIErrInternal)
	}
}

func RestTransformerInstance() IRestTransformer {
	return restTransformerInstance
}

// DomainErrToRestAPIErr transforms DomainError to RestAPIError
func (t *restTransformer) DomainErrToRestAPIErr(domainErr *errors.DomainError) *errors.RestAPIError {
	f := t.mapping[domainErr.Code]
	if f == nil {
		return errors.NewRestAPIErrInternal(fmt.Errorf("can not transform error, DomainError: %v", domainErr))
	}
	return f(domainErr.RootCause, domainErr.ErrorEntities...)
}

// RegisterTransformFunc is used to add new function to transform DomainError to RestAPIError
// if the domainErrCode is already registered, the old transform function will be overridden
func (t *restTransformer) RegisterTransformFunc(domainErrCode int, function restTransformFunc) {
	t.mapping[domainErrCode] = function
}
