package test

import (
	"fmt"
	consumer2 "github.com/CharLemAznable/gofn/consumer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConsumer(t *testing.T) {
	a := assert.New(t)
	s := ""

	normalFn := func(str string) {
		s = s + str
	}
	errorFn := func(str string) error {
		return fmt.Errorf(str)
	}
	consumer := consumer2.Of(normalFn)
	checked := consumer2.Checked(errorFn)
	unchecked := consumer2.Unchecked(errorFn)

	consumer("ok")
	err := checked("error")
	unchecked("error")

	a.Equal("ok", s)
	a.Equal("error", err.Error())
}
