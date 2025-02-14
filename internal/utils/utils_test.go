package utils

import (
	"testing"

	"github.com/DimaKoz/LegionDisbandedBot/internal/testutils"
	"github.com/stretchr/testify/assert"
)

func TestAppendArgs(t *testing.T) {
	args := make([]string, 0)
	want := []string{"a", "b"}
	got := AppendArgs(args, "a", "b")
	assert.Equal(t, want, got)
}

func TestReadFile(t *testing.T) {
	data, err := ReadFile(testutils.GetWD() + "/test/testdata/white_users.json")
	assert.NoError(t, err)
	assert.NotNil(t, data)
}

func TestReadFileNoPath(t *testing.T) {
	b, err := ReadFile("")
	assert.Nil(t, b)
	assert.Error(t, err)
}
