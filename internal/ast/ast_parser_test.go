package ast

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	require := require.New(t)

	for idx, tc := range []struct {
		Program  string
		Expected *Program
	}{{Program: `
	    HAI 1.2
	      VISIBLE "Spoon!"
	      VISIBLE ""
	      VISIBLE "The cake is a lie."
	    KTHXBYE
	`, Expected: &Program{
		Statements: []Statement{
			&Visible{
				Line: 2, Column: 8,
				Expressions: []YARN{
					"Spoon!",
				},
			}, &Visible{
				Line: 3, Column: 8,
				Expressions: []YARN{
					"",
				},
			}, &Visible{
				Line: 4, Column: 8,
				Expressions: []YARN{
					"The cake is a lie.",
				},
			},
		}},
	}, {Program: `
	    HAI 1.2
	      VISIBLE "Kilroy" "was" "here."
	    KTHXBYE
	`, Expected: &Program{
		Statements: []Statement{
			&Visible{
				Line: 2, Column: 8,
				Expressions: []YARN{
					"Kilroy", "was", "here.",
				},
			},
		}},
	}} {
		p := Parser{}
		actual, err := p.Parse(
			strings.NewReader(tc.Program),
		)
		require.NoError(err, idx)
		require.Equal(tc.Expected, actual, idx)
	}
}
