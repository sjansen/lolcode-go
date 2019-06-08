package main_test

import (
	"io/ioutil"
	"path/filepath"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sjansen/lolcode-go/internal/ast"
)

func TestParser(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	filenames, err := filepath.Glob("testdata/*.lol")
	require.NoError(err)

	sort.Strings(filenames)
	for _, filename := range filenames {
		basename := filename[:len(filename)-4]

		p := &ast.Parser{}
		_, parseError := p.ParseFile(filename)
		if parseError == nil {
			expected, err := ioutil.ReadFile(basename + ".stdout")
			require.NoError(err)

			actual := ""
			if !assert.Equal(string(expected), actual) {
				ioutil.WriteFile(basename+".actual", []byte(actual), 0666)
			}
		} else {
			expected, err := ioutil.ReadFile(basename + ".stderr")
			require.NoError(err, parseError)

			actual := parseError.Error()
			if !assert.Equal(string(expected), actual) {
				ioutil.WriteFile(basename+".actual", []byte(actual), 0666)
			}
		}
	}
}
