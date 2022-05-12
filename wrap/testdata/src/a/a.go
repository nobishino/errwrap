package a

import (
	"fmt"

	"github.com/pkg/errors"
)

func f() error {
	err := fmt.Errorf("not wrapped error")
	return errors.WithMessage(err, "message") // want "wrap it"
}

func g() error {
	return f()
}

func h() (n int) {
	n = 10
	return
}
