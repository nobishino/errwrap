package a

import (
	"fmt"

	"github.com/pkg/errors"
)

func f() error {
	return fmt.Errorf("error")
}

func g(flag bool) error {
	err := f()
	if flag {
		return err
	}
	return errors.WithMessage(err, "wrap")
}
