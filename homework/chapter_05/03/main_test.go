package exercise

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  string
	ans string
}{
	{
		"a",
		"42\tJames Bond\tfalse",
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, exercise(tc.N1), "输入:%v", tc)
	}
}
