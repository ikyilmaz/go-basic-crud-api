package lib

import "net/http"

func AppError(msg string, statusCode int) (int, interface{}) {

	var response = map[string]interface{}{
		"status":  "fail",
		"message": msg,
	}

	return statusCode, response
}

func defaultMessage(messages []string, defaultMsg string) string {
	if len(messages) > 0 {
		return messages[0]
	} else {
		return defaultMsg
	}
}

func NotFound(msg ...string) (int, interface{}) {
	return AppError(defaultMessage(msg, "not found"), http.StatusNotFound)
}

func BadRequest(msg ...string) (int, interface{}) {
	return AppError(defaultMessage(msg, "bad request"), http.StatusBadRequest)
}

func Unauthorized(msg ...string) (int, interface{}) {
	return AppError(defaultMessage(msg, "unauthorized"), http.StatusUnauthorized)
}

func Forbidden(msg ...string) (int, interface{}) {
	return AppError(defaultMessage(msg, "forbidden"), http.StatusForbidden)
}
