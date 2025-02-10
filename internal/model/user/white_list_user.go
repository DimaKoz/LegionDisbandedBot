package user

type WhiteListUser struct {
	TelegramNickname string   `json:"telegramNickname"`
	GameNickname     string   `json:"gameNickname"`
	CorpTicker       string   `json:"corpTicker"`
	AllyTicker       string   `json:"allyTicker"`
	DiscordGroups    []string `json:"discordGroups"`
	IsSender         bool     `json:"isSender"`
}
