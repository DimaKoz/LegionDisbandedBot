package repository

import (
	"errors"
	"fmt"
)

var errNoSavedData = errors.New("failed to parse json with error: no data")

// errNotFoundWhiteListUser an error of WhiteListStorage repository.
var errNotFoundWhiteListUser = errors.New("couldn't find the user in the white list")

// errNotFoundTelegramUser an error of TelegramUserStorage repository.
var errNotFoundTelegramUser = errors.New("couldn't find the telegram user in the cache")

// repositoryError wraps error with msg and returns wrapped error.
func repositoryError(err error, msg string) error {
	return fmt.Errorf("%w: %s", err, msg)
}
