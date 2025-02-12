package repository

import (
	"errors"
	"fmt"
)

// errNotFoundWhiteListUser an error of WhiteListStorage repository.
var errNotFoundWhiteListUser = errors.New("couldn't find the user in the white list")

// repositoryError wraps error with msg and returns wrapped error.
func repositoryError(err error, msg string) error {
	return fmt.Errorf("%w: %s", err, msg)
}
