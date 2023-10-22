package compose

import (
	"github.com/CharLemAznable/gofn/consumer"
	"github.com/CharLemAznable/gofn/runnable"
	"github.com/CharLemAznable/gofn/supplier"
)

func RunInSequence(functions ...runnable.Runnable) runnable.Runnable {
	return func() error {
		for _, fn := range functions {
			if err := fn(); err != nil {
				return err
			}
		}
		return nil
	}
}

func SupplyThenConsume[T any](supplierFn supplier.Supplier[T],
	consumerFn consumer.Consumer[T]) runnable.Runnable {
	return func() error {
		t, err := supplierFn()
		if err != nil {
			return err
		}
		return consumerFn(t)
	}
}
