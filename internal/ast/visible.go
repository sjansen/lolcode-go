package ast

import "fmt"

func (p *Program) addVISIBLE(line, column int) {
	p.Statements = append(p.Statements, &Visible{
		Line:   line,
		Column: column,
	})
}

type Visible struct {
	Line   int
	Column int

	Expressions []YARN
}

func (stmt *Visible) Execute(env *Environment) {
	if len(stmt.Expressions) > 0 {
		fmt.Fprint(env.Output, stmt.Expressions[0])
	}
	for _, expr := range stmt.Expressions[1:] {
		fmt.Fprint(env.Output, " ", expr)
	}
	fmt.Fprint(env.Output, "\n")
}

func (stmt *Visible) addYARN(y YARN) {
	stmt.Expressions = append(stmt.Expressions, y)
}
