package reader

import (
	"bufio"
	"io"
)

type LexerReader struct {
	runes    []rune
	pos      int
	ungetFlg bool
	char     rune
	history  []rune
}

func New(r bufio.Reader) LexerReader {
	content, _ := io.ReadAll(&r)
	runes := []rune(string(content))

	return LexerReader{
		runes:    runes,
		pos:      0,
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

	if lr.pos >= len(lr.runes) {
		lr.char = 0
		return 0
	}

	lr.char = lr.runes[lr.pos]
	lr.pos++

	return lr.char
}

func (lr *LexerReader) AppendHistory(r rune) {
	lr.history = append(lr.history, r)
}

func (lr *LexerReader) Unread() {
	lr.ungetFlg = true
}
