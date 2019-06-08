package ast

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Parser struct {
	data       []rune
	errMessage string
	errOffset  int
	lineNumber int
	lineOffset int
}

func (p *Parser) Parse(r io.Reader) (*Program, error) {
	raw, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return p.parse(bytes.Runes(raw))
}

func (p *Parser) ParseFile(filename string) (*Program, error) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return p.parse(bytes.Runes(raw))
}

func (p *Parser) getError(cs int) *parseError {
	return &parseError{
		cs:         cs,
		data:       p.data,
		errMessage: p.errMessage,
		errOffset:  p.errOffset,
		lineNumber: p.lineNumber,
		lineOffset: p.lineOffset,
	}
}

func (p *Parser) setError(offset int, message string) {
	p.errMessage = message
	p.errOffset = offset - p.lineOffset
}

func (p *Parser) startLine(offset int) {
	p.lineNumber++
	p.lineOffset = offset
}

func (p *Parser) trace(r rune, offset int) {
	fmt.Fprintf(os.Stderr, "trace: char=%q offset=%v\n", r, offset)
}
