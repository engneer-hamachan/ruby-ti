package base

import "strings"

func IsKeySuffix(str string) bool {
	return len(str) > 1 && str[len(str)-1:] == ":"
}

func RemoveSuffix(str string) string {
	return str[:len(str)-1]
}

func SplitNameSpace(str string) []string {
	return strings.Split(str, "::")
}

func SeparateNameSpaces(str string) (string, string, string) {
	spaces := SplitNameSpace(str)

	// Hoge
	if len(spaces) == 1 {
		return "", "", spaces[0]
	}

	// Hoge::Fuga
	if len(spaces) == 2 {
		return "", spaces[0], spaces[1]
	}

	// Hoge::Fuga::Piyo
	frame := spaces[0]
	for _, name := range spaces[1 : len(spaces)-2] {
		frame += "::" + name
	}

	return frame, spaces[len(spaces)-2:][0], spaces[len(spaces)-2:][1]
}
