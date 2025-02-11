package utils

import (
	"fmt"
	"io"
	"os"
)

func AppendArgs(target []string, key string, value string) []string {
	if value != "" {
		target = append(target, key)
		target = append(target, value)
	}

	return target
}

func WrapError(err error, msg string) error {
	return fmt.Errorf("%s: %w ", msg, err)
}

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, WrapError(err, "failed to read file")
	}
	defer func() {
		_ = file.Close()
	}()
	result, err := io.ReadAll(file)
	if err != nil {
		return nil, WrapError(err, "failed to read file")
	}

	return result, nil
}
