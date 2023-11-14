package example

type AddExample struct {
	Status int64 `json:"status"`

	UserID int64 `json:"user_id"`

	Name string `json:"name"`

	Detail *string `json:"detail"`

	Count int `json:"count"`

	Config *Config `json:"config"`
}

func AddExampleResponse(config *Config, status int64, userID int64, name string, detail *string, count int) *AddExample {
	return &AddExample{Config: config, Status: status, UserID: userID, Name: name, Detail: detail, Count: count}
}
