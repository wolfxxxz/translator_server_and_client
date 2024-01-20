package clients

import (
	"bytes"
	"client/internal/apperrors"
	"client/internal/config"
	"client/internal/domain/requests"
	"client/internal/domain/responses"
	"client/internal/mappers"
	"client/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	pathLibrary = "/library"
	translate   = "/translate"
)

type LibraryClient interface {
	GetTranslation(word *requests.GetTranslationReq) ([]*models.Library, error)
}
type libraryClient struct {
	config *config.Config
	client *http.Client
	log    *logrus.Logger
	path   string
}

func NewLibraryClient(config *config.Config, client *http.Client, log *logrus.Logger) LibraryClient {
	return &libraryClient{
		config: config,
		client: client,
		log:    log,
		path:   pathLibrary,
	}
}

func (lc libraryClient) GetTranslation(word *requests.GetTranslationReq) ([]*models.Library, error) {
	requestBody, err := json.Marshal(word)
	if err != nil {
		appErr := apperrors.GetTranslationErr.AppendMessage(err)
		lc.log.Error(appErr)
		return nil, appErr
	}

	path := fmt.Sprintf("%v%v%v%v", lc.config.Host, lc.config.AppPort, lc.path, translate)

	req, err := http.NewRequest("GET", path, bytes.NewBuffer(requestBody))
	if err != nil {
		appErr := apperrors.GetTranslationErr.AppendMessage(err)
		lc.log.Error(appErr)
		return nil, appErr
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := lc.client.Do(req)
	if err != nil {
		appErr := apperrors.GetTranslationErr.AppendMessage(err)
		lc.log.Error(appErr)
		return nil, appErr
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			appErr := apperrors.GetTranslationErr.AppendMessage(err)
			lc.log.Error(appErr)
			return nil, appErr
		}

		err = fmt.Errorf("GetTranslation client err [%v] is [%v]", err, string(body))
		appErr := apperrors.GetTranslationErr.AppendMessage(err)
		lc.log.Error(appErr)
		return nil, appErr
	}

	wordsResp := []*responses.GetTranslationResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wordsResp); err != nil {
		appErr := apperrors.GetTranslationErr.AppendMessage(err)
		lc.log.Error(appErr)
		return nil, appErr
	}

	words := mappers.MapGetTranslReqToGetWord(wordsResp)
	return words, nil
}
