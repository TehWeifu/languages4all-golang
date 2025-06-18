package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Users     UserModel
	Languages LanguageModel
	Quizzes   QuizModel
	Points    PointModel
	Questions QuestionModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Users:     UserModel{DB: db},
		Languages: LanguageModel{DB: db},
		Quizzes:   QuizModel{DB: db},
		Points:    PointModel{DB: db},
		Questions: QuestionModel{DB: db},
	}
}
