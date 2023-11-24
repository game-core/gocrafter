package example

type ExampleSub int

const (
	ExampleSub1 ExampleSub = iota
	ExampleSub2
	ExampleSub3
)

func (e ExampleSub) ToString() string {
	names := [...]string{"ExampleSub1", "ExampleSub2", "ExampleSub3"}
	if e < ExampleSub1 || e > ExampleSub3 {
		return "Unknown"
	}

	return names[e]
}
