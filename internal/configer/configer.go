package configer

import (
	"errors"
	"flag"
	"fmt"

	"github.com/DimaKoz/LegionDisbandedBot/internal/model/config"
	"github.com/caarlos0/env/v6"
)

const (
	flagNameTelegramToken                = "t"
	flagNameDiscordToken                 = "d"
	flagNamePathAllowedTelegramUsersList = "u"
)

var (
	errNoPassedTelegramToken = errors.New("you must pass a telegram token to '-t' flag " +
		"or set 'LEGION_BOT_TELEGRAM_TOKEN' environment variable")
	errNoPassedDiscordToken = errors.New("you must pass a discord token to '-d' flag " +
		"or set 'LEGION_BOT_DISCORD_TOKEN' environment variable")
	errNoPassedUsersList = errors.New("you must pass telegram users list to '-u' flag " +
		"or set 'LEGION_BOT_ALLOWED_USERS' environment variable")
)

// LoadLegionBotConfig returns *LegionBotConfig.
// we use the next order of priority:
// 1 - Command line options
// 2 - Environment vars
// so:
// If you specify an option by using a parameter on the command line,
// it overrides any value from either the corresponding environment variable.
func LoadLegionBotConfig() (*config.LegionBotConfig, error) {
	cfg := config.NewEmptyLegionBotConfig()

	if err := processEnv(cfg); err != nil {
		return nil, fmt.Errorf("cannot process ENV variables: %w", err)
	}

	processLegionBotFlags(cfg)

	if cfg.TelegramToken == "" {
		return nil, errNoPassedTelegramToken
	}

	if cfg.DiscordToken == "" {
		return nil, errNoPassedDiscordToken
	}

	if cfg.PathAllowedTelegramUsersList == "" {
		return nil, errNoPassedUsersList
	}

	return cfg, nil
}

func processLegionBotFlags(cfg *config.LegionBotConfig) {
	flag.CommandLine.ErrorHandling()
	var flagT, flagD, flagU string
	addChecksStringFlag(flagNameTelegramToken, &flagT)
	addChecksStringFlag(flagNameDiscordToken, &flagD)
	addChecksStringFlag(flagNamePathAllowedTelegramUsersList, &flagU)
	flag.Parse()

	if flagT != "" {
		cfg.TelegramToken = flagT
	}
	if flagD != "" {
		cfg.DiscordToken = flagD
	}
	if flagU != "" {
		cfg.PathAllowedTelegramUsersList = flagU
	}
}

func addChecksStringFlag(flagName string, passedVar *string) {
	if flag.Lookup(flagName) == nil {
		flag.StringVar(passedVar, flagName, "", "")
	}
}

var processEnv = func(config *config.LegionBotConfig) error {
	err := env.Parse(config)
	if err != nil {
		return fmt.Errorf("failed to parse an environment, error: %w", err)
	}

	return nil
}
