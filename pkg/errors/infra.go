package errors

import (
	"fmt"

	c "github.com/KyberNetwork/kyberswap-error/pkg/constants"
)

type InfraError struct {
	Code          string
	Message       string
	ErrorEntities []string
	RootCause     error
}

func NewInfraError(code string, message string, entities []string, rootCause error) *InfraError {
	return &InfraError{
		Code:          code,
		Message:       message,
		ErrorEntities: entities,
		RootCause:     rootCause,
	}
}

func (e *InfraError) Error() string {
	return fmt.Sprintf("INFRA ERROR: {Code: %s, Message: %s, ErrorEntities: %v, RootCause: %v}", e.Code, e.Message, e.ErrorEntities, e.RootCause)
}

func NewInfraErrorDBConnect(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeDBConnect, c.InfraErrMsgDBConnect, entities, rootCause)
}

func NewInfraErrorDBNotFound(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeDBNotFound, c.InfraErrMsgDBNotFound, entities, rootCause)
}

func NewInfraErrorDBSelect(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeDBSelect, c.InfraErrMsgDBSelect, entities, rootCause)
}

func NewInfraErrorDBInsert(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeDBInsert, c.InfraErrMsgDBInsert, entities, rootCause)
}

func NewInfraErrorDBUpdate(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeDBUpdate, c.InfraErrMsgDBUpdate, entities, rootCause)
}

func NewInfraErrorDBDelete(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeDBDelete, c.InfraErrMsgDBDelete, entities, rootCause)
}

func NewInfraErrorDBUnknown(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeDBUnknown, c.InfraErrMsgDBUnknown, entities, rootCause)
}

func NewInfraErrorRedisConnect(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeRedisConnect, c.InfraErrMsgRedisConnect, entities, rootCause)
}

func NewInfraErrorRedisNotFound(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeRedisNotFound, c.InfraErrMsgRedisNotFound, entities, rootCause)
}

func NewInfraErrorRedisGet(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeRedisGet, c.InfraErrMsgRedisGet, entities, rootCause)
}

func NewInfraErrorRedisSet(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeRedisSet, c.InfraErrMsgRedisSet, entities, rootCause)
}

func NewInfraErrorRedisUnknown(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeRedisUnknown, c.InfraErrMsgRedisUnknown, entities, rootCause)
}

func NewInfraErrorHTTPUnknown(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeHTTPUnknown, c.InfraErrMsgHTTPUnknown, entities, rootCause)
}

func NewInfraErrorHTTPNotFound(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeHTTPNotFound, c.InfraErrMsgHTTPNotFound, entities, rootCause)
}

func NewInfraErrorRPCUnknown(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeRPCUnknown, c.InfraErrMsgRPCUnknown, entities, rootCause)
}

func NewInfraErrorElsConnect(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeElsConnect, c.InfraErrMsgElsConnect, entities, rootCause)
}

func NewInfraErrorElsUnknown(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeElsUnknown, c.InfraErrMsgElsUnknown, entities, rootCause)
}

func NewInfraErrorElsNotFound(rootCause error, entities ...string) *InfraError {
	return NewInfraError(c.InfraErrCodeElsNotFound, c.InfraErrMsgElsNotFound, entities, rootCause)
}
