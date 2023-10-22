package compose

import (
	"github.com/CharLemAznable/gofn/function"
	"github.com/CharLemAznable/gofn/predicate"
	"github.com/CharLemAznable/gofn/supplier"
)

func ThenApply[T any, R any, V any](functionFn1 function.Function[T, R],
	functionFn2 function.Function[R, V],
	errorFn supplier.Supplier[V]) function.Function[T, V] {
	return func(t T) (V, error) {
		r, err := functionFn1(t)
		if err != nil {
			return errorFn()
		}
		return functionFn2(r)
	}
}

func CheckThenSupply[T any, R any](predicateFn predicate.Predicate[T],
	supplierFn supplier.Supplier[R],
	errorFn supplier.Supplier[R]) function.Function[T, R] {
	return func(t T) (R, error) {
		if !predicateFn.Test(t) {
			return errorFn()
		}
		return supplierFn()
	}
}

func CheckThenApply[T any, R any](predicateFn predicate.Predicate[T],
	functionFn function.Function[T, R],
	errorFn supplier.Supplier[R]) function.Function[T, R] {
	return func(t T) (R, error) {
		if !predicateFn.Test(t) {
			return errorFn()
		}
		return functionFn(t)
	}
}
