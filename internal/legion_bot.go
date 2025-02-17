package internal

import (
	"github.com/DimaKoz/LegionDisbandedBot/internal/configer"
	"github.com/DimaKoz/LegionDisbandedBot/internal/repository"
	"go.uber.org/zap"
)

func StartLegionBot(logger *zap.Logger) {
	initLogger(logger)
	defer loggerSync()
	zap.S().Infoln("StartLegionBot(), logger is ready")

	config, err := configer.LoadLegionBotConfig()
	if err != nil {
		zap.S().Warnln("LoadLegionBotConfig() failed by error:\n", err)

		return
	}
	zap.S().Infoln("config:", config)

	if err = repository.LoadWhiteListUser(config.PathWhiteListAA); err != nil {
		zap.S().Warnln("LoadWhiteListUser() failed by error:\n", err)

		return
	}

	if err = repository.LoadCachedTelegramUser(config.PathTelegramUsers); err != nil {
		zap.S().Warnln("LoadCachedTelegramUser() failed by error:\n", err)

		return
	}

	zap.S().Infoln("exiting")
}

func loggerSync() {
	_ = zap.L().Sync()
}

func initLogger(logger *zap.Logger) {
	zap.ReplaceGlobals(logger)
}
