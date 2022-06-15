package constants

const (
	DomainErrCodeRequired = 4000
	DomainErrMsgRequired  = "Missing required fields"

	DomainErrCodeNotAcceptedValue = 4001
	DomainErrMsgNotAcceptedValue  = "Input is not in the accepted values"

	DomainErrCodeOutOfRange = 4002
	DomainErrMsgOutOfRange  = "Input is out of range"

	DomainErrCodeInvalidFormat = 4003
	DomainErrMsgInvalidFormat  = "Input has an invalid format"

	DomainErrCodeUnauthenticated = 4010
	DomainErrMsgUnauthenticated  = "Unauthenticated"

	DomainErrCodeNotFound = 4040
	DomainErrMsgNotFound  = "Not found"

	DomainErrCodeDuplicate = 4090
	DomainErrMsgDuplicate  = "Duplicate data"

	DomainErrCodeAlreadyExists = 4091
	DomainErrMsgAlreadyExists  = "Data already exists"

	DomainErrCodeInternal = 5000
	DomainErrMsgInternal  = "Internal Server Error"
)
