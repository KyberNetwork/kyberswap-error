package constants

const (
	// HTTP 200 - OK
	ApiErrCodeOK = 0
	ApiErrMsgOK  = "Succeeded"

	//HTTP 400 - Bad Request
	ApiErrCodeRequired = 4000
	ApiErrMsgRequired  = "Missing required fields"

	ApiErrCodeNotAcceptedValue = 4001
	ApiErrMsgNotAcceptedValue  = "Input is not in the accepted values"

	ApiErrCodeOutOfRange = 4002
	ApiErrMsgOutOfRange  = "Input is out of range"

	ApiErrCodeInvalidFormat = 4003
	ApiErrMsgInvalidFormat  = "Input has an invalid format"

	//HTTP 401 - Unauthorized
	ApiErrCodeUnauthenticated = 4010
	ApiErrMsgUnauthenticated  = "Unauthenticated"

	//HTTP 404 - Not found
	ApiErrCodeNotFound = 4040
	ApiErrMsgNotFound  = "Not found"

	//HTTP 409 - Duplicate
	ApiErrCodeDuplicate = 4090
	ApiErrMsgDuplicate  = "Duplicate data"

	ApiErrCodeAlreadyExists = 4091
	ApiErrMsgAlreadyExists  = "Data already exists"

	//HTTP 500 - Internal Server Error
	ApiErrCodeInternal = 5000
	ApiErrMsgInternal  = "Internal Server Error"
)
