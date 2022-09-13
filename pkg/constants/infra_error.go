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

	InfraErrCodeHTTPUnknown = "INFRA:HTTP:UNKNOWN"
	InfraErrMsgHTTPUnknown  = "Infra error: Unknown HTTP error"

	InfraErrCodeHTTPNotFound = "INFRA:HTTP:NOTFOUND"
	InfraErrMsgHTTPNotFound  = "Infra error: NotFound HTTP error"

	InfraErrCodeRPCUnknown = "INFRA:RPC:UNKNOWN"
	InfraErrMsgRPCUnknown  = "Infra error: Unknown RPC error"

	InfraErrCodeElsConnect = "INFRA:ELS:CONNECT"
	InfraErrMsgElsConnect  = "Infra error: Failed to connect to Els"

	InfraErrCodeElsUnknown = "INFRA:ELS:UNKNOWN"
	InfraErrMsgElsUnknown  = "Infra error: Unknown Els error"

	InfraErrCodeElsNotFound = "INFRA:ELS:NOT_FOUND"
	InfraErrMsgElsNotFound  = "Infra error: Not found resource in Els"
)
