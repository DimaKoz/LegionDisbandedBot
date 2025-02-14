package testutils

import (
	"os"
	"path/filepath"
	"strings"
)

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
