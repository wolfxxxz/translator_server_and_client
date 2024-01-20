package clients

import (
	"bytes"
	"client/internal/apperrors"
	"client/internal/config"
	"client/internal/domain/requests"
	"client/internal/domain/responses"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

const (
	users          = "/users"
	login          = "/login"
	user           = "/user"
	words          = "/words"
	moveToLearned  = "/move-word-to-learned"
	learn          = "/learn"
	addWordToLearn = "/add-word-to-learn"
)

type UserClient interface {
	CreateUser(createUsReq *requests.CreateUserRequest) (*responses.CreateUserResponse, error)
	Login(loginReq *requests.LoginRequest) (*responses.LoginResponse, error)
	GetUserWithWordsByIDLimit(getWordsReq *requests.GetWordsByUsIdAndLimitRequest, token string) ([]*responses.WordResp, error)
	MoveWordToLearned(getWordsReq *requests.MoveWordToLearnedRequest, token string) error
	AddWordToLearn(getWordsReq *requests.MoveWordToLearnedRequest, token string) error
	GetUserWithLearnByIDLimit(getWordsReq *requests.GetWordsByUsIdAndLimitRequest, token string) ([]*responses.WordResp, error)
	DeleteLearnWordFromUserByWord(deleteWordFromLearn *requests.DeleteLearnFromUserByIDRequest, token string) error
}

type userClient struct {
	config   *config.Config
	client   *http.Client
	log      *logrus.Logger
	pathHttp string
}

func NewUserClient(config *config.Config, client *http.Client, log *logrus.Logger) UserClient {
	return &userClient{
		config:   config,
		client:   client,
		log:      log,
		pathHttp: users,
	}
}

func (uc *userClient) CreateUser(createUsReq *requests.CreateUserRequest) (*responses.CreateUserResponse, error) {
	requestBody, err := json.Marshal(createUsReq)
	if err != nil {
		appErr := apperrors.CreateUserErr.AppendMessage(err)
		uc.log.Error(appErr)
		return nil, appErr
	}

	path := fmt.Sprintf("%v%v%v", uc.config.Host, uc.config.AppPort, uc.pathHttp)

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(requestBody))
	if err != nil {
		appErr := apperrors.CreateUserErr.AppendMessage(err)
		uc.log.Error(appErr)
		return nil, appErr
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := uc.client.Do(req)
	if err != nil {
		appErr := apperrors.CreateUserErr.AppendMessage(err)
		uc.log.Error(appErr)
		return nil, appErr
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			appErr := apperrors.CreateUserErr.AppendMessage(err)
			uc.log.Error(appErr)
			return nil, appErr
		}

		msg := fmt.Sprintf("error [%v] is [%v]", err, string(body))
		appErr := apperrors.CreateUserErr.AppendMessage(msg)
		uc.log.Error(appErr)
		return nil, appErr
	}

	userResp := &responses.CreateUserResponse{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(userResp); err != nil {
		appErr := apperrors.CreateUserErr.AppendMessage(err)
		uc.log.Error(appErr)
		return nil, appErr
	}

	return userResp, nil
}

func (uc *userClient) Login(loginReq *requests.LoginRequest) (*responses.LoginResponse, error) {
	requestBody, err := json.Marshal(loginReq)
	if err != nil {
		appErr := apperrors.LoginErr.AppendMessage(err)
		uc.log.Error(appErr)
		return nil, appErr
	}

	path := fmt.Sprintf("%v%v%v%v", uc.config.Host, uc.config.AppPort, uc.pathHttp, login)

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(requestBody))
	if err != nil {
		appErr := apperrors.LoginErr.AppendMessage(err)
		uc.log.Error(appErr)
		return nil, appErr
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := uc.client.Do(req)
	if err != nil {
		appErr := apperrors.LoginErr.AppendMessage(err)
		uc.log.Error(appErr)
		return nil, appErr
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			appErr := apperrors.LoginErr.AppendMessage(err)
			uc.log.Error(appErr)
			return nil, appErr
		}

		msg := fmt.Sprintf("err [%v] is [%v]", err, string(body))
		appErr := apperrors.LoginErr.AppendMessage(msg)
		uc.log.Error(appErr)
		return nil, appErr
	}

	userResp := &responses.LoginResponse{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(userResp); err != nil {
		appErr := apperrors.LoginErr.AppendMessage(err)
		uc.log.Error(appErr)
		return nil, appErr
	}

	return userResp, nil
}

func (uc *userClient) GetUserWithWordsByIDLimit(getWordsReq *requests.GetWordsByUsIdAndLimitRequest, token string) ([]*responses.WordResp, error) {
	requestBody, err := json.Marshal(getWordsReq)
	if err != nil {
		appErr := apperrors.GetUserWithWordsByIDLimitErr.AppendMessage(err)
		uc.log.Error(appErr)
		return nil, appErr
	}

	path := fmt.Sprintf("%v%v%v%v", uc.config.Host, uc.config.AppPort, user, words)

	req, err := http.NewRequest("GET", path, bytes.NewBuffer(requestBody))
	if err != nil {
		appErr := apperrors.GetUserWithWordsByIDLimitErr.AppendMessage(err)
		uc.log.Error(appErr)
		return nil, appErr
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := uc.client.Do(req)
	if err != nil {
		appErr := apperrors.GetUserWithWordsByIDLimitErr.AppendMessage(err)
		uc.log.Error(appErr)
		return nil, appErr
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			appErr := apperrors.GetUserWithWordsByIDLimitErr.AppendMessage(err)
			uc.log.Error(appErr)
			return nil, appErr
		}

		msg := fmt.Sprintf("err [%v] is [%v]", err, string(body))
		appErr := apperrors.GetUserWithWordsByIDLimitErr.AppendMessage(msg)
		uc.log.Error(appErr)
		return nil, appErr
	}

	wordsResp := []*responses.WordResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wordsResp); err != nil {
		appErr := apperrors.GetUserWithWordsByIDLimitErr.AppendMessage(err)
		uc.log.Error(appErr)
		return nil, appErr
	}

	return wordsResp, nil
}

