package unaryoperator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnaryOperator(t *testing.T) {
	a := assert.New(t)

	normalFn := func(str string) string {
		return str
	}
	errorFn := func(str string) (string, error) {
		return "", fmt.Errorf(str)
	}
	unaryOp := Of(normalFn)
	checked := Checked(errorFn)
	unchecked := Unchecked(errorFn)

	ret0 := unaryOp("ok")
	ret1, err := checked("error")
	ret2 := unchecked("error")

	a.Equal("ok", ret0)
	a.Equal("", ret1)
	a.Equal("error", err.Error())
	a.Equal("", ret2)
}
