package ast

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnescape(t *testing.T) {
	require := require.New(t)

	for raw, expected := range map[string]YARN{
		`Spoon!`:               YARN("Spoon!"),
		`a:)b:>c`:              YARN("a\nb\tc"),
		`:: : :"Spoon!:" : ::`: YARN(`: : "Spoon!" : :`),
		"m:n:o":                YARN("m:n\a"),
	} {
		actual := unescape(raw)
		require.Equal(expected, actual, raw)
	}
}
