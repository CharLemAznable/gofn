package bifunction

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBiFunction(t *testing.T) {
	a := assert.New(t)

	normalFn := func(str1, str2 string) string {
		return str1 + str2
	}
	errorFn := func(str1, str2 string) (string, error) {
		return "", fmt.Errorf(str1 + str2)
	}
	function := Of(normalFn)
	checked := Checked(errorFn)
	unchecked := Unchecked(errorFn)

	ret0 := function("ok", "k")
	ret1, err := checked("error", " fail")
	ret2 := unchecked("error", " fail")

	a.Equal("okk", ret0)
	a.Equal("", ret1)
	a.Equal("error fail", err.Error())
	a.Equal("", ret2)
}
