package example

type ListExamples []ListExample

type ListExample struct {
	Status int64 `json:"status"`

	Items *Examples `json:"items"`
}
