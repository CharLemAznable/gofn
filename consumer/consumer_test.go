package consumer_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/consumer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOf(t *testing.T) {
	fn := func(t int) error {
		// test implementation
		return nil
	}

	c := consumer.Of(fn)

	err := c(10)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestCast(t *testing.T) {
	fn := func(t int) {
		// test implementation
	}

	c := consumer.Cast(fn)

	err := c(10)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestConsumerFn(t *testing.T) {
	fn := func(t int) error {
		// test implementation
		return nil
	}

	c := consumer.Of(fn)

	c.Fn(10)
}

func TestConsumerAccept(t *testing.T) {
	fn := func(t int) error {
		// test implementation
		return nil
	}

	c := consumer.Of(fn)

	c.Accept(10)
}

func TestConsumerFrom(t *testing.T) {
	e := errors.New("error")

	supplierFn := func() (int, error) { return 42, nil }
	consumerFn := consumer.Of(func(t int) error { return nil })

	fn := consumerFn.From(supplierFn)
	err := fn()

	assert.NoError(t, err)

	supplierFn = func() (int, error) { return 0, e }
	fn = consumerFn.From(supplierFn)
	err = fn()

	assert.Equal(t, e, err)
}
