package repository

import (
	"io"
	"os"
	"strconv"
	"testing"
	"time"

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

func TestSaveCachedTelegramUsers(t *testing.T) {
	type args struct {
		filepath string
		tgUsers  []*user.CachedTelegramUser
	}
	tests := []struct {
		name            string
		args            args
		isWantErr       bool
		expectedContent string
	}{
		{
			name: "empty path", isWantErr: true, expectedContent: "", args: args{filepath: "", tgUsers: nil},
		},
		{
			name: "ok", isWantErr: false,
			args: args{
				filepath: testutils.GetWD() + "/tempfile_cached_users" + strconv.FormatInt(time.Now().Unix(), 10) + ".tmp",
				tgUsers: []*user.CachedTelegramUser{
					{
						ID: 123, Nickname: "@Nickname1", FirstName: "FirstName1",
						LastName: "LastName1", IsBot: true, IsBanned: true,
					},
					{
						ID: 1234, Nickname: "@Nickname2", FirstName: "FirstName2",
						LastName: "LastName2", IsBot: false, IsBanned: false,
					},
				},
			},
			expectedContent: "[{\"id\":123,\"nickname\":\"@Nickname1\",\"isBot\"" +
				":true,\"isBanned\":true,\"firstName\":\"FirstName1\"," +
				"\"lastName\":\"LastName1\"},{\"id\":1234,\"nickname\":\"@Nickname2\"" +
				",\"isBot\":false,\"firstName\":\"FirstName2\",\"lastName\":\"LastName2\"}]",
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
			for _, tgUser := range tCase.args.tgUsers {
				AddTelegramUser(tgUser.Nickname, tgUser)
			}
			err := SaveCachedTelegramUsers(tUnit.args.filepath)
			if tCase.isWantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				file, err := os.Open(tUnit.args.filepath)
				assert.NoError(t, err)
				t.Cleanup(func() { _ = os.Remove(tUnit.args.filepath) })
				defer func() { _ = file.Close() }()
				b, err := io.ReadAll(file)
				assert.NoError(t, err)
				assert.Equal(t, tUnit.expectedContent, string(b))
			}
		})
	}
}
