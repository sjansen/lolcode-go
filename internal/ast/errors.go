package ast

import (
	"fmt"
)

type parseError struct {
	cs         int
	data       []rune
	errMessage string
	errColumn  int
	lineNumber int
	lineOffset int
}

func (e *parseError) Error() string {
	line := e.data[e.lineOffset:]
	for i := range line {
		if line[i] == '\n' {
			line = line[:i]
			break
		}
	}

	return fmt.Sprintf(
		"error: %s (line=%d col=%d)\n  %s\n  %*s\n",
		e.errMessage,
		e.lineNumber+1,
		e.errColumn+1,
		string(line),
		e.errColumn+1,
		"^",
	)
}
