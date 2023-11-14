package error

type Error struct {
	Status int64 `json:"status"`

	ErrorMessage string `json:"error_message"`
}

func ErrorResponse(status int64, errorMessage string) *Error {
	return &Error{Status: status, ErrorMessage: errorMessage}
}
