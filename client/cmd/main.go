package main

import (
	"client/internal/clients"
	"client/internal/competition"
	"client/internal/config"
	"client/internal/log"
	"client/internal/repositories"
	"context"
	"net/http"
)

func main() {
	logger, err := log.NewLogAndSetLevel("info")
	if err != nil {
		logger.Fatal(err)
	}

	conf, err := config.NewConfig(logger)
	if err != nil {
		logger.Fatal(err)
	}

	if err = log.SetLevel(logger, conf.LogLevel); err != nil {
		logger.Fatal(err)
	}

	httpClient := &http.Client{}
	translClient := clients.NewLibraryClient(conf, httpClient, logger)
	backup := repositories.NewBackUpCopyRepo(logger)
	userClient := clients.NewUserClient(conf, httpClient, logger)

	comp := competition.NewCompetition(translClient, userClient, backup, logger)

	ctx := context.Background()
	logger.Info("Start competition")
	err = comp.StartCompetition(ctx)
	if err != nil {
		logger.Fatal(err)
	}
}
