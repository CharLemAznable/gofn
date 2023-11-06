package consumer_test

import (
	"errors"
	"fmt"
	"github.com/CharLemAznable/gofn/combinate"
	"github.com/CharLemAznable/gofn/consumer"
	"testing"
)

func TestOf(t *testing.T) {
	fn := func(i int) error {
		return nil
	}

	con := consumer.Of(fn)
	err := con(10)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestCast(t *testing.T) {
	fn := func(i int) {
		// do something
	}

	con := consumer.Cast(fn)
	err := con(10)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestConsumer_Accept(t *testing.T) {
	fn := func(i int) error {
		return errors.New("error")
	}

	con := consumer.Of(fn)
	err := consumer.Cast(con.Accept)(10)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}

func TestExecute(t *testing.T) {
	// 创建一个可执行对象
	fn := consumer.Of(func(str string) error {
		return errors.New(str)
	})

	// 创建一个上下文对象
	ctx := combinate.NewContext("error")
	// 调用 Execute 方法
	fn.Execute(ctx)
	// 验证结果是否符合预期
	result, err := ctx.Get(), ctx.GetErr()
	if result != nil {
		t.Errorf("Expected nil, but got: %v", result)
	}
	if err.Error() != "error" {
		t.Errorf("Expected 'error', but got: %v", err.Error())
	}

	ctx = combinate.NewContext(0)
	fn.Execute(ctx)
	result, err = ctx.Get(), ctx.GetErr()
	if result != nil {
		t.Errorf("Expected nil, but got: %v", result)
	}
	expectedErr := fmt.Sprintf("%#v type mismatch %T", 0, "")
	if err.Error() != expectedErr {
		t.Errorf("Expected '%s', but got: %v", expectedErr, err.Error())
	}
}
