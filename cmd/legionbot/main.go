package main

import (
	"fmt"
	"log"

	"github.com/DimaKoz/LegionDisbandedBot/internal"
	"go.uber.org/zap"
)

func main() {
	s := "gopher"

	// permit
	fmt.Printf("Hello and welcome, %s!\n", s) //nolint:forbidigo

	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't start by %v", err)
	}
	internal.StartLegionBot(logger)
}
