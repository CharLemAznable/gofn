package compose

import (
	"github.com/CharLemAznable/gofn/errs"
	"github.com/CharLemAznable/gofn/function"
	"github.com/CharLemAznable/gofn/predicate"
	"github.com/CharLemAznable/gofn/supplier"
)

func ThenApply[T any, R any, V any](functionFn1 function.Function[T, R], functionFn2 function.Function[R, V], errVal V) function.Function[T, V] {
	return func(t T) (V, error) {
		if r, err := functionFn1(t); err != nil {
			return errVal, err
		} else {
			return functionFn2(r)
		}
	}
}

func TestThenSupply[T any, R any](predicateFn predicate.Predicate[T], supplierFn supplier.Supplier[R], errVal R) function.Function[T, R] {
	return func(t T) (R, error) {
		if b, err := predicateFn(t); err != nil || !b {
			return errVal, errs.DefaultError(err, "predicate failed")
		}
		return supplierFn()
	}
}
