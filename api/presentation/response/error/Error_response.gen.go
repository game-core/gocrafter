package error

type Errors []Error

type Error struct {
	Status int64 `json:"status"`

	ErrorMessage string `json:"error_message"`
}
