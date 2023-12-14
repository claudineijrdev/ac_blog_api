package exception

type HttpError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func NewHttpError(message string, code int) *HttpError {
	return &HttpError{
		Message: message,
		Code:    code,
	}
}
