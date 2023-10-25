package supplier_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/combinate"
	"github.com/CharLemAznable/gofn/supplier"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {
	// Test case 1
	fn := func() (int, error) {
		return 10, nil
	}
	s := supplier.Of(fn)
	result, err := s()
	assert.NoError(t, err)
	assert.Equal(t, 10, result)

	// Test case 2
	fn = func() (int, error) {
		return 0, nil
	}
	s = supplier.Of(fn)
	result, err = s()
	assert.NoError(t, err)
	assert.Equal(t, 0, result)

	// Test case 3
	fn = func() (int, error) {
		return 0, errors.New("error")
	}
	s = supplier.Of(fn)
	result, err = s()
	assert.Error(t, err)
	assert.Zero(t, result)
}

func TestCast(t *testing.T) {
	// Test case 1
	fn := func() int {
		return 10
	}
	s := supplier.Cast(fn)
	result, err := s()
	assert.NoError(t, err)
	assert.Equal(t, 10, result)

	// Test case 2
	fn = func() int {
		return 0
	}
	s = supplier.Cast(fn)
	result, err = s()
	assert.NoError(t, err)
	assert.Equal(t, 0, result)
}

func TestSupplier_Get(t *testing.T) {
	// Test case 1
	fn := func() (int, error) {
		return 10, nil
	}
	s := supplier.Of(fn)
	result := s.Get()
	assert.Equal(t, 10, result)

	// Test case 2
	fn = func() (int, error) {
		return 0, nil
	}
	s = supplier.Of(fn)
	result = s.Get()
	assert.Equal(t, 0, result)

	// Test case 3
	fn = func() (int, error) {
		return 0, errors.New("error")
	}
	s = supplier.Of(fn)
	result = s.Get()
	assert.Zero(t, result)
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
	assert.Equal(t, "notOk", result)
	assert.Equal(t, err, e)
}
