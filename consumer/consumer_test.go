package consumer_test

import (
	"github.com/CharLemAznable/gofn/consumer"
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
