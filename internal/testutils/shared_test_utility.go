package testutils

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	LegionBotTelegramToken = "LEGION_BOT_TELEGRAM_TOKEN" //nolint:gosec
	LegionBotDiscordToken  = "LEGION_BOT_DISCORD_TOKEN"  //nolint:gosec
	LegionBotWhiteListAa   = "LEGION_BOT_WHITE_LIST_AA"
	LegionBotTelegramUsers = "LEGION_BOT_TELEGRAM_USERS"
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

func EnvArgsInitConfig(t *testing.T, key string, value string) {
	t.Helper()
	if value != "" {
		origValue := os.Getenv(key)
		err := os.Setenv(key, value)
		assert.NoError(t, err)
		t.Cleanup(func() { _ = os.Setenv(key, origValue) })
	}
}
