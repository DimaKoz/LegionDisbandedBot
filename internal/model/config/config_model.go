package config

// LegionBotConfig represents a passed config while the bot started.
type LegionBotConfig struct {
	TelegramToken   string `env:"LEGION_BOT_TELEGRAM_TOKEN"` // flag "-t"
	DiscordToken    string `env:"LEGION_BOT_DISCORD_TOKEN"`  // flag "-d"
	PathWhiteListAA string `env:"LEGION_BOT_WHITE_LIST_AA"`  // flag "-u"
}

func NewEmptyLegionBotConfig() *LegionBotConfig {
	return &LegionBotConfig{
		TelegramToken:   "",
		DiscordToken:    "",
		PathWhiteListAA: "",
	}
}
