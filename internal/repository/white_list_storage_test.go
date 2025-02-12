package repository

import (
	"testing"

	"github.com/DimaKoz/LegionDisbandedBot/internal/model/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddGetHs(t *testing.T) {
	type args struct {
		key    string
		wlUser *user.WhiteListUser
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{
				key: "telegramNickname",
				wlUser: &user.WhiteListUser{
					TelegramNickname: "telegramNickname", GameNickname: "gameNickname",
					CorpTicker: "corpTicker", AllyTicker: "allyTicker",
					DiscordGroups: []string{"discordGroups"}, IsSender: true,
				},
			},
		},
	}
	for _, tCase := range tests {
		tUnit := tCase
		t.Run(tUnit.name, func(t *testing.T) {
			AddWhiteListUser(tUnit.args.key, nil)
			got, err := GetWhiteListUser(tUnit.args.key)
			require.Error(t, err)
			assert.Nil(t, got)
			AddWhiteListUser(tUnit.args.key, tUnit.args.wlUser)
			got, err = GetWhiteListUser(tUnit.args.key)
			require.NoError(t, err)
			assert.NotNil(t, got)
			assert.Equal(t, *tUnit.args.wlUser, *got)
		})
	}
}
