package mappers

import (
	"client/internal/domain/requests"
	"client/internal/domain/responses"
	"client/internal/models"
)

func MapGetTranslReqToGetWord(getTrResp []*responses.GetTranslationResp) []*models.Library {
	words := []*models.Library{}
	for _, word := range getTrResp {
		word := &models.Library{
			English: word.English,
			Russian: word.Russian,
		}

		words = append(words, word)
	}

	return words
}

func MapCreateUserReqToUser(createUsReq *requests.CreateUserRequest) *models.User {
	return &models.User{
		Email:    createUsReq.Email,
		Name:     createUsReq.Name,
		LastName: createUsReq.LastName,
		Password: createUsReq.Password,
		Role:     createUsReq.Role,
	}
}
