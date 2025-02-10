package user

type CachedTelegramUser struct {
	ID        int64  `json:"id"`
	Nickname  string `json:"nickname"`
	IsBot     bool   `json:"isBot"`
	IsBanned  bool   `json:"isBanned,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}
