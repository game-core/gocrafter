package response

type Success struct {
	Types  string      `json:"types"`
	Status int         `json:"status"`
	Items  interface{} `json:"items"`
}

type Error struct {
	Types  string      `json:"types"`
	Status int         `json:"status"`
	Errors interface{} `json:"items"`
}

func SuccessWith(types string, status int, items interface{}) *Success {
	return &Success{
		Types:  types,
		Status: status,
		Items:  items,
	}
}

func ErrorWith(types string, status int, errors interface{}) *Error {
	return &Error{
		Types:  types,
		Status: status,
		Errors: errors,
	}
}
