package ast

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

func (stmt *Visible) Head() (int, int) {
	return stmt.Line, stmt.Column
}

func (stmt *Visible) addYARN(y YARN) {
	stmt.Expressions = append(stmt.Expressions, y)
}
