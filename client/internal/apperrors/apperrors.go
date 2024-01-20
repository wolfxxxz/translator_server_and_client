package apperrors

import "fmt"

type AppError struct {
	Message string
	Code    string
}

func NewAppError() *AppError {
	return &AppError{}
}

var (
	SetupDatabaseErr = AppError{
		Message: "Failed SetupDatabaseErr",
		Code:    envInit,
	}
	EnvConfigLoadError = AppError{
		Message: "Failed to load env file",
		Code:    envInit,
	}
	EnvConfigParseError = AppError{
		Message: "Failed to parse env file",
		Code:    envParse,
	}
	InitPostgressErr = AppError{
		Message: "Failed to InitPostgress",
		Code:    envParse,
	}
	NewLoggerErr = AppError{
		Message: "Failed to NewLog",
		Code:    log,
	}
	//Repo
	GetUserFromBackUpErr = AppError{
		Message: "Failed to GetUserFromBackUpErr",
		Code:    backUpRepo,
	}
	SaveUserErr = AppError{
		Message: "Failed to SaveUserErr",
		Code:    backUpRepo,
	}
	GetAllWordsLibErr = AppError{
		Message: "Failed to GetAllWordsLibErr",
		Code:    repoLibrary,
	}
	UpdateWordErr = AppError{
		Message: "Failed to UpdateWord",
		Code:    repoLibrary,
	}
	UpdateUserErr = AppError{
		Message: "Failed to UpdateUserErr",
		Code:    repoUsers,
	}
	AddWordToLearnErr = AppError{
		Message: "Failed to AddWordToLearnErr",
		Code:    repoUsers,
	}
	AddWordToLearnedErr = AppError{
		Message: "Failed to AddWordToLearnedErr",
		Code:    repoUsers,
	}
	DeleteWordFromUserByWordErr = AppError{
		Message: "Failed to DeleteWordFromUserByWordErr",
		Code:    repoUsers,
	}
	CreateUserErr = AppError{
		Message: "Failed to CreateUserErr",
		Code:    repoUsers,
	}
	LoginErr = AppError{
		Message: "Failed to LoginErr",
		Code:    repoUsers,
	}
	GetUserWithLearnByNameErr = AppError{
		Message: "Failed to GetUserWithLearnByNameErr",
		Code:    repoUsers,
	}
	GetUserByNameErr = AppError{
		Message: "Failed to GetUserByNameErr",
		Code:    repoUsers,
	}
	DeleteLearnWordFromUserByWordErr = AppError{
		Message: "Failed to DeleteLearnWordFromUserByWordErr",
		Code:    clientUser,
	}
	GetUserWithWordsByIDLimitErr = AppError{
		Message: "Failed to GetUserWithWordsByNameLimitErr",
		Code:    clientUser,
	}
	GetUserWithLearnByIDLimitErr = AppError{
		Message: "Failed to GetUserWithLearnByIDLimitErr",
		Code:    clientUser,
	}
	MoveWordToLearnedErr = AppError{
		Message: "Failed to MoveWordToLearnedErr",
		Code:    clientUser,
	}
	GetTranslationErr = AppError{
		Message: "Failed to GetTranslationErr",
		Code:    clientLibrary,
	}
	StartCompetitionErr = AppError{
		Message: "Failed to StartCompetitionErr",
		Code:    competition,
	}
	LearnWordsErr = AppError{
		Message: "Failed to LearnWordsErr",
		Code:    serviceUser,
	}
	TestWordsErr = AppError{
		Message: "Failed to TestWordsErr",
		Code:    serviceUser,
	}
	TranslateErr = AppError{
		Message: "Failed to TranslateErr",
		Code:    serviceLibrary,
	}
)

func (appError *AppError) Error() string {
	return appError.Code + ": " + appError.Message
}

func (appError *AppError) AppendMessage(anyErrs ...interface{}) *AppError {
	return &AppError{
		Message: fmt.Sprintf("%v : %v", appError.Message, anyErrs),
		Code:    appError.Code,
	}
}

func IsAppError(err1 error, err2 *AppError) bool {
	err, ok := err1.(*AppError)
	if !ok {
		return false
	}

	return err.Code == err2.Code
}
