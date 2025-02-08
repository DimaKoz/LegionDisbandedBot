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
