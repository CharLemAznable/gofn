package test

import (
	"fmt"
	function2 "github.com/CharLemAznable/gofn/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunction(t *testing.T) {
	a := assert.New(t)

	normalFn := func(str string) string {
		return str
	}
	errorFn := func(str string) (string, error) {
		return "", fmt.Errorf(str)
	}
	function := function2.Of(normalFn)
	checked := function2.Checked(errorFn)
	unchecked := function2.Unchecked(errorFn)

	ret0 := function("ok")
	ret1, err := checked("error")
	ret2 := unchecked("error")

	a.Equal("ok", ret0)
	a.Equal("", ret1)
	a.Equal("error", err.Error())
	a.Equal("", ret2)
}
