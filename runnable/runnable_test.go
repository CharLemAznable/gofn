package runnable_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/combinate"
	"github.com/CharLemAznable/gofn/runnable"
	"testing"
)

func TestOf(t *testing.T) {
	err := errors.New("error")
	fn := func() error {
		return err
	}

	r := runnable.Of(fn)
	e := r()
	if err != e {
		t.Errorf("Expected error '%v', but got '%v'", err, e)
	}
}

func TestCast(t *testing.T) {
	called := false
	fn := func() {
		called = true
	}

	r := runnable.Cast(fn)
	err := r()
	if err != nil {
		t.Errorf("Expected error is nil, but got '%v'", err)
	}
	if !called {
		t.Error("Expected called, but not called")
	}
}

func TestRunnable_Run(t *testing.T) {
	err := errors.New("error")
	fn := func() error {
		return err
	}

	r := runnable.Of(fn)
	err = runnable.Cast(r.Run)()
	if err != nil {
		t.Errorf("Expected error is nil, but got '%v'", err)
	}
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
	if result != nil {
		t.Errorf("Expected result is nil, but got '%v'", result)
	}
	if err != e {
		t.Errorf("Expected error '%v', but got '%v'", err, e)
	}
}
