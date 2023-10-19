package supplier

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSupplier(t *testing.T) {
	a := assert.New(t)

	normalFn := func() string {
		return "ok"
	}
	errorFn := func() (string, error) {
		return "", fmt.Errorf("error")
	}
	supplier := Of(normalFn)
	checked := Checked(errorFn)
	unchecked := Unchecked(errorFn)

	ret0 := supplier()
	ret1, err := checked()
	ret2 := unchecked()

	a.Equal("ok", ret0)
	a.Equal("", ret1)
	a.Equal("error", err.Error())
	a.Equal("", ret2)
}
