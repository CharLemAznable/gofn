package compose

import (
	"github.com/CharLemAznable/gofn/function"
	"github.com/CharLemAznable/gofn/supplier"
)

func SupplyThenApply[T any, R any](supplierFn supplier.Supplier[T], functionFn function.Function[T, R], errVal R) supplier.Supplier[R] {
	return func() (R, error) {
		if t, err := supplierFn(); err != nil {
			return errVal, err
		} else {
			return functionFn(t)
		}
	}
}
