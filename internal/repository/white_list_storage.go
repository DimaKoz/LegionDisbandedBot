package repository

import (
	"sync"

	"github.com/DimaKoz/LegionDisbandedBot/internal/model/user"
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
	addHsImpl(key, *user)
}

func addHsImpl(key string, hs user.WhiteListUser) {
	wlStorage.storage[key] = hs
}

// GetWhitListUser returns a *user.WhiteListUser if found or error otherwise.
func GetWhitListUser(key string) (*user.WhiteListUser, error) {
	wlStorageSync.Lock()
	defer wlStorageSync.Unlock()

	if found, ok := wlStorage.storage[key]; ok {
		return &found, nil
	}

	return nil, repositoryError(errNotFoundWhiteListUser, key)
}
