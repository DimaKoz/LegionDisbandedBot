package configer

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/DimaKoz/LegionDisbandedBot/internal/model/config"
	"github.com/DimaKoz/LegionDisbandedBot/internal/testutils"
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

type argTestConfig struct {
	flagT    string
	flagD    string
	flagU    string
	flagM    string
	envFlagT string
	envFlagD string
	envFlagU string
	envFlagM string
}

func TestLoadLegionBotConfig(t *testing.T) {
	tests := []struct {
		name    string
		args    argTestConfig
		want    *config.LegionBotConfig
		wantErr bool
	}{
		{
			name:    "PathEmpty",
			args:    argTestConfig{flagT: "1", flagD: "2", flagU: ""}, //nolint:exhaustruct
			want:    nil,
			wantErr: true,
		},
		{
			name:    "Ok",
			args:    argTestConfig{flagT: "1", flagD: "2", flagU: "3", flagM: "4"}, //nolint:exhaustruct
			want:    &config.LegionBotConfig{TelegramToken: "1", DiscordToken: "2", PathWhiteListAA: "3"},
			wantErr: false,
		},
		{
			name:    "TelegramTokenEmpty",
			args:    argTestConfig{flagT: "", flagD: "2", flagU: "3"}, //nolint:exhaustruct
			want:    nil,
			wantErr: true,
		},
		{
			name:    "DiscordTokenEmpty",
			args:    argTestConfig{flagT: "1", flagD: "", flagU: "3"}, //nolint:exhaustruct
			want:    nil,
			wantErr: true,
		},
		{
			name:    "PathTelegramUsersEmpty",
			args:    argTestConfig{flagT: "1", flagD: "2", flagU: "3"}, //nolint:exhaustruct
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
			os.Args = utils.AppendArgs(os.Args, "-"+flagNameWhiteListAA, tCase.args.flagU)
			os.Args = utils.AppendArgs(os.Args, "-"+flagNamePathTelegramUsers, tCase.args.flagM)

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
	tests := []struct {
		name string
		args argTestConfig
		want *config.LegionBotConfig
	}{
		{
			name: "OnlyEnv",
			args: argTestConfig{ //nolint:exhaustruct
				flagT: "", flagD: "", flagU: "", envFlagT: "1", envFlagD: "2", envFlagU: "3", envFlagM: "4",
			},
			want: &config.LegionBotConfig{TelegramToken: "1", DiscordToken: "2", PathWhiteListAA: "3", PathTelegramUsers: "4"},
		},
		{
			name: "Env+CommandLine",
			args: argTestConfig{
				flagT: "11", flagD: "22", flagU: "33", flagM: "44", envFlagT: "1", envFlagD: "2", envFlagU: "3", envFlagM: "4",
			},
			want: &config.LegionBotConfig{TelegramToken: "11", DiscordToken: "22", PathWhiteListAA: "33"}, //nolint:exhaustruct
		},
	}
	for _, tCase := range tests {
		t.Run(tCase.name, func(t *testing.T) {
			testutils.EnvArgsInitConfig(t, testutils.LegionBotTelegramToken, tCase.args.envFlagT)
			testutils.EnvArgsInitConfig(t, testutils.LegionBotDiscordToken, tCase.args.envFlagD)
			testutils.EnvArgsInitConfig(t, testutils.LegionBotWhiteListAa, tCase.args.envFlagU)
			testutils.EnvArgsInitConfig(t, testutils.LegionBotTelegramUsers, tCase.args.envFlagM)
			osArgOrig := os.Args
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
			flag.CommandLine.SetOutput(io.Discard)
			os.Args = make([]string, 0)
			os.Args = append(os.Args, osArgOrig[0])
			if tCase.args.flagT != "" {
				os.Args = utils.AppendArgs(os.Args, "-"+flagNameTelegramToken, tCase.args.flagT)
			}
			if tCase.args.flagD != "" {
				os.Args = utils.AppendArgs(os.Args, "-"+flagNameDiscordToken, tCase.args.flagD)
			}
			if tCase.args.flagU != "" {
				os.Args = utils.AppendArgs(os.Args, "-"+flagNameWhiteListAA, tCase.args.flagU)
			}
			if tCase.args.flagM != "" {
				os.Args = utils.AppendArgs(os.Args, "-"+flagNamePathTelegramUsers, tCase.args.flagM)
			}

			t.Cleanup(func() { os.Args = osArgOrig })
			got, err := LoadLegionBotConfig()

			require.NoErrorf(t, err, "Configs - got error: %v", err)
			assert.NotNil(t, got)
		})
	}
}
