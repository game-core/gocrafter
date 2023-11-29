package example

type ListExamples []ListExample

type ListExample struct {
	Status int64 `json:"status"`

	Examples *Examples `json:"examples"`
}

func ToListExample(Status int64, Examples *Examples) *ListExample {
	return &ListExample{
		Status:   Status,
		Examples: Examples,
	}
}
