package compose

import (
	"github.com/CharLemAznable/gofn/consumer"
	"github.com/CharLemAznable/gofn/errs"
	"github.com/CharLemAznable/gofn/function"
	"github.com/CharLemAznable/gofn/predicate"
	"github.com/CharLemAznable/gofn/runnable"
)

func ApplyThenConsume[T any, R any](functionFn function.Function[T, R],
	consumerFn consumer.Consumer[R]) consumer.Consumer[T] {
	return func(t T) error {
		r, err := functionFn(t)
		if err != nil {
			return err
		}
		return consumerFn(r)
	}
}

func CheckThenRun[T any](predicateFn predicate.Predicate[T],
	runnableFn runnable.Runnable) consumer.Consumer[T] {
	return func(t T) error {
		if b, err := predicateFn(t); err != nil || !b {
			return errs.DefaultError(err, "predicate failed")
		}
		return runnableFn()
	}
}

func CheckThenConsume[T any](predicateFn predicate.Predicate[T],
	consumerFn consumer.Consumer[T]) consumer.Consumer[T] {
	return func(t T) error {
		if b, err := predicateFn(t); err != nil || !b {
			return errs.DefaultError(err, "predicate failed")
		}
		return consumerFn(t)
	}
}
