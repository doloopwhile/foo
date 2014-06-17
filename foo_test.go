package foo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type fooTest struct {
	in, out int
}

func TestFoo(t *testing.T) {
	tests := []fooTest{
		fooTest{1, 1},
		fooTest{2, 2},
	}
	for _, ft := range tests {
		assert.Equal(t, Foo(ft.in), ft.out)
	}
}
