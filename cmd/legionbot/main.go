package main

import (
	"log"

	"github.com/DimaKoz/LegionDisbandedBot/internal"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't start by %v", err)
	}
	internal.StartLegionBot(logger)
}
