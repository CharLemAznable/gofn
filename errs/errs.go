package errs

import "errors"

func DefaultError(err error, defMsg string) error {
	if err != nil {
		return err
	}
	return errors.New(defMsg)
}
