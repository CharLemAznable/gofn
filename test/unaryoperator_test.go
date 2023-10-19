package test

import (
	"fmt"
	"github.com/CharLemAznable/gofn/unaryoperator"
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
	unaryOp := unaryoperator.Of(normalFn)
	checked := unaryoperator.Checked(errorFn)
	unchecked := unaryoperator.Unchecked(errorFn)

	ret0 := unaryOp("ok")
	ret1, err := checked("error")
	ret2 := unchecked("error")

	a.Equal("ok", ret0)
	a.Equal("", ret1)
	a.Equal("error", err.Error())
	a.Equal("", ret2)
}
