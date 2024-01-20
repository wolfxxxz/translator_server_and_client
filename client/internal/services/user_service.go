package services

import (
	"client/internal/apperrors"
	"client/internal/clients"
	"client/internal/domain/requests"
	"client/internal/mappers"
	"client/internal/models"
	"client/internal/repositories"
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type UserService struct {
	clientUser    clients.UserClient
	clientLibrary clients.LibraryClient
	repoBackup    repositories.BackupRepo
	log           *logrus.Logger
}

func NewUserService(clientUser clients.UserClient, clientLibrary clients.LibraryClient,
	repoBackup repositories.BackupRepo, log *logrus.Logger) *UserService {
	return &UserService{clientUser: clientUser, repoBackup: repoBackup, clientLibrary: clientLibrary, log: log}
}

func (us *UserService) UserExistsOrRegistration(ctx context.Context) (*models.User, error) {
	userFromBackup, err := us.repoBackup.GetUserFromBackUp()
	if err != nil {
		us.log.Error(err)
		return nil, err
	}

	user := &models.User{}
	user = userFromBackup

	if user.ID == "" {
		createUsReq := scanUser()
		user = mappers.MapCreateUserReqToUser(createUsReq)
		userId, err := us.clientUser.CreateUser(createUsReq)
		if err != nil {
			us.log.Error(err)
			return nil, err
		}

		user.ID = userId.UserId
	}

	time.Sleep(time.Millisecond * 15)
	for {
		fmt.Printf("[%v] Enter your password", user.Name)
		pass, err := scanLine()
		if err != nil {
			us.log.Error(err)
		}

		loginUsReq := &requests.LoginRequest{Email: user.Email, Password: pass}
		tokenReq, err := us.clientUser.Login(loginUsReq)
		if err != nil {
			us.log.Error(err)
			fmt.Println("wrong password")
			continue
		}

		user.Token = tokenReq.Token
		user.TokenExpired = tokenReq.ExpiresIn
		break
	}

	us.log.Infof("LOGIN_CLIENT success %+v", user.Token)

	userForBackup := *user
	userForBackup.Password = ""
	err = us.repoBackup.SaveUser(&userForBackup)
	if err != nil {
		us.log.Error(err)
	}

	us.log.Info("UserExistsOrRegistration invoked success")
	return user, nil
}

func (c *UserService) TestWords(ctx context.Context, user *models.User, quantity int) error {
	startTime := time.Now()
	limit := strconv.Itoa(quantity)
	getWordsReq := &requests.GetWordsByUsIdAndLimitRequest{ID: user.ID, Limit: limit}
	testTable, err := c.clientUser.GetUserWithWordsByIDLimit(getWordsReq, user.Token)
	if err != nil {
		c.log.Error(err)
		return err
	}

	c.log.Info(len(testTable))

	if len(testTable) == 0 {
		fmt.Println("There aren't what to test")
		return nil
	}

	fmt.Println("                     START")
	var right int
	var wrong int
	fmt.Println("TEST WORDS")

	for {
		word := testTable[0]
		fmt.Println(word.Russian)
		englishAnswer, err := scanLine()
		if err != nil {
			appErr := apperrors.TestWordsErr.AppendMessage(err)
			c.log.Error(appErr)
			return appErr
		}
		englishAnswerIgnoreSpace := ignorSpace(englishAnswer)
		englishWordQuest := ignorSpace(word.English)

		if strings.EqualFold(englishWordQuest, englishAnswerIgnoreSpace) {
			right++
			fmt.Println("Yes")
			moveToLearnedReq := &requests.MoveWordToLearnedRequest{WordID: word.ID, UserID: user.ID}
			err := c.clientUser.MoveWordToLearned(moveToLearnedReq, user.Token)
			if err != nil {
				c.log.Error(err)
				return err
			}

			if len(testTable) == 1 {
				break
			}

			testTable = testTable[1:]
			continue
		}

		if compareStringsLevenshtein(englishWordQuest, englishAnswerIgnoreSpace) {
			right++
			fmt.Println("Yes")
			fmt.Println("Spelling mistake ", word.English)
			moveToLearnedReq := &requests.MoveWordToLearnedRequest{WordID: word.ID, UserID: user.ID}
			err := c.clientUser.MoveWordToLearned(moveToLearnedReq, user.Token)
			if err != nil {
				c.log.Error(err)
				return err
			}

			if len(testTable) == 1 {
				break
			}

			testTable = testTable[1:]
			continue
		}

		wrong++
		getTranslReq := &requests.GetTranslationReq{Word: word.English}
		lib, err := c.clientLibrary.GetTranslation(getTranslReq)
		if err != nil {
			c.log.Error(err)
			return err
		}

		printAll(lib)
		for {
			englishAnswer, err := scanLine()
			if err != nil {
				appErr := apperrors.TestWordsErr.AppendMessage(err)
				c.log.Error(appErr)
				return appErr
			}

			englishAnswerIgnoreSpace := ignorSpace(englishAnswer)
			if strings.EqualFold(englishWordQuest, englishAnswerIgnoreSpace) {
				break
			}
		}

		moveToLearnedReq := &requests.MoveWordToLearnedRequest{WordID: word.ID, UserID: user.ID}
		err = c.clientUser.AddWordToLearn(moveToLearnedReq, user.Token)
		if err != nil {
			return err
		}

		if len(testTable) == 1 {
			break
		}

		testTable = testTable[1:]
	}

	duration := time.Since(startTime)
	printTime(duration)
	fmt.Println(right, wrong)

	return nil
}

func (us *UserService) LearnWords(ctx context.Context, quantity int, user *models.User) error {
	startTime := time.Now()
	limit := strconv.Itoa(quantity)
	getWordsReq := &requests.GetWordsByUsIdAndLimitRequest{ID: user.ID, Limit: limit}
	testTable, err := us.clientUser.GetUserWithLearnByIDLimit(getWordsReq, user.Token)
	if err != nil {
		us.log.Error(err)
		return err
	}

	if len(testTable) == 0 {
		fmt.Println("There isn't what to learn")
		return nil
	}

	fmt.Println("                 START")
	fmt.Println("LEARN WORDS")

	for {
		word := testTable[0]
		fmt.Println(word.Russian)
		englishAnswer, err := scanLine()
		if err != nil {
			appErr := apperrors.LearnWordsErr.AppendMessage(err)
			us.log.Error(appErr)
			return appErr
		}

		englishAnswerIgnoreSpace := ignorSpace(englishAnswer)
		englishWordQust := ignorSpace(word.English)

		if strings.EqualFold(englishWordQust, englishAnswerIgnoreSpace) {
			fmt.Println("Yes")
			deleteLearnReq := &requests.DeleteLearnFromUserByIDRequest{UserID: user.ID, WordID: word.ID}
			err := us.clientUser.DeleteLearnWordFromUserByWord(deleteLearnReq, user.Token)
			if err != nil {
				us.log.Error(err)
				return err
			}

			if len(testTable) == 1 {
				break
			}

			testTable = testTable[1:]
			continue
		}

		if compareStringsLevenshtein(englishWordQust, englishAnswerIgnoreSpace) {
			fmt.Println("Yes")
			fmt.Println("Spelling mistake ", word.English)
			deleteLearnReq := &requests.DeleteLearnFromUserByIDRequest{UserID: user.ID, WordID: word.ID}
			err := us.clientUser.DeleteLearnWordFromUserByWord(deleteLearnReq, user.Token)
			if err != nil {
				us.log.Error(err)
				return err
			}

			if len(testTable) == 1 {
				break
			}

			testTable = testTable[1:]
			continue
		}

		getTranslReq := &requests.GetTranslationReq{Word: word.English}
		lib, err := us.clientLibrary.GetTranslation(getTranslReq)
		if err == nil {
			us.log.Error(err)
			return err
		}

		printAll(lib)
		if len(testTable) == 1 {
			break
		}

		testTable = append(testTable, word)
		testTable = testTable[1:]

	}

	duration := time.Since(startTime)
	printTime(duration)
	return nil
}
