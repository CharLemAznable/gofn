package runnable_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/combinate"
	"github.com/CharLemAznable/gofn/runnable"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {
	err := errors.New("error")
	fn := func() error {
		return err
	}

	r := runnable.Of(fn)
	err = r()
	assert.Equal(t, err, err)
}

func TestCast(t *testing.T) {
	called := false
	fn := func() {
		called = true
	}

	r := runnable.Cast(fn)
	err := r()
	assert.NoError(t, err)
	assert.True(t, called)
}

func TestRunnable_Run(t *testing.T) {
	err := errors.New("error")
	fn := func() error {
		return err
	}

	r := runnable.Of(fn)
	err = runnable.Cast(r.Run)()
	assert.NoError(t, err)
}

func TestExecute(t *testing.T) {
	// 创建一个可执行对象
	err := errors.New("error")
	fn := runnable.Of(func() error {
		return err
	})

	// 创建一个上下文对象
	ctx := combinate.NewContext("nil")
	// 调用 Execute 方法
	fn.Execute(ctx)
	// 验证结果是否符合预期
	result, e := ctx.Get(), ctx.GetErr()
	assert.Nil(t, result)
	assert.Equal(t, err, e)
}
