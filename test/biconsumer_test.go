package test

import (
	"fmt"
	"github.com/CharLemAznable/gofn/biconsumer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBiConsumer(t *testing.T) {
	a := assert.New(t)
	s := ""

	normalFn := func(str1, str2 string) {
		s = s + str1 + str2
	}
	errorFn := func(str1, str2 string) error {
		return fmt.Errorf(str1 + str2)
	}
	consumer := biconsumer.Of(normalFn)
	checked := biconsumer.Checked(errorFn)
	unchecked := biconsumer.Unchecked(errorFn)

	consumer("ok", "k")
	err := checked("error", " fail")
	unchecked("error", " fail")

	a.Equal("okk", s)
	a.Equal("error fail", err.Error())
}
