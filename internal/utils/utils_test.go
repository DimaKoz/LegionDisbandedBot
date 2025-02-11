package utils

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendArgs(t *testing.T) {
	args := make([]string, 0)
	want := []string{"a", "b"}
	got := AppendArgs(args, "a", "b")
	assert.Equal(t, want, got)
}

func TestReadFile(t *testing.T) {
	data, err := ReadFile(GetWD() + "/test/testdata/white_users.json")
	assert.NoError(t, err)
	assert.NotNil(t, data)
}

func TestReadFileNoPath(t *testing.T) {
	b, err := ReadFile("")
	assert.Nil(t, b)
	assert.Error(t, err)
}

var workDir = ""

func GetWD() string {
	if workDir != "" {
		return workDir
	}
	wDirTemp, _ := os.Getwd()

	for !strings.HasSuffix(wDirTemp, "internal") {
		wDirTemp = filepath.Dir(wDirTemp)
	}
	wDirTemp = filepath.Dir(wDirTemp)
	workDir = wDirTemp

	return workDir
}
