package repository

import (
	"testing"

	"github.com/DimaKoz/LegionDisbandedBot/internal/model/user"
	"github.com/DimaKoz/LegionDisbandedBot/internal/testutils"
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

func TestLoadCachedTelegramUser(t *testing.T) {
	type args struct {
		filepath string
	}
	wdir := testutils.GetWD()
	tests := []struct {
		name             string
		args             args
		isWantErr        bool
		wantTelegramUser *user.CachedTelegramUser
		expectedNumbers  int
	}{
		{
			name:             "empty path",
			args:             args{filepath: ""},
			wantTelegramUser: nil, isWantErr: true, expectedNumbers: 0,
		},
		{
			name: "ok",
			args: args{filepath: wdir + "/test/testdata/cached_users.json"},
			wantTelegramUser: &user.CachedTelegramUser{
				ID: 123, Nickname: "@Nickname1", FirstName: "FirstName1",
				LastName: "LastName1", IsBot: true, IsBanned: true,
			},
			isWantErr: false, expectedNumbers: 2,
		},
		{
			name: "bad json",
			args: args{
				filepath: wdir + "/test/testdata/white_users_bad.json",
			}, wantTelegramUser: nil, isWantErr: true, expectedNumbers: 0,
		},
		{
			name: "no tg users",
			args: args{
				filepath: wdir + "/test/testdata/white_users_zero.json",
			}, wantTelegramUser: nil, isWantErr: true, expectedNumbers: 0,
		},
	}
	for _, tCase := range tests {
		tUnit := tCase
		t.Run(tUnit.name, func(t *testing.T) {
			tgUserStorageOrig := tgUserStorage
			tgUserStorage = TelegramUserStorage{
				storage: make(map[string]user.CachedTelegramUser),
			}
			t.Cleanup(func() { tgUserStorage = tgUserStorageOrig })
			err := LoadCachedTelegramUser(tUnit.args.filepath)
			if tCase.isWantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tUnit.expectedNumbers, len(tgUserStorage.storage), "Wrong numbers of telegram users")
			if tUnit.wantTelegramUser != nil {
				got, err := GetTelegramUser(tUnit.wantTelegramUser.Nickname)
				assert.NoError(t, err)
				assert.Equal(t, *tUnit.wantTelegramUser, *got)
			}
		})
	}
}
