package base

import (
	"strings"
	"unicode"
)

func IsKeySuffix(str string) bool {
	return len(str) > 1 && str[len(str)-1:] == ":"
}

func IsSetterSuffix(str string) bool {
	return len(str) > 1 && str[len(str)-1:] == "="
}

func IsSymbol(str string) bool {
	return len(str) > 1 && str[0] == ':'
}

func IsDoubleAsteriskPrefix(str string) bool {
	return len(str) > 2 && str[0] == '*' && str[1] == '*'
}

func IsAsteriskPrefix(str string) bool {
	return len(str) > 1 && str[0] == '*'
}

func IsAtmarkPrefix(str string) bool {
	return len(str) > 1 && str[0] == '@'
}

func IsAmpersandPrefix(str string) bool {
	return len(str) > 1 && str[0] == '&'
}

func IsEqualPrefix(str string) bool {
	return len(str) > 1 && str[0] == '='
}

func IsNameSpace(str string) bool {
	return len(strings.Split(str, "::")) > 1
}

func IsUpper(str string) bool {
	if len(str) < 1 {
		return false
	}

	return unicode.IsUpper(rune(str[0]))
}

func IsContainsSymbol(str string) bool {
	for _, r := range str {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '_' && r != '=' {
			return true
		}
	}
	return false
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
