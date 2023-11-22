package loginReward

type Items []Item

type Item struct {
	ID int64 `json:"id"`

	Name string `json:"name"`

	Detail string `json:"detail"`
}
