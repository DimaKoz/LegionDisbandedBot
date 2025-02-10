package configer

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/DimaKoz/LegionDisbandedBot/internal/model/config"
	"github.com/DimaKoz/LegionDisbandedBot/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	errAny                    = errors.New("any error")
	errWantTestProcessEnvMock = errors.New("cannot process ENV variables: any error")
)

func TestProcessEnvMock(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	flag.CommandLine.SetOutput(io.Discard)

	osArgOrig := os.Args
	os.Args = make([]string, 0)
	os.Args = append(os.Args, osArgOrig[0])

	t.Cleanup(func() {
		os.Args = osArgOrig
	})
	origProcessEnv := processEnv
	processEnv = func(_ *config.LegionBotConfig) error {
		return errAny
	}
	t.Cleanup(func() { processEnv = origProcessEnv })
	_, gotErr := LoadLegionBotConfig()
	wantErr := errWantTestProcessEnvMock
	assert.NotNil(t, gotErr)
	assert.Equal(t, wantErr.Error(), gotErr.Error(), "Configs - got error: %v, want: %v", gotErr, wantErr)
}

var errTestProcessEnvError = errors.New("env: expected a pointer to a Struct")

func TestProcessEnvError(t *testing.T) {
	wantErr := fmt.Errorf("failed to parse an environment, error: %w", errTestProcessEnvError)
	gotErr := processEnv(nil)

	assert.Equal(t, wantErr, gotErr, "Configs - got error: %v, want: %v", gotErr, wantErr)
}

func TestLoadLegionBotConfig(t *testing.T) {
	type argTestConfig struct {
		flagT string
		flagD string
		flagU string
	}

	tests := []struct {
		name    string
		args    argTestConfig
		want    *config.LegionBotConfig
		wantErr bool
	}{
		{
			name:    "PathEmpty",
			args:    argTestConfig{flagT: "1", flagD: "2", flagU: ""},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Ok",
			args:    argTestConfig{flagT: "1", flagD: "2", flagU: "3"},
			want:    &config.LegionBotConfig{TelegramToken: "1", DiscordToken: "2", PathWhiteListAA: "3"},
			wantErr: false,
		},
		{
			name:    "TelegramTokenEmpty",
			args:    argTestConfig{flagT: "", flagD: "2", flagU: "3"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "DiscordTokenEmpty",
			args:    argTestConfig{flagT: "1", flagD: "", flagU: "3"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			osArgOrig := os.Args
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
			flag.CommandLine.SetOutput(io.Discard)
			os.Args = make([]string, 0)
			os.Args = append(os.Args, osArgOrig[0])
			os.Args = utils.AppendArgs(os.Args, "-"+flagNameTelegramToken, tCase.args.flagT)
			os.Args = utils.AppendArgs(os.Args, "-"+flagNameDiscordToken, tCase.args.flagD)
			os.Args = utils.AppendArgs(os.Args, "-"+flagNamePathAllowedTelegramUsersList, tCase.args.flagU)

			t.Cleanup(func() { os.Args = osArgOrig })
			got, err := LoadLegionBotConfig()
			if tCase.wantErr {
				require.Error(t, err)
				assert.Nil(t, got)
			} else {
				require.NoErrorf(t, err, "Configs - got error: %v", err)
				assert.NotNil(t, got)
			}
		})
	}
}

func TestConfigEnv(t *testing.T) {
	type argTestConfig struct {
		flagT    string
		flagD    string
		flagU    string
		envFlagT string
		envFlagD string
		envFlagU string
	}

	tests := []struct {
		name string
		args argTestConfig
		want *config.LegionBotConfig
	}{
		{
			name: "OnlyEnv",
			args: argTestConfig{flagT: "", flagD: "", flagU: "", envFlagT: "1", envFlagD: "2", envFlagU: "3"},
			want: &config.LegionBotConfig{TelegramToken: "1", DiscordToken: "2", PathWhiteListAA: "3"},
		},
		{
			name: "Env+CommandLine",
			args: argTestConfig{flagT: "11", flagD: "22", flagU: "33", envFlagT: "1", envFlagD: "2", envFlagU: "3"},
			want: &config.LegionBotConfig{TelegramToken: "11", DiscordToken: "22", PathWhiteListAA: "33"},
		},
	}
	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			envArgsInitConfig(t, "LEGION_BOT_TELEGRAM_TOKEN", tCase.args.envFlagT)
			envArgsInitConfig(t, "LEGION_BOT_DISCORD_TOKEN", tCase.args.envFlagD)
			envArgsInitConfig(t, "LEGION_BOT_WHITE_LIST_AA", tCase.args.envFlagU)
			osArgOrig := os.Args
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
			flag.CommandLine.SetOutput(io.Discard)
			os.Args = make([]string, 0)
			os.Args = append(os.Args, osArgOrig[0])
			if tCase.args.envFlagT != "" {
				os.Args = utils.AppendArgs(os.Args, "-"+flagNameTelegramToken, tCase.args.flagT)
			}
			if tCase.args.envFlagD != "" {
				os.Args = utils.AppendArgs(os.Args, "-"+flagNameDiscordToken, tCase.args.flagD)
			}
			if tCase.args.envFlagU != "" {
				os.Args = utils.AppendArgs(os.Args, "-"+flagNamePathAllowedTelegramUsersList, tCase.args.flagU)
			}

			t.Cleanup(func() { os.Args = osArgOrig })
			got, err := LoadLegionBotConfig()

			require.NoErrorf(t, err, "Configs - got error: %v", err)
			assert.NotNil(t, got)
		})
	}
}

func envArgsInitConfig(t *testing.T, key string, value string) {
	t.Helper()
	if value != "" {
		origValue := os.Getenv(key)
		err := os.Setenv(key, value)
		assert.NoError(t, err)
		t.Cleanup(func() { _ = os.Setenv(key, origValue) })
	}
}
