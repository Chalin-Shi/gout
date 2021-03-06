package e

const (
	ERROR = "500"

	ERROR_AUTH_CHECK_TOKEN_NULL    = "20000"
	ERROR_AUTH_CHECK_TOKEN_FAIL    = "20001"
	ERROR_AUTH_CHECK_TOKEN_EXPIRED = "20002"
	ERROR_AUTH_TOKEN               = "20003"
	ERROR_AUTH                     = "20004"

	UNKNOW_ERROR           = "-1"
	SUCCESS                = "100000"
	API_INTERNAL_ERROR     = "200000"
	PERMISSION_DENIED      = "210000"
	INVALID_PARAMS         = "220000"
	SEND_EMAIL_ERROR       = "230000"
	GIT_UTIL_ERROR         = "240000"
	GIT_FLOW_ERROR         = "250000"
	DATABASE_ERROR         = "300000"
	RECORD_HAS_EXISTED     = "310000"
	RECORD_MULTI_EXIST     = "320000"
	RECORD_NOT_EXIST       = "330000"
	FILE_NOT_EXIST         = "340000"
	FILE_UPLOAD_FAILED     = "350000"
	APP_NOT_INSTALLED      = "360000"
	OOS_ERROR              = "380000"
	KERBEROS_ERROR         = "390000"
	VALIDATION_ERROR       = "400000"
	PASSWORD_NOT_MATCH     = "410000"
	ORIGIN_PASSWORD_ERROR  = "420000"
	VERIFICATION_NOT_MATCH = "430000"
	CLUSTER_NOT_EXIST      = "500000"
	HTTP_REQUEST_ERROR     = "600000"
	PLATFORM_REQUEST_ERROR = "700000"
	INTERNAL_CALL_ERROR    = "800000"
	PYTHON_CALL_ERROR      = "810000"
	SHELL_COMMAND_ERROR    = "820000"
)
