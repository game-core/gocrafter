package item

type Items []Item

type Item struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	Detail string `json:"detail"`

	Count int `json:"count"`
}

func ToItem(ID int64, Name string, Detail string, Count int) *Item {
	return &Item{
		ID:     ID,
		Name:   Name,
		Detail: Detail,
		Count:  Count,
	}
}
