package e

var eMap = map[string]string{
	ERROR: "fail",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token check failed",
	ERROR_AUTH_CHECK_TOKEN_EXPIRED: "Token is expired",
	ERROR_AUTH_TOKEN:               "Token generate failed",
	ERROR_AUTH:                     "Token is null",

	UNKNOW_ERROR:           "Unknow error",
	SUCCESS:                "Success",
	API_INTERNAL_ERROR:     "API internal error",
	PERMISSION_DENIED:      "Permission denied",
	INVALID_PARAMS:         "Invalid params",
	SEND_EMAIL_ERROR:       "Send eamil error",
	GIT_UTIL_ERROR:         "Git util error",
	GIT_FLOW_ERROR:         "Git flow error",
	DATABASE_ERROR:         "Database error",
	RECORD_HAS_EXISTED:     "Record has existed",
	RECORD_MULTI_EXIST:     "Record multi exist",
	RECORD_NOT_EXIST:       "Record not exist",
	FILE_NOT_EXIST:         "File not exist",
	FILE_UPLOAD_FAILED:     "File upload failed",
	APP_NOT_INSTALLED:      "App not installed error",
	OOS_ERROR:              "Aliyun oss error",
	KERBEROS_ERROR:         "Kerberos error",
	VALIDATION_ERROR:       "Validation error",
	PASSWORD_NOT_MATCH:     "Username and password don't match",
	ORIGIN_PASSWORD_ERROR:  "Origin password not match",
	VERIFICATION_NOT_MATCH: "Verification not match",
	CLUSTER_NOT_EXIST:      "Cluster not exist",
	HTTP_REQUEST_ERROR:     "Http request error",
	PLATFORM_REQUEST_ERROR: "Platform request error",
	INTERNAL_CALL_ERROR:    "Internal call error",
	PYTHON_CALL_ERROR:      "Python call error",
	SHELL_COMMAND_ERROR:    "Shell command error",
}

func GetMsg(code string) map[string]string {
	info, ok := eMap[code]
	if !ok {
		info = eMap[UNKNOW_ERROR]
	}
	message := map[string]string{"desc": info, "raw": info}
	if code != SUCCESS {
		message["name"] = code
	}
	return message
}
