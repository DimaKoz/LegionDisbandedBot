package repository

import (
	"sync"

	"github.com/DimaKoz/LegionDisbandedBot/internal/model/user"
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
