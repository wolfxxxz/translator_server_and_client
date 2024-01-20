package services

import (
	"client/internal/apperrors"
	"client/internal/clients"
	"client/internal/domain/requests"
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
)

type LibraryService struct {
	clientLibrary clients.LibraryClient
	log           *logrus.Logger
}

func NewLibraryService(clientLibrary clients.LibraryClient, log *logrus.Logger) *LibraryService {
	return &LibraryService{clientLibrary: clientLibrary, log: log}
}

func (sl *LibraryService) Translate(ctx context.Context) error {
	for {
		fmt.Println()
		word, err := scanLine()
		if err != nil {
			appErr := apperrors.TranslateErr.AppendMessage(err)
			sl.log.Error(appErr)
			return appErr
		}

		if word == "exit" {
			break
		}

		wordRequest := &requests.GetTranslationReq{Word: word}
		words, err := sl.clientLibrary.GetTranslation(wordRequest)
		if err != nil {
			appErr := apperrors.TranslateErr.AppendMessage(err)
			sl.log.Error(appErr)
			return appErr
		}

		printAll(words)
		continue

	}

	return nil
}
