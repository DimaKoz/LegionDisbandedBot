package user

type CachedTelegramUser struct {
	ID        int64  `json:"id"`
	Nickname  string `json:"nickname"`
	IsBot     bool   `json:"isBot"`
	IsBanned  bool   `json:"isBanned,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

// Clone creates a copy of instance of CachedTelegramUser and returns it.
func (tgUser CachedTelegramUser) Clone() CachedTelegramUser {
	return CachedTelegramUser{
		ID:        tgUser.ID,
		Nickname:  tgUser.Nickname,
		IsBot:     tgUser.IsBot,
		IsBanned:  tgUser.IsBanned,
		FirstName: tgUser.FirstName,
		LastName:  tgUser.LastName,
	}
}
