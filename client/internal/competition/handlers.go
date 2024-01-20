package competition

import (
	"client/internal/models"
	"client/internal/services"
	"context"
	"fmt"
)

func (c *Competition) translator(ctx context.Context) error {
	fmt.Println(enterAWorldOfAPart)
	libServ := services.NewLibraryService(c.clientLibrary, c.log)
	return libServ.Translate(ctx)
}

func (c *Competition) userExistOrRegistration(ctx context.Context) (*models.User, error) {
	userService := services.NewUserService(c.clientUser, c.clientLibrary, c.repoBackup, c.log)
	return userService.UserExistsOrRegistration(ctx)
}

func (c *Competition) test(ctx context.Context, user *models.User) error {
	var quantity int
	fmt.Println(numberOfWordsForTheTest)
	fmt.Scan(&quantity)
	userService := services.NewUserService(c.clientUser, c.clientLibrary, c.repoBackup, c.log)
	return userService.TestWords(ctx, user, quantity)
}

func (c *Competition) learn(ctx context.Context, user *models.User) error {
	var quantity int
	fmt.Println(numberOfWordsForTheTest)
	fmt.Scan(&quantity)
	userService := services.NewUserService(c.clientUser, c.clientLibrary, c.repoBackup, c.log)
	return userService.LearnWords(ctx, quantity, user)
}
