package compose_test

import (
	"errors"
	"github.com/CharLemAznable/gofn/compose"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSupplyThenApply(t *testing.T) {
	t.Run("supplierFn returns value, functionFn returns value", func(t *testing.T) {
		expectedResult := "result"
		supplierFn := func() (string, error) {
			return "input", nil
		}
		functionFn := func(input string) (string, error) {
			return expectedResult, nil
		}
		errorFn := func() (string, error) {
			return "", nil
		}

		resultFn := compose.SupplyThenApply(supplierFn, functionFn, errorFn)
		result, err := resultFn()

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})

	t.Run("supplierFn returns error", func(t *testing.T) {
		expectedError := errors.New("supplier error")
		supplierFn := func() (string, error) {
			return "", errors.New("ignored")
		}
		functionFn := func(input string) (string, error) {
			return "", nil
		}
		errorFn := func() (string, error) {
			return "", expectedError
		}

		resultFn := compose.SupplyThenApply(supplierFn, functionFn, errorFn)
		_, err := resultFn()

		assert.Equal(t, expectedError, err)
	})

	t.Run("functionFn returns error", func(t *testing.T) {
		expectedError := errors.New("function error")
		supplierFn := func() (string, error) {
			return "input", nil
		}
		functionFn := func(input string) (string, error) {
			return "", expectedError
		}
		errorFn := func() (string, error) {
			return "", nil
		}

		resultFn := compose.SupplyThenApply(supplierFn, functionFn, errorFn)
		_, err := resultFn()

		assert.Equal(t, expectedError, err)
	})

	t.Run("errorFn returns value", func(t *testing.T) {
		expectedResult := "error result"
		supplierFn := func() (string, error) {
			return "", errors.New("supplier error")
		}
		functionFn := func(input string) (string, error) {
			return "", errors.New("function error")
		}
		errorFn := func() (string, error) {
			return expectedResult, nil
		}

		resultFn := compose.SupplyThenApply(supplierFn, functionFn, errorFn)
		result, err := resultFn()

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
	})
}
