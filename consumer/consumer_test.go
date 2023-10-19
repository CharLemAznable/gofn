package consumer

import (
	"fmt"
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
	consumer := Of(normalFn)
	checked := Checked(errorFn)
	unchecked := Unchecked(errorFn)

	consumer("ok")
	err := checked("error")
	unchecked("error")

	a.Equal("ok", s)
	a.Equal("error", err.Error())
}
