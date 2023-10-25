package combinate

import "github.com/CharLemAznable/gofn/common"

func NewContext(value interface{}) common.Context {
	return &context{value: value}
}

type context struct {
	value     interface{}
	error     error
	interrupt bool
}

func (ctx *context) Get() interface{} {
	return ctx.value
}

func (ctx *context) Set(value interface{}) {
	ctx.value = value
}

func (ctx *context) GetErr() error {
	return ctx.error
}

func (ctx *context) SetErr(err error) {
	ctx.error = err
}

func (ctx *context) Interrupted() bool {
	return ctx.interrupt
}

func (ctx *context) SetInterrupt(interrupt bool) {
	ctx.interrupt = interrupt
}
