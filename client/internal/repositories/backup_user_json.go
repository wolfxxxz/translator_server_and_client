package repositories

import (
	"client/internal/apperrors"
	"client/internal/models"
	"encoding/json"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

const backup = "backup/user.json"

type BackupRepo interface {
	GetUserFromBackUp() (*models.User, error)
	SaveUser(user *models.User) error
}
type backupRepo struct {
	path string
	log  *logrus.Logger
}

func NewBackUpCopyRepo(log *logrus.Logger) BackupRepo {
	return &backupRepo{path: backup, log: log}
}

func (br *backupRepo) GetUserFromBackUp() (*models.User, error) {
	filejson, err := os.Open(br.path)
	if err != nil {
		appErr := apperrors.GetUserFromBackUpErr.AppendMessage(err)
		br.log.Error(appErr)
		return nil, nil
	}

	defer filejson.Close()
	data, err := io.ReadAll(filejson)
	if err != nil {
		appErr := apperrors.GetUserFromBackUpErr.AppendMessage(err)
		br.log.Error(appErr)
		return nil, appErr
	}

	user := &models.User{}

	err = json.Unmarshal(data, user)
	if err != nil {
		appErr := apperrors.GetUserFromBackUpErr.AppendMessage(err)
		br.log.Error(appErr)
		return nil, appErr
	}

	return user, nil
}

func (br *backupRepo) SaveUser(user *models.User) error {
	byteArr, err := json.MarshalIndent(user, "", "   ")
	if err != nil {
		appErr := apperrors.SaveUserErr.AppendMessage(err)
		br.log.Error(appErr)
		return err
	}

	err = os.WriteFile(br.path, byteArr, 0666) //-rw-rw-rw- 0664
	if err != nil {
		appErr := apperrors.SaveUserErr.AppendMessage(err)
		br.log.Error(appErr)
		return err
	}

	return nil
}
