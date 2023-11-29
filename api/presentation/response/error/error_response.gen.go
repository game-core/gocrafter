package error

type Errors []Error

type Error struct {
	Status int64 `json:"status"`

	ErrorMessage string `json:"error_message"`
}

func ToError(Status int64, ErrorMessage string) *Error {
	return &Error{
		Status:       Status,
		ErrorMessage: ErrorMessage,
	}
}
