package errors

import (
	"fmt"
	"net/http"
	"strings"

	c "github.com/KyberNetwork/kyberswap-error/pkg/constants"
)

type RestAPIError struct {
	HttpStatus    int           `json:"-"`
	Code          int           `json:"code"`
	Message       string        `json:"message"`
	ErrorEntities []string      `json:"errorEntities"`
	Details       []interface{} `json:"details"`
	RootCause     error         `json:"-"`
}

func NewRestAPIError(httpStatus int, code int, message string, entities []string, rootCause error) *RestAPIError {
	return &RestAPIError{
		HttpStatus:    httpStatus,
		Code:          code,
		Message:       message,
		ErrorEntities: entities,
		RootCause:     rootCause,
	}
}

func (e *RestAPIError) Error() string {
	return fmt.Sprintf("API ERROR: {Code: %d, Message: %s, ErrorEntities: %v, RootCause: %v}", e.Code, e.Message, e.ErrorEntities, e.RootCause)
}

func AppendEntitiesToErrMsg(message string, entities []string) string {
	if len(entities) > 0 {
		message += ": "
		message += strings.Join(entities, ",")
	}
	return message
}

func NewRestAPIErrRequired(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ClientErrMsgRequired, entities)
	return NewRestAPIError(http.StatusBadRequest, c.ClientErrCodeRequired, message, entities, rootCause)
}

func NewRestAPIErrInvalidFormat(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ClientErrMsgInvalidFormat, entities)
	return NewRestAPIError(http.StatusBadRequest, c.ClientErrCodeInvalidFormat, message, entities, rootCause)
}

func NewRestAPIErrInvalid(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ClientErrMsgInvalid, entities)
	return NewRestAPIError(http.StatusBadRequest, c.ClientErrCodeInvalid, message, entities, rootCause)
}

func NewRestAPIErrNotAcceptedValue(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ClientErrMsgNotAcceptedValue, entities)
	return NewRestAPIError(http.StatusBadRequest, c.ClientErrCodeNotAcceptedValue, message, entities, rootCause)
}

func NewRestAPIErrOutOfRange(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ClientErrMsgOutOfRange, entities)
	return NewRestAPIError(http.StatusBadRequest, c.ClientErrCodeOutOfRange, message, entities, rootCause)
}

func NewRestAPIErrUnauthenticated(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ClientErrMsgUnauthenticated, entities)
	return NewRestAPIError(http.StatusUnauthorized, c.ClientErrCodeUnauthenticated, message, entities, rootCause)
}

func NewRestAPIErrUnauthorized(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ClientErrMsgUnauthorized, entities)
	return NewRestAPIError(http.StatusUnauthorized, c.ClientErrCodeUnauthorized, message, entities, rootCause)
}

func NewRestAPIErrNotFound(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ClientErrMsgNotFound, entities)
	return NewRestAPIError(http.StatusNotFound, c.ClientErrCodeNotFound, message, entities, rootCause)
}

func NewRestAPIErrDuplicate(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ClientErrMsgDuplicate, entities)
	return NewRestAPIError(http.StatusConflict, c.ClientErrCodeDuplicate, message, entities, rootCause)
}

func NewRestAPIErrAlreadyExits(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ClientErrMsgAlreadyExists, entities)
	return NewRestAPIError(http.StatusConflict, c.ClientErrCodeAlreadyExists, message, entities, rootCause)
}

func NewRestAPIErrInternal(rootCause error, entities ...string) *RestAPIError {
	return NewRestAPIError(http.StatusInternalServerError, c.ClientErrCodeInternal, c.ClientErrMsgInternal, entities, rootCause)
}
