package compose

import (
	"github.com/CharLemAznable/gofn/function"
	"github.com/CharLemAznable/gofn/supplier"
)

func SupplyThenApply[T any, R any](supplierFn supplier.Supplier[T],
	functionFn function.Function[T, R]) supplier.Supplier[R] {
	return func() (R, error) {
		t, err := supplierFn()
		if err != nil {
			var n interface{}
			r, _ := n.(R)
			return r, err
		}
		return functionFn(t)
	}
}
