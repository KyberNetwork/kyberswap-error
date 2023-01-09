package transformers

import (
	"fmt"

	c "github.com/KyberNetwork/kyberswap-error/pkg/constants"
	"github.com/KyberNetwork/kyberswap-error/pkg/errors"
)

type IInfraTransformer interface {
	RestAPIErrorToInfraErr(restAPIError *errors.RestAPIError) *errors.InfraError
	RegisterTransformFunc(restAPIErrCode int, transformFunc infraTransformFunc)
}

type infraTransformFunc func(rootCause error, entities ...string) *errors.InfraError

type infraTransformer struct {
	mapping map[int]infraTransformFunc
}

var infraTransformerInstance *infraTransformer

func initInfraTransformerInstance() {
	if infraTransformerInstance == nil {
		infraTransformerInstance = &infraTransformer{}
		infraTransformerInstance.mapping = map[int]infraTransformFunc{}
		infraTransformerInstance.RegisterTransformFunc(c.ClientErrCodeRequired, errors.NewInfraErrorHTTPRequired)
		infraTransformerInstance.RegisterTransformFunc(c.ClientErrCodeNotAcceptedValue, errors.NewInfraErrorHTTPNotAcceptedValue)
		infraTransformerInstance.RegisterTransformFunc(c.ClientErrCodeOutOfRange, errors.NewInfraErrorHTTPOutOfRange)
		infraTransformerInstance.RegisterTransformFunc(c.ClientErrCodeInvalidFormat, errors.NewInfraErrorHTTPInvalidFormat)
		infraTransformerInstance.RegisterTransformFunc(c.ClientErrCodeInvalid, errors.NewInfraErrorHTTPInvalid)
		infraTransformerInstance.RegisterTransformFunc(c.ClientErrCodeUnauthenticated, errors.NewInfraErrorHTTPUnauthenticated)
		infraTransformerInstance.RegisterTransformFunc(c.ClientErrCodeUnauthorized, errors.NewInfraErrorHTTPUnauthorized)
		infraTransformerInstance.RegisterTransformFunc(c.ClientErrCodeNotFound, errors.NewInfraErrorHTTPNotFound)
		infraTransformerInstance.RegisterTransformFunc(c.ClientErrCodeDuplicate, errors.NewInfraErrorHTTPDuplicate)
		infraTransformerInstance.RegisterTransformFunc(c.ClientErrCodeAlreadyExists, errors.NewInfraErrorHTTPAlreadyExists)
		infraTransformerInstance.RegisterTransformFunc(c.ClientErrCodeInternal, errors.NewInfraErrorHTTPUnknown)
	}
}

func InfraTransformerInstance() IInfraTransformer {
	return infraTransformerInstance
}

// RestAPIErrorToInfraErr transforms RestAPIError to InfraError
func (t *infraTransformer) RestAPIErrorToInfraErr(restAPIError *errors.RestAPIError) *errors.InfraError {
	f := t.mapping[restAPIError.Code]
	if f == nil {
		return errors.NewInfraErrorHTTPUnknown(fmt.Errorf("can not transform error, RestAPIError: %v", restAPIError))
	}
	return f(restAPIError, restAPIError.ErrorEntities...)
}

// RegisterTransformFunc is used to add new function to transform RestAPIError to InfraError
// if the restAPIErrorCode is already registered, the old transform function will be overridden
func (t *infraTransformer) RegisterTransformFunc(restAPIErrCode int, transformFunc infraTransformFunc) {
	t.mapping[restAPIErrCode] = transformFunc
}
