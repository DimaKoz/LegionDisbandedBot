package internal

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestLegionBotFinished(t *testing.T) {
	want := "exiting"

	logger, logs := setupLogsCapture(zapcore.DebugLevel)
	prevL := zap.L()
	defer zap.ReplaceGlobals(prevL)
	zap.ReplaceGlobals(logger)
	StartLegionBot(logger)
	assert.NotEmpty(t, logs, "No logs")

	entry := logs.All()[len(logs.All())-1]
	assert.Containsf(t, entry.Message, want, "Invalid log entry %v", entry)

	t.Cleanup(func() {
		entries, err := os.ReadDir("./")
		require.NoError(t, err)

		for _, e := range entries {
			if !e.IsDir() && strings.Contains(e.Name(), "report1") {
				err = os.Remove(e.Name())
				require.NoError(t, err)
			}
		}
	})
}

func setupLogsCapture(enab zapcore.LevelEnabler) (*zap.Logger, *observer.ObservedLogs) {
	core, logs := observer.New(enab)

	return zap.New(core), logs
}
