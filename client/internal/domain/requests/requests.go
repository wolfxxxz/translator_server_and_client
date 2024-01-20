package requests

type GetTranslationReq struct {
	Word string `json:"word"`
}

type CreateUserRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type GetWordsByUsIdAndLimitRequest struct {
	Limit string `json:"limit"`
	ID    string `json:"user_id"`
}

type DeleteWordFromUserByIDRequest struct {
	UserID string `json:"user_id"`
	WordID string `json:"word_id"`
}

type DeleteLearnFromUserByIDRequest struct {
	UserID string `json:"user_id"`
	WordID string `json:"word_id"`
}

type MoveWordToLearnedRequest struct {
	UserID string `json:"user_id"`
	WordID string `json:"word_id"`
}

type AddWordToLearnedRequest struct {
	UserID string `json:"user_id"`
	WordID string `json:"word_id"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
