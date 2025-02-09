package configer

import (
	"errors"
	"flag"
	"fmt"

	"github.com/DimaKoz/LegionDisbandedBot/internal/model/config"
)

const (
	flagNameTelegramToken                = "t"
	flagNameDiscordToken                 = "d"
	flagNamePathAllowedTelegramUsersList = "u"
)

var (
	errNoPassedTelegramToken = errors.New("you must pass a telegram token to '-t' flag")
	errNoPassedDiscordToken  = errors.New("you must pass a discord token to '-d' flag")
	errNoPassedUsersList     = errors.New("you must pass telegram users list to '-u' flag")
)

// LoadLegionBotConfig returns *LegionBotConfig.
func LoadLegionBotConfig() (*config.LegionBotConfig, error) {
	cfg := config.NewEmptyLegionBotConfig()
	if err := processLegionBotFlags(cfg); err != nil {
		return nil, fmt.Errorf("cannot process flags variables: %w", err)
	}

	return cfg, nil
}

func processLegionBotFlags(cfg *config.LegionBotConfig) error {
	flag.CommandLine.ErrorHandling()
	var flagT, flagD, flagU string
	addChecksStringFlag(flagNameTelegramToken, &flagT)
	addChecksStringFlag(flagNameDiscordToken, &flagD)
	addChecksStringFlag(flagNamePathAllowedTelegramUsersList, &flagU)
	flag.Parse()

	if flagT == "" {
		return errNoPassedTelegramToken
	}
	cfg.TelegramToken = flagT

	if flagD == "" {
		return errNoPassedDiscordToken
	}
	cfg.DiscordToken = flagD

	if flagU == "" {
		return errNoPassedUsersList
	}
	cfg.PathAllowedTelegramUsersList = flagU

	return nil
}

func addChecksStringFlag(flagName string, passedVar *string) {
	if flag.Lookup(flagName) == nil {
		flag.StringVar(passedVar, flagName, "", "")
	}
}
