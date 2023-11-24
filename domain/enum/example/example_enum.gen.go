package example

type Example int

const (
	Example1 Example = iota
	Example2
	Example3
)

func (e Example) ToString() string {
	names := [...]string{"Example1", "Example2", "Example3"}
	if e < Example1 || e > Example3 {
		return "Unknown"
	}

	return names[e]
}
