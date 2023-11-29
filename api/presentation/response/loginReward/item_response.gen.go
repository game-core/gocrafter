package loginReward

type Items []Item

type Item struct {
	Name string `json:"name"`

	Count int `json:"count"`
}

func ToItem(Name string, Count int) *Item {
	return &Item{
		Name:  Name,
		Count: Count,
	}
}
