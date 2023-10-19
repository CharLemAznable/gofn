package runnable

import (
	"fmt"
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
	runnable := Of(normalFn)
	checked := Checked(errorFn)
	unchecked := Unchecked(errorFn)

	runnable()
	_ = checked()
	unchecked()

	a.Equal(1, i)
}
