package transformers

import (
	"fmt"

	c "github.com/KyberNetwork/kyberswap-error/pkg/constants"
	"github.com/KyberNetwork/kyberswap-error/pkg/errors"
)

type IDomainTransformer interface {
	InfraErrToDomainErr(infraErr *errors.InfraError) *errors.DomainError
	RegisterTransformFunc(infraErrCode string, transformFunc domainTransformFunc)
}

type domainTransformFunc func(rootCause error, entities ...string) *errors.DomainError

type domainTransformer struct {
	mapping map[string]domainTransformFunc
}

var domainTransformerInstance *domainTransformer

func initDomainTransformerInstance() {
	if domainTransformerInstance == nil {
		domainTransformerInstance = &domainTransformer{}
		domainTransformerInstance.mapping = map[string]domainTransformFunc{}
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeDBConnect, errors.NewDomainErrorUnknown)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeDBNotFound, errors.NewDomainErrorNotFound)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeDBSelect, errors.NewDomainErrorUnknown)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeDBInsert, errors.NewDomainErrorUnknown)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeDBUpdate, errors.NewDomainErrorUnknown)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeDBDelete, errors.NewDomainErrorUnknown)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeDBUnknown, errors.NewDomainErrorUnknown)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeRedisConnect, errors.NewDomainErrorUnknown)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeRedisNotFound, errors.NewDomainErrorNotFound)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeRedisGet, errors.NewDomainErrorUnknown)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeRedisSet, errors.NewDomainErrorUnknown)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeRedisUnknown, errors.NewDomainErrorUnknown)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeHTTPUnknown, errors.NewDomainErrorUnknown)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeHTTPNotFound, errors.NewDomainErrorNotFound)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeHTTPUnauthorized, errors.NewDomainErrorUnauthorized)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeRPCUnknown, errors.NewDomainErrorUnknown)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeElsConnect, errors.NewDomainErrorUnknown)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeElsUnknown, errors.NewDomainErrorUnknown)
		domainTransformerInstance.RegisterTransformFunc(c.InfraErrCodeElsNotFound, errors.NewDomainErrorNotFound)
	}
}

func DomainTransformerInstance() IDomainTransformer {
	return domainTransformerInstance
}

// InfraErrToDomainErr transforms InfraError to DomainError
func (t *domainTransformer) InfraErrToDomainErr(infraError *errors.InfraError) *errors.DomainError {
	f := t.mapping[infraError.Code]
	if f == nil {
		return errors.NewDomainErrorUnknown(fmt.Errorf("can not transform error, InfraError: %v", infraError))
	}
	return f(infraError, infraError.ErrorEntities...)
}

// RegisterTransformFunc is used to add new function to transform InternalError to DomainError
// if the infraErrCode is already registered, the old transform function will be overridden
func (t *domainTransformer) RegisterTransformFunc(infraErrCode string, function domainTransformFunc) {
	t.mapping[infraErrCode] = function
}
