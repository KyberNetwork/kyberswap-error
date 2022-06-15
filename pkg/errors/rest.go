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
	return fmt.Sprintf("An API Error occured: Code: %d, Messaage: %s, ErrorEntities: %v, RootCause: %v", e.Code, e.Message, e.ErrorEntities, e.RootCause)
}

func AppendEntitiesToErrMsg(message string, entities []string) string {
	if len(entities) > 0 {
		message += ": "
		message += strings.Join(entities, ",")
	}
	return message
}

func NewRestAPIErrRequired(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ApiErrMsgRequired, entities)
	return NewRestAPIError(http.StatusBadRequest, c.ApiErrCodeRequired, message, entities, rootCause)
}

func NewRestAPIErrInvalidFormat(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ApiErrMsgInvalidFormat, entities)
	return NewRestAPIError(http.StatusBadRequest, c.ApiErrCodeInvalidFormat, message, entities, rootCause)
}

func NewRestAPIErrNotAcceptedValue(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ApiErrMsgNotAcceptedValue, entities)
	return NewRestAPIError(http.StatusBadRequest, c.ApiErrCodeNotAcceptedValue, message, entities, rootCause)
}

func NewRestAPIErrOutOfRange(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ApiErrMsgOutOfRange, entities)
	return NewRestAPIError(http.StatusBadRequest, c.ApiErrCodeOutOfRange, message, entities, rootCause)
}

func NewRestAPIErrUnauthenticated(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ApiErrMsgUnauthenticated, entities)
	return NewRestAPIError(http.StatusUnauthorized, c.ApiErrCodeUnauthenticated, message, entities, rootCause)
}

func NewRestAPIErrNotFound(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ApiErrMsgNotFound, entities)
	return NewRestAPIError(http.StatusNotFound, c.ApiErrCodeNotFound, message, entities, rootCause)
}

func NewRestAPIErrDuplicate(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ApiErrMsgDuplicate, entities)
	return NewRestAPIError(http.StatusConflict, c.ApiErrCodeDuplicate, message, entities, rootCause)
}

func NewRestAPIErrAlreadyExits(rootCause error, entities ...string) *RestAPIError {
	message := AppendEntitiesToErrMsg(c.ApiErrMsgAlreadyExists, entities)
	return NewRestAPIError(http.StatusConflict, c.ApiErrCodeAlreadyExists, message, entities, rootCause)
}

func NewRestAPIErrInternal(rootCause error, entities ...string) *RestAPIError {
	return NewRestAPIError(http.StatusInternalServerError, c.ApiErrCodeInternal, c.ApiErrMsgInternal, entities, rootCause)
}
