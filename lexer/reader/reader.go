package reader

import (
	"bufio"
)

type LexerReader struct {
	reader   bufio.Reader
	ungetFlg bool
	char     rune
	history  []rune
}

func New(r bufio.Reader) LexerReader {
	return LexerReader{
		reader:   r,
		ungetFlg: false,
	}
}

func (lr *LexerReader) Read() rune {
	if lr.ungetFlg {
		lr.ungetFlg = false

		return lr.char
	}

	if len(lr.history) > 0 {
		lr.char = lr.history[0]
		lr.history = lr.history[1:]

		return lr.char
	}

	lr.char, _, _ = lr.reader.ReadRune()

	return lr.char
}

func (lr *LexerReader) AppendHistory(r rune) {
	lr.history = append(lr.history, r)
}

func (lr *LexerReader) Unread() {
	lr.ungetFlg = true
}
