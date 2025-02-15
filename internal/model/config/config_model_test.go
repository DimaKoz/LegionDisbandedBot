package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmptyLegionBotConfig(t *testing.T) {
	want := LegionBotConfig{} //nolint:exhaustruct
	got := NewEmptyLegionBotConfig()
	assert.Equal(t, want, *got)
}

func TestLegionBotConfigString(t *testing.T) {
	want := "LegionBotConfig {TelegramToken:1, DiscordToken:2, PathWhiteListAA:3, PathTelegramUsers:4}"
	config := LegionBotConfig{
		TelegramToken:     "1",
		DiscordToken:      "2",
		PathWhiteListAA:   "3",
		PathTelegramUsers: "4",
	}
	got := config.String()

	assert.Equal(t, want, got)
}
