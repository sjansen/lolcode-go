package main_test

import (
	"bytes"
	"io/ioutil"
	"os"
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

		parser := &ast.Parser{}
		program, parseError := parser.ParseFile(filename)
		if parseError == nil {
			expected, err := ioutil.ReadFile(basename + ".stdout")
			if os.IsNotExist(err) {
				expected = []byte{}
			} else {
				require.NoError(err)
			}

			buf := &bytes.Buffer{}
			env := &ast.Environment{
				Output: buf,
			}

			program.Execute(env)

			actual := buf.String()
			if !assert.Equal(string(expected), actual, filename) {
				ioutil.WriteFile(basename+".actual", []byte(actual), 0666)
			}
		} else {
			expected, err := ioutil.ReadFile(basename + ".stderr")
			if os.IsNotExist(err) {
				expected = []byte{}
			} else {
				require.NoError(err)
			}

			actual := parseError.Error()
			if !assert.Equal(string(expected), actual, filename) {
				ioutil.WriteFile(basename+".actual", []byte(actual), 0666)
			}
		}
	}
}
