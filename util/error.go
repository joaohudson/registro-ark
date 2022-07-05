package util

import "strconv"

const DefaultInternalServerError = "Erro interno do servidor, por favor tente mais tarde."

type ApiError struct {
	Message    string
	StatusCode int
}

func (err *ApiError) Error() string {
	return "[" + strconv.FormatInt(int64(err.StatusCode), 10) + "] - " + err.Message
}

func ThrowApiError(message string, statusCode int) *ApiError {
	return &ApiError{Message: message, StatusCode: statusCode}
}
