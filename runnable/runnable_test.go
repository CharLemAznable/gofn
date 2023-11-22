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
	fn := func() error {
		return errors.New("error")
	}

	r := runnable.Of(fn)
	err := runnable.Cast(r.Run).Fn()
	if err != nil {
		t.Errorf("Expected error is nil, but got '%v'", err)
	}

	func() {
		defer func() {
			rec := recover()
			if rec == nil {
				t.Error("Expected recover error, but got nil")
			}
			recErr, ok := rec.(error)
			if !ok {
				t.Errorf("Expected recover error, but got %v", rec)
			} else if recErr.Error() != "error" {
				t.Errorf("Expected error message 'error', but got '%s'", recErr.Error())
			}
		}()
		r.MustRun()
	}()
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
