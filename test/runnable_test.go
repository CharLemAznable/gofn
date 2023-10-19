package test

import (
	"fmt"
	runnable2 "github.com/CharLemAznable/gofn/runnable"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunnable(t *testing.T) {
	a := assert.New(t)
	i := 0

	normalFn := func() {
		i += 1
	}
	errorFn := func() error {
		return fmt.Errorf("error")
	}
	runnable := runnable2.Of(normalFn)
	checked := runnable2.Checked(errorFn)
	unchecked := runnable2.Unchecked(errorFn)

	runnable()
	_ = checked()
	unchecked()

	a.Equal(1, i)
}
