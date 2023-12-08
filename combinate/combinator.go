package combinate

import (
	"github.com/CharLemAznable/ge"
	"github.com/CharLemAznable/gofn/common"
	"github.com/CharLemAznable/gofn/consumer"
	"github.com/CharLemAznable/gofn/function"
	"github.com/CharLemAznable/gofn/predicate"
	"github.com/CharLemAznable/gofn/runnable"
	"github.com/CharLemAznable/gofn/supplier"
)

type Combinator interface {
	Append(...common.Executable) Combinator
	Execute(common.Context) common.Context
}

func Runnable(combinator Combinator) runnable.Runnable {
	return func() error {
		ctx := combinator.Execute(NewContext(nil))
		return ctx.GetErr()
	}
}

func Supplier[T any](combinator Combinator) supplier.Supplier[T] {
	return func() (T, error) {
		ctx := combinator.Execute(NewContext(nil))
		t, e := ge.CastOrZero[T](ctx.Get())
		return t, ge.DefaultErrorFn(e, ctx.GetErr)
	}
}

func Consumer[T any](combinator Combinator) consumer.Consumer[T] {
	return func(t T) error {
		ctx := combinator.Execute(NewContext(t))
		return ctx.GetErr()
	}
}

func Function[T any, R any](combinator Combinator) function.Function[T, R] {
	return func(t T) (R, error) {
		ctx := combinator.Execute(NewContext(t))
		r, e := ge.CastOrZero[R](ctx.Get())
		return r, ge.DefaultErrorFn(e, ctx.GetErr)
	}
}

func Predicate[T any](combinator Combinator) predicate.Predicate[T] {
	return func(t T) (bool, error) {
		ctx := combinator.Execute(NewContext(t))
		return !ctx.Interrupted(), ctx.GetErr()
	}
}

func NewCombinator(executables ...common.Executable) Combinator {
	return &combinator{executables: append(make([]common.Executable, 0), executables...)}
}

type combinator struct {
	executables []common.Executable
}

func (c *combinator) Append(executables ...common.Executable) Combinator {
	c.executables = append(c.executables, executables...)
	return c
}

func (c *combinator) Execute(ctx common.Context) common.Context {
	for _, executable := range c.executables {
		executable.Execute(ctx)
		if ctx.GetErr() != nil || ctx.Interrupted() {
			return ctx
		}
	}
	return ctx
}
