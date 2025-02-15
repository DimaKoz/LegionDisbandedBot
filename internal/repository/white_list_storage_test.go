package repository

import (
	"testing"

	"github.com/DimaKoz/LegionDisbandedBot/internal/model/user"
	"github.com/DimaKoz/LegionDisbandedBot/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddGetWhiteListUser(t *testing.T) {
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

func TestLoadWhiteListUser(t *testing.T) {
	type args struct {
		filepath string
	}
	wdir := testutils.GetWD()
	tests := []struct {
		name            string
		args            args
		isWantErr       bool
		expectedNumbers int
	}{
		{
			name: "empty path",
			args: args{
				filepath: "",
			},
			isWantErr: true, expectedNumbers: 0,
		},
		{
			name: "ok",
			args: args{
				filepath: wdir + "/test/testdata/white_users.json",
			},
			isWantErr: false, expectedNumbers: 3,
		},
		{
			name: "bad json",
			args: args{
				filepath: wdir + "/test/testdata/white_users_bad.json",
			},
			isWantErr: true, expectedNumbers: 0,
		},
		{
			name: "no white users",
			args: args{
				filepath: wdir + "/test/testdata/white_users_zero.json",
			},
			isWantErr: true, expectedNumbers: 0,
		},
	}
	for _, tCase := range tests {
		tUnit := tCase
		t.Run(tUnit.name, func(t *testing.T) {
			wlStorageOrig := wlStorage
			wlStorage = WhiteListStorage{
				storage: make(map[string]user.WhiteListUser),
			}
			t.Cleanup(func() {
				wlStorage = wlStorageOrig
			})
			err := LoadWhiteListUser(tUnit.args.filepath)
			if tCase.isWantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tUnit.expectedNumbers, len(wlStorage.storage), "Wrong numbers of white users")
		})
	}
}

func TestLoadWhiteListUserCheckParsing(t *testing.T) {
	wantNumbers := 1
	wantUser := user.WhiteListUser{
		TelegramNickname: "@Nickname1", GameNickname: "Manager1",
		CorpTicker: "corpTicker1", AllyTicker: "allyTicker1",
		IsSender: true,
		DiscordGroups: []string{
			"group1",
			"group2",
		},
	}
	filepath := testutils.GetWD() + "/test/testdata/white_users_one.json"
	wlStorageOrig := wlStorage
	wlStorage = WhiteListStorage{
		storage: make(map[string]user.WhiteListUser),
	}
	t.Cleanup(func() {
		wlStorage = wlStorageOrig
	})
	err := LoadWhiteListUser(filepath)
	assert.NoError(t, err)

	assert.Equal(t, wantNumbers, len(wlStorage.storage), "Wrong numbers of white users")

	assert.Equal(t, wantUser, wlStorage.storage[wantUser.TelegramNickname])
}
