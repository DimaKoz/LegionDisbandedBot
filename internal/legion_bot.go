package internal

import "go.uber.org/zap"

func StartLegionBot(logger *zap.Logger) {
	initLogger(logger)
	defer loggerSync()
	zap.S().Infoln("StartMonitor(), logger is ready")

	zap.S().Infoln("exiting")
}

func loggerSync() {
	_ = zap.L().Sync()
}

func initLogger(logger *zap.Logger) {
	zap.ReplaceGlobals(logger)
}
