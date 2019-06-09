package ast

//go:generate ragel-go -G2 -o ast_parser.go ast_parser.rl

type Program struct {
	Statements []Statement
}

type Statement interface {
	Head() (int, int)
	addYARN(YARN)
}

func (p *Program) addYARN(y YARN) {
	stmt := p.Statements[len(p.Statements)-1]
	stmt.addYARN(y)
}
