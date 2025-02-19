package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"sync"

	"github.com/DimaKoz/LegionDisbandedBot/internal/model/user"
	"go.uber.org/zap"
)

var (
	tgUserStorageSync = &sync.Mutex{}
	tgUserStorage     = TelegramUserStorage{
		storage: make(map[string]user.CachedTelegramUser),
	}
)

// TelegramUserStorage represents a storage of user.CachedTelegramUser.
type TelegramUserStorage struct {
	storage map[string]user.CachedTelegramUser
}

// AddTelegramUser adds user.CachedTelegramUser to TelegramUserStorage repository.
func AddTelegramUser(key string, user *user.CachedTelegramUser) {
	tgUserStorageSync.Lock()
	defer tgUserStorageSync.Unlock()
	if user == nil {
		delete(tgUserStorage.storage, key)

		return
	}
	addTelegramUserImpl(key, *user)
}

func addTelegramUserImpl(key string, user user.CachedTelegramUser) {
	tgUserStorage.storage[key] = user
}

// GetTelegramUser returns a *user.CachedTelegramUser if found or error otherwise.
func GetTelegramUser(key string) (*user.CachedTelegramUser, error) {
	tgUserStorageSync.Lock()
	defer tgUserStorageSync.Unlock()

	if found, ok := tgUserStorage.storage[key]; ok {
		return &found, nil
	}

	return nil, repositoryError(errNotFoundTelegramUser, key)
}

func LoadCachedTelegramUser(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("TelegramUserStorage: can't read '%s' file with error: %w", filepath, err)
	}
	defer func(file *os.File) {
		if err = file.Close(); err != nil {
			zap.S().Warnf("TelegramUserStorage: can't close the file: %s with err: %s \n", filepath, err)
		}
	}(file)
	var tgUsers []user.CachedTelegramUser

	const bufferSize = 128
	tgUserReader := bufio.NewReaderSize(file, bufferSize)
	if err = json.NewDecoder(tgUserReader).Decode(&tgUsers); err != nil {
		return fmt.Errorf("TelegramUserStorage: failed to parse json with error: %w", err)
	}

	tgUserStorageSync.Lock()
	defer tgUserStorageSync.Unlock()
	for _, v := range tgUsers {
		tgUserStorage.storage[v.Nickname] = v
	}

	zap.S().Infof("TelegramUserStorage repository: loaded: %d \n", len(tgUsers))

	return nil
}

func getCachedTelegramUsers() []user.CachedTelegramUser {
	result := make([]user.CachedTelegramUser, 0)

	tgUserStorageSync.Lock()
	defer tgUserStorageSync.Unlock()

	for _, v := range tgUserStorage.storage {
		result = append(result, v.Clone())
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})

	return result
}

func SaveCachedTelegramUsers(filepath string) error {
	tgUsers := getCachedTelegramUsers()

	var (
		savingJSON []byte
		err        error
	)
	if savingJSON, err = json.Marshal(tgUsers); err != nil {
		return fmt.Errorf("can't marshal json with error: %w", err)
	}
	var perm os.FileMode = 0o666
	if err = os.WriteFile(filepath, savingJSON, perm); err != nil {
		return fmt.Errorf("can't write '%s' file with error: %w", filepath, err)
	}

	zap.S().Infof("'TelegramUserStorage' repository: saved: %d \n", len(tgUsers))

	return nil
}
