package pointer

func PointerToString(str *string) string {
	return *str
}

func StringToPointer(str string) *string {
	return &str
}
