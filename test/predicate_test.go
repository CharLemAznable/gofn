package test

import (
	"fmt"
	predicate2 "github.com/CharLemAznable/gofn/predicate"
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
	predicate := predicate2.Of(normalFn)
	checked := predicate2.Checked(errorFn)
	unchecked := predicate2.Unchecked(errorFn)

	ret0 := predicate("ok")
	ret1, err := checked("error")
	ret2 := unchecked("error")

	a.True(ret0)
	a.False(ret1)
	a.Equal("error", err.Error())
	a.False(ret2)
}
