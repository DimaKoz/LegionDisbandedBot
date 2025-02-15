package repository

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/DimaKoz/LegionDisbandedBot/internal/model/user"
	"go.uber.org/zap"
)

var (
	wlStorageSync = &sync.Mutex{}
	wlStorage     = WhiteListStorage{
		storage: make(map[string]user.WhiteListUser, 0),
	}
)

// WhiteListStorage represents a storage of user.WhiteListUser.
type WhiteListStorage struct {
	storage map[string]user.WhiteListUser
}

// AddWhiteListUser adds user.WhiteListUser to WhiteListStorage repository.
func AddWhiteListUser(key string, user *user.WhiteListUser) {
	wlStorageSync.Lock()
	defer wlStorageSync.Unlock()
	if user == nil {
		delete(wlStorage.storage, key)

		return
	}
	addWhiteListUserImpl(key, *user)
}

func addWhiteListUserImpl(key string, user user.WhiteListUser) {
	wlStorage.storage[key] = user
}

// GetWhiteListUser returns a *user.WhiteListUser if found or error otherwise.
func GetWhiteListUser(key string) (*user.WhiteListUser, error) {
	wlStorageSync.Lock()
	defer wlStorageSync.Unlock()

	if found, ok := wlStorage.storage[key]; ok {
		return &found, nil
	}

	return nil, repositoryError(errNotFoundWhiteListUser, key)
}

func LoadWhiteListUser(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("WhiteListStorage: can't read '%s' file with error: %w", filepath, err)
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			zap.S().Warnf("WhiteListStorage: can't close the file: %s with err: %s \n", filepath, err)
		}
	}(file)
	var wlUsers []user.WhiteListUser

	const bufferSize = 128
	wlUserReader := bufio.NewReaderSize(file, bufferSize)
	err = json.NewDecoder(wlUserReader).Decode(&wlUsers)
	if err != nil {
		return fmt.Errorf("WhiteListStorage: failed to parse json with error: %w", err)
	}

	if len(wlUsers) == 0 {
		return errNoSavedData
	}

	wlStorageSync.Lock()
	defer wlStorageSync.Unlock()
	for _, v := range wlUsers {
		wlStorage.storage[v.TelegramNickname] = v
	}

	zap.S().Infof("WhiteListStorage repository: loaded: %d \n", len(wlUsers))

	return nil
}
