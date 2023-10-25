package biconsumer_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/biconsumer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {
	fn := func(i int, s string) error {
		return nil
	}

	con := biconsumer.Of(fn)
	err := con(10, "test")

	assert.NoError(t, err)
}

func TestCast(t *testing.T) {
	fn := func(i int, s string) {
		// do something
	}

	con := biconsumer.Cast(fn)
	err := con(10, "test")

	assert.NoError(t, err)
}

func TestBiConsumer_Accept(t *testing.T) {
	fn := func(i int, s string) error {
		return errors.New("error")
	}

	con := biconsumer.Of(fn)
	err := biconsumer.Cast(con.Accept)(10, "test")

	assert.NoError(t, err)
}
