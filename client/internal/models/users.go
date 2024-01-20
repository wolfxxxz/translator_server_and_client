package models

import (
	"strings"
)

type User struct {
	ID           string  `json:"id"`
	Email        string  `json:"user_email"`
	Name         string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Password     string  `json:"password"`
	Role         string  `json:"role"`
	Words        []*Word `json:"user_words"`
	Learn        []*Word `json:"user_learn"`
	Learned      []*Word `json:"user_learned"`
	Token        string
	TokenExpired string
}

type Word struct {
	ID      string `json:"id"`
	English string `json:"english"`
	Russian string `json:"russian"`
	Theme   string `json:"theme"`
}

func CreateAndInitMapWords(s []*Word) *map[string][]string {
	maps := make(map[string][]string)
	for _, w := range s {
		maps[w.Russian] = append(maps[w.Russian], strings.ToLower(w.English))
	}

	return &maps
}
