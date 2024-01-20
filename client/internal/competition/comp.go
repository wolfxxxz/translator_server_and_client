package competition

import (
	"client/internal/clients"
	"client/internal/models"
	"client/internal/repositories"
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type Competition struct {
	clientLibrary clients.LibraryClient
	clientUser    clients.UserClient
	repoBackup    repositories.BackupRepo
	log           *logrus.Logger
}

func NewCompetition(clientLibrary clients.LibraryClient, clientUser clients.UserClient,
	repoBackup repositories.BackupRepo, log *logrus.Logger) *Competition {
	return &Competition{
		clientLibrary: clientLibrary,
		clientUser:    clientUser,
		repoBackup:    repoBackup,
		log:           log,
	}
}

func (c *Competition) StartCompetition(ctx context.Context) error {
	user, err := c.userExistOrRegistration(ctx)

	if err != nil {
		c.log.Error(err)
	}

	for {
		exit, err := c.scanCommand(ctx, user)
		if err != nil {
			c.log.Error(err)
			return err
		}

		if exit {
			break
		}
	}

	return nil

}

func (c *Competition) scanCommand(ctx context.Context, user *models.User) (bool, error) {
	printInfoMenu()
	var command string
	fmt.Scan(&command)
	switch command {
	case test:
		if err := c.test(ctx, user); err != nil {
			return false, err
		}

	case learn:
		if err := c.learn(ctx, user); err != nil {
			return false, err
		}

	case translate:
		if err := c.translator(ctx); err != nil {
			c.log.Error(err)
			return false, err
		}

	case exit:
		fmt.Println("    Good buy, have a good day !!!")
		return true, nil
	}

	return false, nil
}

func printInfoMenu() {
	menu := []string{
		fmt.Sprintf("      Test knowledge:   [%v]\n", test),
		fmt.Sprintf("      Learn words:     [%v]\n", learn),
		fmt.Sprintf("      Translator:  [%v]\n", translate),
		fmt.Sprintf("          Exit:        [%v]\n", exit),
	}

	for _, pos := range menu {
		fmt.Println(pos)
		time.Sleep(20 * time.Millisecond)
	}
}
