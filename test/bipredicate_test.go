package test

import (
	"fmt"
	"github.com/CharLemAznable/gofn/bipredicate"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBiPredicate(t *testing.T) {
	a := assert.New(t)

	normalFn := func(str1, str2 string) bool {
		return true
	}
	errorFn := func(str1, str2 string) (bool, error) {
		return false, fmt.Errorf(str1 + str2)
	}
	predicate := bipredicate.Of(normalFn)
	checked := bipredicate.Checked(errorFn)
	unchecked := bipredicate.Unchecked(errorFn)

	ret0 := predicate("ok", "k")
	ret1, err := checked("error", " fail")
	ret2 := unchecked("error", " fail")

	a.True(ret0)
	a.False(ret1)
	a.Equal("error fail", err.Error())
	a.False(ret2)
}
