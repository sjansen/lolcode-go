package ast

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Parser struct {
	Trace bool

	data []rune

	errMessage string
	errColumn  int

	lineNumber int
	lineOffset int

	markOffset int
	markLineno int
	markColumn int
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
		errColumn:  p.errColumn,
		lineNumber: p.lineNumber,
		lineOffset: p.lineOffset,
	}
}

func (p *Parser) getMark() (int, int) {
	return p.markLineno, p.markColumn
}

func (p *Parser) getYARN(offset int) YARN {
	runes := p.data[p.markOffset+1 : offset-1]
	s := unescape(string(runes))
	return YARN(s)
}

func (p *Parser) setError(offset int, message string) {
	p.errMessage = message
	p.errColumn = offset - p.lineOffset
}

func (p *Parser) setMark(offset int) {
	p.markOffset = offset
	p.markLineno = p.lineNumber
	p.markColumn = offset - p.lineOffset + 1
}

func (p *Parser) startLine(offset int) {
	p.lineNumber++
	p.lineOffset = offset
}

func (p *Parser) trace(r rune, offset int) {
	if p.Trace {
		fmt.Fprintf(
			os.Stderr, "trace: char=%-4q\toffset=%-3d\tlineNumber=%-3d\tlineOffset=%-3d\n",
			r, offset, p.lineNumber, p.lineOffset,
		)
	}
}
