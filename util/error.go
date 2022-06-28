package util

type AppError struct {
	message string
}

func (err *AppError) Error() string {
	return err.message
}

func ThrowAppError(message string) *AppError {
	return &AppError{message: message}
}
