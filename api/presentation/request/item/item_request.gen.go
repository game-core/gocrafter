package item

type Items []Item

type Item struct {
	Name string `json:"name"`

	Count int `json:"count"`
}
