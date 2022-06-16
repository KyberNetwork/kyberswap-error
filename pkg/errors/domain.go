package errors

import (
	"fmt"

	c "github.com/KyberNetwork/kyberswap-error/pkg/constants"
)

type DomainError struct {
	Code          string
	Message       string
	ErrorEntities []string
	RootCause     error
}

func NewDomainError(code string, message string, entities []string, rootCause error) *DomainError {
	return &DomainError{
		Code:          code,
		Message:       message,
		ErrorEntities: entities,
		RootCause:     rootCause,
	}
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("DOMAIN ERROR: {Code: %s, Messaage: %s, ErrorEntities: %v, RootCause: %v}", e.Code, e.Message, e.ErrorEntities, e.RootCause)
}

func NewDomainErrorRequired(rootCause error, entities ...string) *DomainError {
	return NewDomainError(c.DomainErrCodeRequired, c.DomainErrMsgRequired, entities, rootCause)
}

func NewDomainErrorInvalidFormat(rootCause error, entities ...string) *DomainError {
	return NewDomainError(c.DomainErrCodeInvalidFormat, c.DomainErrMsgInvalidFormat, entities, rootCause)
}

func NewDomainErrorNotAcceptedValue(rootCause error, entities ...string) *DomainError {
	return NewDomainError(c.DomainErrCodeNotAcceptedValue, c.DomainErrMsgNotAcceptedValue, entities, rootCause)
}

func NewDomainErrorOutOfRange(rootCause error, entities ...string) *DomainError {
	return NewDomainError(c.DomainErrCodeOutOfRange, c.DomainErrMsgOutOfRange, entities, rootCause)
}

func NewDomainErrorUnauthenticated(rootCause error, entities ...string) *DomainError {
	return NewDomainError(c.DomainErrCodeUnauthenticated, c.DomainErrMsgUnauthenticated, entities, rootCause)
}

func NewDomainErrorNotFound(rootCause error, entities ...string) *DomainError {
	return NewDomainError(c.DomainErrCodeNotFound, c.DomainErrMsgNotFound, entities, rootCause)
}

func NewDomainErrorDuplicate(rootCause error, entities ...string) *DomainError {
	return NewDomainError(c.DomainErrCodeDuplicate, c.DomainErrMsgDuplicate, entities, rootCause)
}

func NewDomainErrorAlreadyExits(rootCause error, entities ...string) *DomainError {
	return NewDomainError(c.DomainErrCodeAlreadyExists, c.DomainErrMsgAlreadyExists, entities, rootCause)
}

func NewDomainErrorUnknown(rootCause error, entities ...string) *DomainError {
	return NewDomainError(c.DomainErrCodeUnknown, c.DomainErrMsgUnknown, entities, rootCause)
}
