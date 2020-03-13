package kerr

import (
	"fmt"

	"github.com/cockroachdb/errors"
)

type Kind string

var SecureKind Kind = "secure"

func E(args ...interface{}) error {
	e := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Kind:
			e.Kind = arg
		case error:
			e.error = arg
		default:
			panic("bad call to E")
		}
	}

	return e
}

type Error struct {
	Kind  Kind // category of errors
	error      // the wrapped error
}

func (e *Error) Unwrap() error {
	return e.error
}

func (e *Error) Error() string {
	return fmt.Sprintf("kerr: %s: %v", e.Kind, e.error)
}

func (e *Error) Format(s fmt.State, verb rune) { errors.FormatError(e, s, verb) }

func IsSecureKind(err error) bool {
	var e *Error
	if errors.As(err, &e) {
		return e.Kind == SecureKind
	}

	return false
}
