package common

import "errors"

func DefaultErrorFn(err error, defErrFn func() error) error {
	if err != nil {
		return err
	}
	return defErrFn()
}

func DefaultErrorMsg(err error, defMsg string) error {
	return DefaultErrorFn(err, func() error {
		return errors.New(defMsg)
	})
}

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
