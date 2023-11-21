package example

type Examples []Example

type Example struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	Detail *string `json:"detail"`

	Count int `json:"count"`
}
