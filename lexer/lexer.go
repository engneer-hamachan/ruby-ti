package lexer

import (
	"strconv"
	"strings"
	"ti/base"
	"ti/lexer/reader"
	"unicode"
)

type Lexer struct {
	tok         rune
	val         any
	IsSpace     bool
	IsSpacePrev bool
	reader      reader.LexerReader
	LastComment string
}

var reserved map[string]any = make(map[string]any)

func New(lr reader.LexerReader) Lexer {
	reserved["nil"] = rune(base.NIL)

	return Lexer{
		reader: lr,
	}
}

func (l *Lexer) Token() rune {
	return l.tok
}

func (l *Lexer) Value() any {
	return l.val
}

func (l *Lexer) setDigitToken(buf strings.Builder) {
	var err error

	l.val, err = strconv.ParseInt(buf.String(), 10, 64)
	l.tok = base.INT

	if err != nil {
		l.val, _ = strconv.ParseFloat(buf.String(), 64)
		l.tok = base.FLOAT

		return
	}
}

func (l *Lexer) lexToSpaceTokenEat(currentChar rune) strings.Builder {
	var buf strings.Builder

	buf.WriteRune(currentChar)

	for {
		char := l.reader.Read()

		if unicode.IsSpace(char) {
			if char != '\n' {
				l.IsSpace = true
			}

			l.reader.Unread()

			return buf
		}

		buf.WriteRune(char)
	}
}

func (l *Lexer) lexToNotIdentifierTokenEat(currentChar rune) strings.Builder {
	var buf strings.Builder

	buf.WriteRune(currentChar)

	for {
		char := l.reader.Read()

		if !isIdentifierChar(char) {
			if char == ' ' {
				l.IsSpace = true
			}

			l.reader.Unread()

			return buf
		}

		buf.WriteRune(char)
	}
}

func (l *Lexer) lexHexDigits(currentChar rune) strings.Builder {
	var buf strings.Builder

	buf.WriteRune(currentChar)

	for {
		char := l.reader.Read()

		if char == 'x' || char == 'o' || char == 'b' {
			buf.WriteRune(char)
			continue
		}

		if !((char >= '0' && char <= '9') || (char >= 'a' && char <= 'f') || (char >= 'A' && char <= 'F')) {
			l.reader.Unread()

			return buf
		}

		buf.WriteRune(char)
	}
}

func (l *Lexer) lexDigit() {
	var buf strings.Builder

	for {
		char := l.reader.Read()

		switch char {
		//0xff
		case 'x', 'o', 'b':
			l.reader.Unread()

			l.tok = base.UNKNOWN

			l.lexHexDigits(char)

			l.val = int64(0)
			l.tok = base.INT

			return

		// 1.23
		case '.':
			nextChar := l.reader.Read()

			if !unicode.IsDigit(nextChar) {
				l.reader.AppendHistory(char)
				l.reader.AppendHistory(nextChar)

				l.setDigitToken(buf)

				return
			}

		default:
			if !unicode.IsDigit(char) {
				l.reader.Unread()

				l.setDigitToken(buf)

				return
			}
		}

		buf.WriteRune(char)
	}
}

func (l *Lexer) lexIdentifier(currentChar rune) {
	l.tok = base.UNKNOWN

	var buf strings.Builder

	for {
		char := l.reader.Read()

		// *=
		if currentChar == '*' && char == '=' {
			buf.WriteRune(char)
			break
		}

		if !isIdentifierChar(char) {
			if strings.Contains(buf.String(), ":\"") && char != '\n' && char != '"' {
				buf.WriteRune(char)
				continue
			}

			l.reader.Unread()
			break
		}

		buf.WriteRune(char)
	}

	str := buf.String()
	l.val = Intern(str)

	val, ok := reserved[str]
	if ok {
		l.tok = val.(rune)
	}
}

func (l *Lexer) lexString(start rune) {
	var buf strings.Builder

	for {
		char := l.reader.Read()

		if char == start {
			break
		}

		if char == '\\' {
			char = l.reader.Read()
		}

		buf.WriteRune(char)
	}

	l.val = buf.String()
}

func (l *Lexer) skipSpace() {
	char := l.reader.Read()
	if unicode.IsSpace(char) {
		l.IsSpace = true
	}

	for {
		if !unicode.IsSpace(char) || char == '\n' {
			break
		}

		char = l.reader.Read()

		l.IsSpace = true
	}

	l.reader.Unread()
}

func (l *Lexer) skipLineComment() {
	var char rune
	var buf strings.Builder

	for {
		char = l.reader.Read()

		if char == '\n' {
			break
		}

		buf.WriteRune(char)
	}

	comment := buf.String()
	if strings.Contains(comment, "ti-doc:") {
		l.LastComment = strings.TrimSpace(strings.Replace(comment, "ti-doc:", "", 1))
	}

	l.reader.Unread()
}

