package configer

import (
	"flag"
	"io"
	"os"
	"testing"

	"github.com/DimaKoz/LegionDisbandedBot/internal/model/config"
	"github.com/DimaKoz/LegionDisbandedBot/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadMonitorConfig(t *testing.T) {
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
			want:    &config.LegionBotConfig{TelegramToken: "1", DiscordToken: "2", PathAllowedTelegramUsersList: "3"},
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
