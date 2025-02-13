package repository

import (
	"testing"

	"github.com/DimaKoz/LegionDisbandedBot/internal/model/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddGetGetTelegramUser(t *testing.T) {
	type args struct {
		key    string
		tgUser *user.CachedTelegramUser
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "ok",
			args: args{
				key: "telegramNickname",
				tgUser: &user.CachedTelegramUser{
					ID: 123, Nickname: "telegramNickname", FirstName: "First",
					LastName: "Last", IsBot: true, IsBanned: true,
				},
			},
		},
	}
	for _, tCase := range tests {
		tUnit := tCase
		t.Run(tUnit.name, func(t *testing.T) {
			AddTelegramUser(tUnit.args.key, nil)
			got, err := GetTelegramUser(tUnit.args.key)
			require.Error(t, err)
			assert.Nil(t, got)

			AddTelegramUser(tUnit.args.key, tUnit.args.tgUser)
			got, err = GetTelegramUser(tUnit.args.key)
			require.NoError(t, err)
			assert.NotNil(t, got)
			assert.Equal(t, *tUnit.args.tgUser, *got)

			AddTelegramUser(tUnit.args.key, nil)
			got, err = GetTelegramUser(tUnit.args.key)
			require.Error(t, err)
			assert.Nil(t, got)
		})
	}
}
