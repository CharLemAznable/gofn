package predicate_test

import (
	"fmt"
	"github.com/CharLemAznable/gofn/combinate"
	"github.com/CharLemAznable/gofn/predicate"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {
	// Test case 1: Test with a function that returns true
	fn := func(n int) (bool, error) {
		return n > 0, nil
	}
	p := predicate.Of(fn)
	result, err := p(5)
	assert.True(t, result)
	assert.NoError(t, err)

	// Test case 2: Test with a function that returns false
	fn = func(n int) (bool, error) {
		return n < 0, nil
	}
	p = predicate.Of(fn)
	result, err = p(-5)
	assert.True(t, result)
	assert.NoError(t, err)
}

func TestCast(t *testing.T) {
	// Test case 1: Test with a function that returns true
	fn := func(n int) bool {
		return n > 0
	}
	p := predicate.Cast(fn)
	result, err := p(5)
	assert.True(t, result)
	assert.NoError(t, err)

	// Test case 2: Test with a function that returns false
	fn = func(n int) bool {
		return n < 0
	}
	p = predicate.Cast(fn)
	result, err = p(-5)
	assert.True(t, result)
	assert.NoError(t, err)
}

func TestPredicate_Test(t *testing.T) {
	// Test case 1: Test with a function that returns true
	fn := func(n int) (bool, error) {
		return n > 0, nil
	}
	p := predicate.Of(fn)
	result := p.Test(5)
	assert.True(t, result)

	// Test case 2: Test with a function that returns false
	fn = func(n int) (bool, error) {
		return n < 0, nil
	}
	p = predicate.Of(fn)
	result = p.Test(-5)
	assert.True(t, result)
}

func TestExecute(t *testing.T) {
	// 创建一个可执行对象
	fn := predicate.Cast(func(str string) bool {
		return false
	})

	// 创建一个上下文对象
	ctx := combinate.NewContext("ok")
	// 调用 Execute 方法
	fn.Execute(ctx)
	// 验证结果是否符合预期
	result, err, interrupt := ctx.Get(), ctx.GetErr(), ctx.Interrupted()
	assert.Nil(t, result)
	assert.NoError(t, err)
	assert.True(t, interrupt)

	ctx = combinate.NewContext(0)
	fn.Execute(ctx)
	result, err, interrupt = ctx.Get(), ctx.GetErr(), ctx.Interrupted()
	assert.Nil(t, result)
	assert.Equal(t, fmt.Sprintf(
		"%#v type mismatch %T", 0, ""), err.Error())
	assert.False(t, interrupt)
}
