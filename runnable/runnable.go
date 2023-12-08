package runnable

import (
	"github.com/CharLemAznable/ge"
	"github.com/CharLemAznable/gofn/common"
)

type Runnable func() error

func Of(fn func() error) Runnable {
	return fn
}

func Cast(fn func()) Runnable {
	return func() error {
		fn()
		return nil
	}
}

func (fn Runnable) Fn() error {
	return fn()
}

func (fn Runnable) Run() {
	_ = fn()
}

func (fn Runnable) MustRun() {
	ge.PanicIfError(fn())
}

func (fn Runnable) Execute(ctx common.Context) {
	err := fn()
	ctx.SetErr(err)
	if err != nil {
		ctx.Set(nil)
	}
}
