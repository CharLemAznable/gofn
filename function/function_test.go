package function_test

import (
	"fmt"
	"github.com/CharLemAznable/gofn/combinate"
	"github.com/CharLemAznable/gofn/function"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOf(t *testing.T) {
	// Test case 1: Test with a function that returns a string
	fn := function.Of(func(t string) (string, error) {
		return "Hello, " + t, nil
	})

	result, err := fn("World")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expected := "Hello, World"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	// Test case 2: Test with a function that returns an integer
	fn2 := function.Of(func(t int) (int, error) {
		return t * 2, nil
	})

	result2, err := fn2(5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	expectedInt := 10
	if result2 != expectedInt {
		t.Errorf("Expected %d, but got %d", expectedInt, result2)
	}
}

func TestCast(t *testing.T) {
	// Test case 1: Test with a function that takes a string and returns an integer
	fn := function.Cast(func(t string) int {
		return len(t)
	})

	result := fn.Apply("Hello")
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 2: Test with a function that takes an integer and returns a string
	fn2 := function.Cast(func(t int) string {
		if t%2 == 0 {
			return "even"
		}
		return "odd"
	})

	resultStr := fn2.Apply(7)
	expectedStr := "odd"
	if resultStr != expectedStr {
		t.Errorf("Expected %s, but got %s", expectedStr, resultStr)
	}
}

func TestExecute(t *testing.T) {
	// 创建一个可执行对象
	fn := function.Cast(func(str string) string {
		return "[" + str + "]"
	})

	// 创建一个上下文对象
	ctx := combinate.NewContext("ok")
	// 调用 Execute 方法
	fn.Execute(ctx)
	// 验证结果是否符合预期
	result, err := ctx.Get(), ctx.GetErr()
	assert.Equal(t, "[ok]", result)
	assert.NoError(t, err)

	ctx = combinate.NewContext(0)
	fn.Execute(ctx)
	result, err = ctx.Get(), ctx.GetErr()
	assert.Nil(t, result)
	assert.Equal(t, fmt.Sprintf(
		"%#v type mismatch %T", 0, ""), err.Error())
}