func (uc *userClient) MoveWordToLearned(getWordsReq *requests.MoveWordToLearnedRequest, token string) error {
	requestBody, err := json.Marshal(getWordsReq)
	if err != nil {
		appErr := apperrors.MoveWordToLearnedErr.AppendMessage(err)
		uc.log.Error(appErr)
		return appErr
	}

	path := fmt.Sprintf("%v%v%v%v", uc.config.Host, uc.config.AppPort, user, moveToLearned)

	req, err := http.NewRequest("PUT", path, bytes.NewBuffer(requestBody))
	if err != nil {
		appErr := apperrors.MoveWordToLearnedErr.AppendMessage(err)
		uc.log.Error(appErr)
		return appErr
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := uc.client.Do(req)
	if err != nil {
		appErr := apperrors.MoveWordToLearnedErr.AppendMessage(err)
		uc.log.Error(appErr)
		return appErr
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			appErr := apperrors.MoveWordToLearnedErr.AppendMessage(err)
			uc.log.Error(appErr)
			return appErr
		}

		msg := fmt.Errorf("err [%v] is [%v]", err, string(body))
		appErr := apperrors.MoveWordToLearnedErr.AppendMessage(msg)
		uc.log.Error(appErr)
		return appErr
	}

	result := &responses.Result{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(result); err != nil {
		appErr := apperrors.MoveWordToLearnedErr.AppendMessage(err)
		uc.log.Error(appErr)
		return appErr
	}

	uc.log.Infof("%+v", result)

	return nil
}

func (uc *userClient) AddWordToLearn(getWordsReq *requests.MoveWordToLearnedRequest, token string) error {
	requestBody, err := json.Marshal(getWordsReq)
	if err != nil {
		appErr := apperrors.AddWordToLearnErr.AppendMessage(err)
		uc.log.Error(appErr)
		return appErr
	}

	path := fmt.Sprintf("%v%v%v%v", uc.config.Host, uc.config.AppPort, user, addWordToLearn)

	req, err := http.NewRequest("POST", path, bytes.NewBuffer(requestBody))
	if err != nil {
		appErr := apperrors.AddWordToLearnErr.AppendMessage(err)
		uc.log.Error(appErr)
		return appErr
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := uc.client.Do(req)
	if err != nil {
		appErr := apperrors.AddWordToLearnErr.AppendMessage(err)
		uc.log.Error(appErr)
		return appErr
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			appErr := apperrors.AddWordToLearnErr.AppendMessage(err)
			uc.log.Error(appErr)
			return appErr
		}

		msg := fmt.Errorf("err [%v] is [%v]", err, string(body))
		appErr := apperrors.AddWordToLearnErr.AppendMessage(msg)
		uc.log.Error(appErr)
		return appErr
	}

	result := &responses.Result{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(result); err != nil {
		appErr := apperrors.AddWordToLearnErr.AppendMessage(err)
		uc.log.Error(appErr)
		return appErr
	}

	uc.log.Infof("%+v", result)

	return nil
}

func (uc *userClient) GetUserWithLearnByIDLimit(getWordsReq *requests.GetWordsByUsIdAndLimitRequest, token string) ([]*responses.WordResp, error) {
	requestBody, err := json.Marshal(getWordsReq)
	if err != nil {
		appErr := apperrors.GetUserWithLearnByIDLimitErr.AppendMessage(err)
		uc.log.Error(appErr)
		return nil, appErr
	}

	path := fmt.Sprintf("%v%v%v%v", uc.config.Host, uc.config.AppPort, user, learn)

	req, err := http.NewRequest("GET", path, bytes.NewBuffer(requestBody))
	if err != nil {
		appErr := apperrors.GetUserWithLearnByIDLimitErr.AppendMessage(err)
		uc.log.Error(appErr)
		return nil, appErr
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := uc.client.Do(req)
	if err != nil {
		appErr := apperrors.GetUserWithLearnByIDLimitErr.AppendMessage(err)
		uc.log.Error(appErr)
		return nil, appErr
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			appErr := apperrors.GetUserWithLearnByIDLimitErr.AppendMessage(err)
			uc.log.Error(appErr)
			return nil, appErr
		}

		msg := fmt.Sprintf("err [%v] is [%v]", err, string(body))
		appErr := apperrors.GetUserWithLearnByIDLimitErr.AppendMessage(msg)
		uc.log.Error(appErr)
		return nil, appErr
	}

	wordsResp := []*responses.WordResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wordsResp); err != nil {
		appErr := apperrors.GetUserWithLearnByIDLimitErr.AppendMessage(err)
		uc.log.Error(appErr)
		return nil, appErr
	}

	return wordsResp, nil
}

func (uc *userClient) DeleteLearnWordFromUserByWord(deleteWordFromLearn *requests.DeleteLearnFromUserByIDRequest, token string) error {
	requestBody, err := json.Marshal(deleteWordFromLearn)
	if err != nil {
		appErr := apperrors.DeleteLearnWordFromUserByWordErr.AppendMessage(err)
		uc.log.Error(appErr)
		return appErr
	}

	path := fmt.Sprintf("%v%v%v%v", uc.config.Host, uc.config.AppPort, user, learn)

	req, err := http.NewRequest("DELETE", path, bytes.NewBuffer(requestBody))
	if err != nil {
		appErr := apperrors.DeleteLearnWordFromUserByWordErr.AppendMessage(err)
		uc.log.Error(appErr)
		return appErr
	}

	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := uc.client.Do(req)
	if err != nil {
		appErr := apperrors.DeleteLearnWordFromUserByWordErr.AppendMessage(err)
		uc.log.Error(appErr)
		return appErr
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			appErr := apperrors.DeleteLearnWordFromUserByWordErr.AppendMessage(err)
			uc.log.Error(appErr)
			return appErr
		}

		msg := fmt.Sprintf("err [%v] is [%v]", err, string(body))
		appErr := apperrors.DeleteLearnWordFromUserByWordErr.AppendMessage(msg)
		uc.log.Error(appErr)
		return appErr
	}

	wordsResp := &responses.Result{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wordsResp); err != nil {
		appErr := apperrors.DeleteLearnWordFromUserByWordErr.AppendMessage(err)
		uc.log.Error(appErr)
		return appErr
	}

	return nil
}
