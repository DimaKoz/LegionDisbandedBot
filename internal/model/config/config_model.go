package config

// LegionBotConfig represents a passed config while the bot started.
type LegionBotConfig struct {
	TelegramToken                string // flag "-t"
	DiscordToken                 string // flag "-d"
	PathAllowedTelegramUsersList string // flag "-u"
}

func NewEmptyLegionBotConfig() *LegionBotConfig {
	return &LegionBotConfig{
		TelegramToken:                "",
		DiscordToken:                 "",
		PathAllowedTelegramUsersList: "",
	}
}
