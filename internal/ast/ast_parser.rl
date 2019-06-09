package ast

// Code generated by ragel; DO NOT EDIT.

%%{
	machine lol;
	alphtype rune;
	write data;
}%%

func c(r int) rune {
	return rune(r)
}

func (parser *Parser) parse(data []rune) (*Program, error) {
	parser.data = data
	program := &Program{}

	eof := len(data);
	cs, p, pe := 0, 0, eof;
	%%{

	action addVISIBLE {
	    program.addVISIBLE(parser.getMark())
	}
	action addYARN {
	    program.addYARN(parser.getYARN(fpc))
	}
	action setMark {
	    parser.setMark(fpc)
	}
	action trace {
	    parser.trace(rune(fc), fpc)
	}

	eol = ( '\r'? '\n' | '\r' '\n'? ) %from{ parser.startLine(fpc) };
	sep = [ \t]+;
	ws = ( sep | eol )+;

	head = sep? ('HAI' sep '1.2') $err{ parser.setError(fpc, "invalid version declaration") };
	tail = sep? ('KTHXBYE') $err{ parser.setError(fpc, "expected: \"KTHXBYE\"") };

	yarn = ('"' [^"]* '"') >setMark @addYARN;
	expr = yarn;

	visible = (
	    'VISIBLE' >setMark sep @addVISIBLE
	    expr $err{ parser.setError(fpc, "expected: expression") }
	    (sep expr $err{ parser.setError(fpc, "expected: expression") })*
	);
	statement = sep? visible;

	main := (eol* head eol
		(statement eol)*
		tail ws? ) $trace;

	write init;
	write exec;
	}%%

	if cs < lol_first_final {
		return nil, parser.getError(cs)
	}

	return program, nil
}
