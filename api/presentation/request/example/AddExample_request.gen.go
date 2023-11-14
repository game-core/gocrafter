package example

type AddExample struct {
	UserID int64 `json:"user_id"`

	Name string `json:"name"`

	Detail *string `json:"detail"`

	Count int `json:"count"`

	Config *Config `json:"config"`
}
