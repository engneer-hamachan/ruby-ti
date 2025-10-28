package base

func IsKeySuffix(str string) bool {
	return str[len(str)-1:] == ":" && len(str) >= 2
}

func RemoveSuffix(str string) string {
	return str[:len(str)-1]
}
