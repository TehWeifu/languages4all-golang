package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Quiz struct {
	ID               int64  `json:"id"`
	Language         int64  `json:"language"`
	Skill            int    `json:"skill"`
	Name             string `json:"name"`
	BackgroundColor  string `json:"background_color"`
	IntroPhrase      string `json:"intro_phrase"`
	ImageResource    string `json:"image_resource"`
	ProgressionOrder int    `json:"order"`
}

type QuizModel struct {
	DB *sql.DB
}

func (m *QuizModel) GetAll(language int64) ([]*Quiz, error) {
	query := fmt.Sprintf(`
		SELECT id, language, skill, name, background_color, intro_phrase, image_resource, progression_order
		FROM quizzes
		WHERE language = $1`)

	args := []any{language}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Hour)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quizzes []*Quiz
	for rows.Next() {
		var quiz Quiz

		err := rows.Scan(
			&quiz.ID,
			&quiz.Language,
			&quiz.Skill,
			&quiz.Name,
			&quiz.BackgroundColor,
			&quiz.IntroPhrase,
			&quiz.ImageResource,
			&quiz.ProgressionOrder,
		)

		if err != nil {
			return nil, err
		}

		quizzes = append(quizzes, &quiz)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return quizzes, nil
}