func (l *Lexer) Advance() bool {
	l.skipSpace()
	char := l.reader.Read()

	switch char {
	case '<', '>':
		buf := l.lexToSpaceTokenEat(char)

		str := buf.String()
		l.tok = base.UNKNOWN
		l.val = Intern(str)

	case '=':
		var buf strings.Builder
		buf.WriteRune(char)

		nextChar := l.reader.Read()

		// =>
		if nextChar == '>' {
			buf.WriteRune(nextChar)
			str := buf.String()
			l.tok = base.UNKNOWN
			l.val = Intern(str)

			return true
		}

		if nextChar != '=' {
			l.reader.Unread()

			str := buf.String()
			l.tok = base.UNKNOWN
			l.val = Intern(str)

			return true
		}

		// ==
		buf.WriteRune(nextChar)

		nextChar = l.reader.Read()
		if nextChar != '=' {
			l.reader.Unread()

			str := buf.String()
			l.tok = base.UNKNOWN
			l.val = Intern(str)

			return true
		}

		// ===
		buf.WriteRune(nextChar)

		str := buf.String()
		l.tok = base.UNKNOWN
		l.val = Intern(str)

		return true

	case '.':
		nextChar := l.reader.Read()

		// 1..2
		if nextChar == '.' {
			var buf strings.Builder

			buf.WriteRune(char)
			buf.WriteRune(nextChar)

			//1...2
			nextChar := l.reader.Read()

			switch nextChar {
			case '.':
				buf.WriteRune(nextChar)
			default:
				l.reader.Unread()
			}

			str := buf.String()
			l.val = Intern(str)

			l.tok = base.UNKNOWN

			break
		}

		// .
		l.tok = char
		l.reader.Unread()

	case '%':
		var buf strings.Builder
		nextChar := l.reader.Read()
		buf.WriteRune(char)

		switch nextChar {
		case '=', 'W', 'w', 'i', 'Q', 'q', 'r', 's', 'l', 'x':
			buf.WriteRune(nextChar)
			str := buf.String()
			l.val = Intern(str)
			l.tok = base.UNKNOWN

		default:
			l.reader.Unread()

			l.tok = base.UNKNOWN

			buf = l.lexToSpaceTokenEat(char)

			str := buf.String()
			l.val = Intern(str)
		}

	case '!', '+', '-', '/':
		var buf strings.Builder
		nextChar := l.reader.Read()
		buf.WriteRune(char)

		switch nextChar {
		case '=':
			buf.WriteRune(nextChar)

		default:
			if char == '-' && nextChar == '>' {
				buf.WriteRune(nextChar)
				break
			}

			l.reader.Unread()
		}

		if (char == '+' || char == '-') && unicode.IsDigit(nextChar) {
			l.lexDigit()
			break
		}

		str := buf.String()
		l.val = Intern(str)
		l.tok = base.UNKNOWN

	case '&':
		var buf strings.Builder
		nextChar := l.reader.Read()
		buf.WriteRune(char)

		switch nextChar {
		case '.', '&':
			buf.WriteRune(nextChar)
			str := buf.String()
			l.val = Intern(str)
			l.tok = base.UNKNOWN

		default:
			l.reader.Unread()

			l.tok = base.UNKNOWN

			buf := l.lexToNotIdentifierTokenEat(char)

			str := buf.String()
			l.val = Intern(str)
		}

	case '|':
		var buf strings.Builder
		buf.WriteRune(char)

		nextChar := l.reader.Read()

		switch nextChar {
		case '|':
			buf.WriteRune(nextChar)

		default:
			l.reader.Unread()
		}

		nextChar = l.reader.Read()

		switch nextChar {
		case '=':
			buf.WriteRune(nextChar)

		default:
			l.reader.Unread()
		}

		str := buf.String()
		l.val = Intern(str)
		l.tok = base.UNKNOWN

	case '^':
		l.tok = char

	case '\n', '(', ')', '`', ',', '{', '}', '[', ']':
		l.tok = char

	case '"', '\'':
		l.lexString(char)
		l.tok = base.STRING

	case '#':
		nextChar := l.reader.Read()

		switch nextChar {
		case '{':
			var buf strings.Builder
			buf.WriteRune(char)
			buf.WriteRune(nextChar)
			str := buf.String()
			l.val = Intern(str)
			l.tok = base.UNKNOWN

		default:
			l.reader.Unread()
			l.skipLineComment()

			return l.Advance()
		}

	default:
		if unicode.IsDigit(char) {
			l.reader.Unread()

			l.lexDigit()

			break
		}

		if isIdentifierChar(char) {
			l.reader.Unread()
			l.lexIdentifier(char)

			break
		}

		return false
	}

	return true
}
