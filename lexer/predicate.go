package lexer

import (
	"ti/base"
	"unicode"
)

func isIdentifierChar(c rune) bool {
	if unicode.IsSpace(c) ||
		c == '\n' ||
		c == '(' ||
		c == ')' ||
		c == ',' ||
		c == '.' ||
		c == '{' ||
		c == '}' ||
		c == '=' ||
		c == '[' ||
		c == ']' ||
		c == '|' ||
		c == '&' ||
		c == '^' ||
		c == base.NIL {

		return false
	}

	return true
}
