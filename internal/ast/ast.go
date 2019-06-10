package ast

import "io"

//go:generate ragel-go -G2 -o ast_parser.go ast_parser.rl

type Program struct {
	Statements []Statement
}

type Environment struct {
	Output io.Writer
}

type Statement interface {
	addYARN(YARN)
	Execute(*Environment)
}

func (p *Program) addYARN(y YARN) {
	stmt := p.Statements[len(p.Statements)-1]
	stmt.addYARN(y)
}

func (p *Program) Execute(env *Environment) {
	for _, stmt := range p.Statements {
		stmt.Execute(env)
	}
}
