package common

import (
	"errors"
	"fmt"
)

func Cast[T any](val interface{}) (T, error) {
	t, ok := val.(T)
	if !ok {
		return t, errors.New(fmt.Sprintf(
			"%#v type mismatch %T", val, t))
	}
	return t, nil
}

func CastOrZero[T any](val interface{}) (T, error) {
	t, err := Cast[T](val)
	if val == nil {
		return t, nil
	} else {
		return t, err
	}
}

func CastQuietly[T any](val interface{}) T {
	t, _ := val.(T)
	return t
}

func Zero[T any]() T {
	return CastQuietly[T](nil)
}
