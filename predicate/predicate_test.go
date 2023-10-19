package predicate

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPredicate(t *testing.T) {
	a := assert.New(t)

	normalFn := func(str string) bool {
		return true
	}
	errorFn := func(str string) (bool, error) {
		return false, fmt.Errorf(str)
	}
	predicate := Of(normalFn)
	checked := Checked(errorFn)
	unchecked := Unchecked(errorFn)

	ret0 := predicate("ok")
	ret1, err := checked("error")
	ret2 := unchecked("error")

	a.True(ret0)
	a.False(ret1)
	a.Equal("error", err.Error())
	a.False(ret2)
}
