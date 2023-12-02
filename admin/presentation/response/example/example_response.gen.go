package example

type Examples []Example

type Example struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	Detail *string `json:"detail"`

	Count int `json:"count"`
}

func ToExample(ID int64, Name string, Detail *string, Count int) *Example {
	return &Example{
		ID:     ID,
		Name:   Name,
		Detail: Detail,
		Count:  Count,
	}
}
