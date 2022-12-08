package handling

import "errors"

func ErrCheck(err ...error) error {
	for _, errz := range err {
		if errz != nil {
			return errors.New("\n\nAn error has occurred, error below.\n+" + errz.Error() + "\n")
		}
	}
	return nil
}
