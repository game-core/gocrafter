package error

type Error struct {
	Status int64 `json:"status"`

	ErrorMessage string `json:"error_message"`
}

func ErrorResponse(errorMessage string, status int64) *Error {
	return &Error{ErrorMessage: errorMessage, Status: status}
}
