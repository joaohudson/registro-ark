package util

type AppError struct {
	message string
}

type ApiError struct {
	Message    string
	StatusCode int
}

func (err *AppError) Error() string {
	return err.message
}

func ThrowAppError(message string) *AppError {
	return &AppError{message: message}
}

func ThrowApiError(message string, statusCode int) *ApiError {
	return &ApiError{Message: message, StatusCode: statusCode}
}
