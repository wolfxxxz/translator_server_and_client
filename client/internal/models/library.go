package models

type Library struct {
	English       string       `json:"english"`
	Russian       string       `json:"russian"`
	Theme         string       `json:"theme"`
	PartsOfSpeech string       `json:"part_of_speech"`
	Phrases       []Phrase     `json:"library_phrases"`
	PhraseVerbs   []PhraseVerb `json:"library_phrase_verbs"`
	Exceptions    string       `json:"exceptions"`
}

type Phrase struct {
	ID        int       `json:"id"`
	English   string    `json:"english"`
	Russian   string    `json:"russian"`
	Libraries []Library `gorm:"many2many:library_phrases;" json:"libraries"`
}

type PhraseVerb struct {
	ID        int       `json:"id"`
	English   string    `json:"english"`
	Russian   string    `json:"russian"`
	Libraries []Library `gorm:"many2many:library_phrase_verbs;" json:"libraries"`
}
