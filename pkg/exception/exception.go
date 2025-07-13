package exception

type AppException struct {
	Code    int
	Message string
}

func (a *AppException) Error() string {
	return a.Message
}

func BadRequest(message string) error {
	if message == "" {
		message = "bad request"
	}
	return &AppException{
		Message: message,
		Code:    400,
	}
}

func InternalServerError() error {

	message := "internal server error"

	return &AppException{
		Code:    500,
		Message: message,
	}
}

func UnprocessableEntity(message string) error {
	if message == "" {
		message = "validation error"
	}
	return &AppException{
		Code:    422,
		Message: message,
	}
}

func NotFound(message string) error {
	if message == "" {
		message = "not found"
	}
	return &AppException{
		Code:    404,
		Message: message,
	}
}
