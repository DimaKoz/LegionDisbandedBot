package internal

import (
	"flag"
	"io"
	"os"
	"testing"

	"github.com/DimaKoz/LegionDisbandedBot/internal/testutils"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

func TestLegionBotFinishedOk(t *testing.T) {
	want := "exiting"

	logger, logs := setupLogsCapture(zapcore.DebugLevel)
	prevL := zap.L()
	defer zap.ReplaceGlobals(prevL)
	zap.ReplaceGlobals(logger)

	testutils.EnvArgsInitConfig(t, testutils.LegionBotTelegramToken, "1")
	testutils.EnvArgsInitConfig(t, testutils.LegionBotDiscordToken, "2")
	testutils.EnvArgsInitConfig(t, testutils.LegionBotWhiteListAa, testutils.GetWD()+"/test/testdata/white_users.json")
	testutils.EnvArgsInitConfig(t, testutils.LegionBotTelegramUsers, "4")

	StartLegionBot(logger)
	assert.NotEmpty(t, logs, "No logs")

	entry := logs.All()[len(logs.All())-1]
	assert.Containsf(t, entry.Message, want, "Invalid log entry %v", entry)
}

func TestLegionBotFinishedByWhiteListUsersRepoError(t *testing.T) {
	want := "LoadWhiteListUser() failed by error:"

	logger, logs := setupLogsCapture(zapcore.DebugLevel)
	prevL := zap.L()
	defer zap.ReplaceGlobals(prevL)
	zap.ReplaceGlobals(logger)

	testutils.EnvArgsInitConfig(t, testutils.LegionBotTelegramToken, "1")
	testutils.EnvArgsInitConfig(t, testutils.LegionBotDiscordToken, "2")
	testutils.EnvArgsInitConfig(t, testutils.LegionBotWhiteListAa, testutils.GetWD()+"/test/testdata/white_users_bad.json")
	testutils.EnvArgsInitConfig(t, testutils.LegionBotTelegramUsers, "4")

	StartLegionBot(logger)
	assert.NotEmpty(t, logs, "No logs")

	entry := logs.All()[len(logs.All())-1]
	assert.Containsf(t, entry.Message, want, "Invalid log entry %v", entry)
}

func TestLegionBotFinishedByConfigErrors(t *testing.T) {
	testutils.EnvArgsUnset(t, testutils.LegionBotTelegramToken)
	testutils.EnvArgsUnset(t, testutils.LegionBotDiscordToken)
	testutils.EnvArgsUnset(t, testutils.LegionBotWhiteListAa)
	testutils.EnvArgsUnset(t, testutils.LegionBotTelegramUsers)

	osArgOrig := os.Args
	t.Cleanup(func() { os.Args = osArgOrig })
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.CommandLine.SetOutput(io.Discard)
	os.Args = make([]string, 0)
	os.Args = append(os.Args, osArgOrig[0])

	want := "failed by error"

	logger, logs := setupLogsCapture(zapcore.DebugLevel)
	prevL := zap.L()
	defer zap.ReplaceGlobals(prevL)
	zap.ReplaceGlobals(logger)
	StartLegionBot(logger)
	assert.NotEmpty(t, logs, "No logs")

	entry := logs.All()[len(logs.All())-1]
	assert.Containsf(t, entry.Message, want, "Invalid log entry %v", entry)
}

func setupLogsCapture(enab zapcore.LevelEnabler) (*zap.Logger, *observer.ObservedLogs) {
	core, logs := observer.New(enab)

	return zap.New(core), logs
}
