package supplier_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/combinate"
	"github.com/CharLemAznable/gofn/supplier"
	"testing"
)

func TestOf(t *testing.T) {
	// Test case 1
	fn := func() (int, error) {
		return 10, nil
	}
	s := supplier.Of(fn)
	result, err := s()
	if err != nil {
		t.Error("Expected no error, but got:", err)
	}
	if result != 10 {
		t.Error("Expected 10, but got:", result)
	}

	// Test case 2
	fn = func() (int, error) {
		return 0, nil
	}
	s = supplier.Of(fn)
	result, err = s()
	if err != nil {
		t.Error("Expected no error, but got:", err)
	}
	if result != 0 {
		t.Error("Expected 0, but got:", result)
	}

	// Test case 3
	fn = func() (int, error) {
		return 0, errors.New("error")
	}
	s = supplier.Of(fn)
	result, err = s()
	if err == nil {
		t.Error("Expected error, but got no error")
	}
	if result != 0 {
		t.Error("Expected 0, but got:", result)
	}

}

func TestCast(t *testing.T) {
	// Test case 1
	fn := func() int {
		return 10
	}
	s := supplier.Cast(fn)
	result, err := s()
	if err != nil {
		t.Error("Expected no error, but got:", err)
	}
	if result != 10 {
		t.Error("Expected 10, but got:", result)
	}

	// Test case 2
	fn = func() int {
		return 0
	}
	s = supplier.Cast(fn)
	result, err = s()
	if err != nil {
		t.Error("Expected no error, but got:", err)
	}
	if result != 0 {
		t.Error("Expected 0, but got:", result)
	}
}

func TestSupplier_Get(t *testing.T) {
	// Test case 1
	fn := func() (int, error) {
		return 10, nil
	}
	s := supplier.Of(fn)
	result := s.Get()
	if result != 10 {
		t.Error("Expected 10, but got:", result)
	}

	// Test case 2
	fn = func() (int, error) {
		return 0, nil
	}
	s = supplier.Of(fn)
	result = s.Get()
	if result != 0 {
		t.Error("Expected 0, but got:", result)
	}

	// Test case 3
	fn = func() (int, error) {
		return 0, errors.New("error")
	}
	s = supplier.Of(fn)
	result = s.Get()
	if result != 0 {
		t.Error("Expected 0, but got:", result)
	}

	// Test case 4
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
		s.MustGet()
	}()
}

func TestExecute(t *testing.T) {
	// 创建一个可执行对象
	err := errors.New("error")
	fn := supplier.Of(func() (string, error) {
		return "notOk", err
	})

	// 创建一个上下文对象
	ctx := combinate.NewContext(nil)
	// 调用 Execute 方法
	fn.Execute(ctx)
	// 验证结果是否符合预期
	result, e := ctx.Get(), ctx.GetErr()
	if result != "notOk" {
		t.Error("Expected 'notOk', but got:", result)
	}
	if e != err {
		t.Error("Expected error, but got:", e)
	}
}
