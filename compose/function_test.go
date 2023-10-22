package compose_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/compose"
	"github.com/CharLemAznable/gofn/predicate"
	"testing"
)

func TestThenApply(t *testing.T) {
	// 定义测试用的函数和供应器
	fn1 := func(t int) (string, error) {
		if t < 0 {
			return "", errors.New("negative number")
		}
		return "positive number", nil
	}
	fn2 := func(s string) (int, error) {
		if s == "positive number" {
			return 1, nil
		}
		return 0, errors.New("unexpected string")
	}
	errFn := func() (int, error) {
		return 0, errors.New("error function called")
	}

	// 创建被测试的函数
	fn := compose.ThenApply(fn1, fn2, errFn)

	// 执行测试用例
	result, err := fn(10)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if result != 1 {
		t.Errorf("expected result 1, got %d", result)
	}

	result, err = fn(-5)
	if err == nil {
		t.Error("expected non-nil error, got nil")
	}
	if result != 0 {
		t.Errorf("expected result 0, got %d", result)
	}
}

func TestCheckThenSupply(t *testing.T) {
	// 定义测试用的断言函数和供应器
	predicateFn := func(t int) bool {
		return t > 0
	}
	supplierFn := func() (string, error) {
		return "positive number", nil
	}
	errFn := func() (string, error) {
		return "", errors.New("error function called")
	}

	// 创建被测试的函数
	fn := compose.CheckThenSupply(predicate.Cast(predicateFn), supplierFn, errFn)

	// 执行测试用例
	result, err := fn(10)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if result != "positive number" {
		t.Errorf("expected result 'positive number', got '%s'", result)
	}

	result, err = fn(-5)
	if err == nil {
		t.Error("expected non-nil error, got nil")
	}
	if result != "" {
		t.Errorf("expected empty result, got '%s'", result)
	}
}

func TestCheckThenApply(t *testing.T) {
	// 定义测试用的断言函数和函数
	predicateFn := func(t int) bool {
		return t > 0
	}
	fn := func(t int) (string, error) {
		if t > 0 {
			return "positive number", nil
		}
		return "", errors.New("negative number")
	}
	errFn := func() (string, error) {
		return "", errors.New("error function called")
	}

	// 创建被测试的函数
	f := compose.CheckThenApply(predicate.Cast(predicateFn), fn, errFn)

	// 执行测试用例
	result, err := f(10)
	if err != nil {
		t.Errorf("expected nil error, got %v", err)
	}
	if result != "positive number" {
		t.Errorf("expected result 'positive number', got '%s'", result)
	}

	result, err = f(-5)
	if err == nil {
		t.Error("expected non-nil error, got nil")
	}
	if result != "" {
		t.Errorf("expected empty result, got '%s'", result)
	}
}
