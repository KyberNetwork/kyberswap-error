package constants

const (
	InfraErrCodeDBConnect = "INFRA:DATABASE:CONNECT"
	InfraErrMsgDBConnect  = "Infra error: Failed to connect to database"

	InfraErrCodeDBNotFound = "INFRA:DATABASE:NOT_FOUND"
	InfraErrMsgDBNotFound  = "Infra error: Not found resource in database"

	InfraErrCodeDBSelect = "INFRA:DATABASE:SELECT"
	InfraErrMsgDBSelect  = "Infra error: Failed to select resources from database"

	InfraErrCodeDBInsert = "INFRA:DATABASE:INSERT"
	InfraErrMsgDBInsert  = "Infra error: Failed to insert into database"

	InfraErrCodeDBUpdate = "INFRA:DATABASE:UPDATE"
	InfraErrMsgDBUpdate  = "Infra error: Failed to update resources in database"

	InfraErrCodeDBDelete = "INFRA:DATABASE:DELETE"
	InfraErrMsgDBDelete  = "Infra error: Failed to delete resources from database"

	InfraErrCodeDBUnknown = "INFRA:DATABASE:UNKNOWN"
	InfraErrMsgDBUnknown  = "Infra error: Unknown database error"

	InfraErrCodeRedisConnect = "INFRA:REDIS:CONNECT"
	InfraErrMsgRedisConnect  = "Infra error: Failed to connect to Redis"

	InfraErrCodeRedisNotFound = "INFRA:REDIS:NOT_FOUND"
	InfraErrMsgRedisNotFound  = "Infra error: Not found resource in Redis"

	InfraErrCodeRedisGet = "INFRA:REDIS:GET"
	InfraErrMsgRedisGet  = "Infra error: Failed to get data from Redis"

	InfraErrCodeRedisSet = "INFRA:REDIS:SET"
	InfraErrMsgRedisSet  = "Infra error: Failed to write data to Redis"

	InfraErrCodeRedisUnknown = "INFRA:REDIS:UNKNOWN"
	InfraErrMsgRedisUnknown  = "Infra error: Unknown redis error"

	InfraErrCodeHTTPRequired = "INFRA:HTTP:REQUIRED"
	InfraErrMsgHTTPRequired  = "Infra error: Missing required fields when calling HTTP error"

	InfraErrCodeHTTPNotAcceptedValue = "INFRA:HTTP:NOT_ACCEPTED_VALUE"
	InfraErrMsgHTTPNotAcceptedValue  = "Infra error: Input is not in accepted value when calling HTTP error"

	InfraErrCodeHTTPOutOfRange = "INFRA:HTTP:OUT_OF_RANGE"
	InfraErrMsgHTTPOutOfRange  = "Infra error: Input is out of range when calling HTTP error"

	InfraErrCodeHTTPInvalidFormat = "INFRA:HTTP:INVALID_FORMAT"
	InfraErrMsgHTTPInvalidFormat  = "Infra error: Input is invalid format when calling HTTP error"

	InfraErrCodeHTTPInvalid = "INFRA:HTTP:INVALID"
	InfraErrMsgHTTPInvalid  = "Infra error: Input is invalid when calling HTTP error"

	InfraErrCodeHTTPUnauthenticated = "INFRA:HTTP:UNAUTHENTICATED"
	InfraErrMsgHTTPUnauthenticated  = "Infra error: Unauthenticated HTTP error"

	InfraErrCodeHTTPUnauthorized = "INFRA:HTTP:UNAUTHORIZED"
	InfraErrMsgHTTPUnauthorized  = "Infra error: Unauthorized HTTP error"

	InfraErrCodeHTTPNotFound = "INFRA:HTTP:NOTFOUND"
	InfraErrMsgHTTPNotFound  = "Infra error: NotFound HTTP error"

	InfraErrCodeHTTPDuplicate = "INFRA:HTTP:DUPLICATE"
	InfraErrMsgHTTPDuplicate  = "Infra error: Input is duplicated when calling HTTP error"

	InfraErrCodeHTTPAlreadyExists = "INFRA:HTTP:ALREADY_EXISTS"
	InfraErrMsgHTTPAlreadyExists  = "Infra error: Input already exists when calling HTTP error"

	InfraErrCodeHTTPUnknown = "INFRA:HTTP:UNKNOWN"
	InfraErrMsgHTTPUnknown  = "Infra error: Unknown HTTP error"

	InfraErrCodeRPCUnknown = "INFRA:RPC:UNKNOWN"
	InfraErrMsgRPCUnknown  = "Infra error: Unknown RPC error"

	InfraErrCodeElsConnect = "INFRA:ELS:CONNECT"
	InfraErrMsgElsConnect  = "Infra error: Failed to connect to Els"

	InfraErrCodeElsUnknown = "INFRA:ELS:UNKNOWN"
	InfraErrMsgElsUnknown  = "Infra error: Unknown Els error"

	InfraErrCodeElsNotFound = "INFRA:ELS:NOT_FOUND"
	InfraErrMsgElsNotFound  = "Infra error: Not found resource in Els"
)
