package function

import (
	"fmt"
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
	function := Of(normalFn)
	checked := Checked(errorFn)
	unchecked := Unchecked(errorFn)

	ret0 := function("ok")
	ret1, err := checked("error")
	ret2 := unchecked("error")

	a.Equal("ok", ret0)
	a.Equal("", ret1)
	a.Equal("error", err.Error())
	a.Equal("", ret2)
}
