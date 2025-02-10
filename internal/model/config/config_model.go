package config

// LegionBotConfig represents a passed config while the bot started.
type LegionBotConfig struct {
	TelegramToken string `env:"LEGION_BOT_TELEGRAM_TOKEN"` // flag "-t"
	DiscordToken  string `env:"LEGION_BOT_DISCORD_TOKEN"`  // flag "-d"
	// Use it to get the emulated list of telegram users from AA
	PathWhiteListAA string `env:"LEGION_BOT_WHITE_LIST_AA"` // flag "-u"
	// Use this path to get&save the list of telegram users matched by the telegram bot
	PathTelegramUsers string `env:"LEGION_BOT_TELEGRAM_USERS"` // flag "-m"
}

func NewEmptyLegionBotConfig() *LegionBotConfig {
	return &LegionBotConfig{
		TelegramToken:     "",
		DiscordToken:      "",
		PathWhiteListAA:   "",
		PathTelegramUsers: "",
	}
}
