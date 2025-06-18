package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Question struct {
	ID               int64  `json:"id"`
	Preposition      string `json:"preposition"`
	ImageResource    string `json:"image_resource"`
	QuizID           int64  `json:"quiz_id"`
	ProgressionOrder int    `json:"order"`
	Answer           string `json:"answer"`
}

type QuestionModel struct {
	DB *sql.DB
}

func (m *QuestionModel) GetCountByLanguage(language int64) (int, error) {
	query := fmt.Sprintf(`
	SELECT count(*) AS count_questions
	FROM questions
	JOIN public.quizzes q on q.id = questions.quiz_id
	WHERE q.language = $1`)

	args := []any{language}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Hour)
	defer cancel()

	var questionCount int

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&questionCount)
	if err != nil {
		return 0, err
	}

	return questionCount, nil
}

func (m *QuestionModel) GetAll(quizID int64) ([]*Question, error) {
	query := fmt.Sprintf(`
	SELECT questions.id, preposition, image_resource, quiz_id, progression_order, a.solution
	FROM questions
	JOIN public.answers a on questions.id = a.question_id
	WHERE quiz_id = $1 OR $1 = 0`)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Hour)
	defer cancel()

	args := []any{quizID}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	questions := []*Question{}

	for rows.Next() {
		var question Question

		err := rows.Scan(
			&question.ID,
			&question.Preposition,
			&question.ImageResource,
			&question.QuizID,
			&question.ProgressionOrder,
			&question.Answer,
		)

		if err != nil {
			return nil, err
		}

		questions = append(questions, &question)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}

func (m *QuestionModel) GetRandomByLanguage(language int64) ([]*Question, error) {
	query := fmt.Sprintf(`
	SELECT q.id, q.preposition, q.image_resource, q.quiz_id, q.progression_order, a.solution
	FROM questions q
		JOIN public.answers a on q.id = a.question_id
		JOIN public.quizzes qzzs on qzzs.id = q.quiz_id
	WHERE qzzs.language = $1
	ORDER BY RANDOM()
	LIMIT 10
	`)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Hour)
	defer cancel()

	args := []any{language}

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	questions := []*Question{}

	for rows.Next() {
		var question Question

		err := rows.Scan(
			&question.ID,
			&question.Preposition,
			&question.ImageResource,
			&question.QuizID,
			&question.ProgressionOrder,
			&question.Answer,
		)

		if err != nil {
			return nil, err
		}

		questions = append(questions, &question)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return questions, nil
}
