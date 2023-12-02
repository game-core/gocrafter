package example

type GetExamples []GetExample

type GetExample struct {
	Status int64 `json:"status"`

	Example *Example `json:"example"`
}

func ToGetExample(Status int64, Example *Example) *GetExample {
	return &GetExample{
		Status:  Status,
		Example: Example,
	}
}
