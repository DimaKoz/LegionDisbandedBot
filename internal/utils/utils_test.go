package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendArgs(t *testing.T) {
	args := make([]string, 0)
	want := []string{"a", "b"}
	got := AppendArgs(args, "a", "b")
	assert.Equal(t, want, got)
}
