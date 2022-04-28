package multierror

import (
	"fmt"
)

func prefixError(err error, prefix string) error {
	return fmt.Errorf("%s %s", prefix, err)
}

func Prefix(err error, prefix string) error {
	if err == nil {
		return nil
	}

	switch err := err.(type) {
	case *Collection:
		if err == nil {
			err = new(Collection)
		}

		for i, e := range err.Errors() {
			err.UpdateError(i, prefixError(e, prefix))
		}

		return err
	default:
		return prefixError(err, prefix)
	}
}
