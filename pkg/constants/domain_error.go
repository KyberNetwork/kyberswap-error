package constants

const (
	DomainErrCodeRequired = "DOMAIN:REQUIRED"
	DomainErrMsgRequired  = "Domain error: Missing required fields"

	DomainErrCodeNotAcceptedValue = "DOMAIN:NOT_ACCEPTED_VALUE"
	DomainErrMsgNotAcceptedValue  = "Domain error: Input is not in the accepted values"

	DomainErrCodeOutOfRange = "DOMAIN:OUT_OF_RANGE"
	DomainErrMsgOutOfRange  = "Domain error: Input is out of range"

	DomainErrCodeInvalidFormat = "DOMAIN:INVALID_FORMAT"
	DomainErrMsgInvalidFormat  = "Domain error: Input has an invalid format"

	DomainErrCodeInvalid = "DOMAIN:INVALID"
	DomainErrMsgInvalid  = "Domain error: Input is invalid "

	DomainErrCodeUnauthenticated = "DOMAIN:UNAUTHENTICATED"
	DomainErrMsgUnauthenticated  = "Domain error: Unauthenticated"

	DomainErrCodeUnauthorized = "DOMAIN:UNAUTHORIZED"
	DomainErrMsgUnauthorized  = "Domain error: Unauthorized"

	DomainErrCodeNotFound = "DOMAIN:NOT_FOUND"
	DomainErrMsgNotFound  = "Domain error: Not found"

	DomainErrCodeDuplicate = "DOMAIN:DUPLICATE"
	DomainErrMsgDuplicate  = "Domain error: Duplicate data"

	DomainErrCodeAlreadyExists = "DOMAIN:ALREADY_EXISTS"
	DomainErrMsgAlreadyExists  = "Domain error: Data already exists"

	DomainErrCodeUnknown = "DOMAIN:UNKNOWN"
	DomainErrMsgUnknown  = "Domain error: Unknown Domain Error"
)
