package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Point struct {
	QuizID               int64  `json:"quiz_id"`
	Points               int    `json:"points"`
	MaxPoints            int    `json:"maxPoints"`
	Completed            int    `json:"completed"`
	CurrentQuestionOrder string `json:"currentQuestionOrder"`
}

type PointModel struct {
	DB *sql.DB
}

func (m *PointModel) GetAll(userId, language int64) ([]*Point, error) {
	query := fmt.Sprintf(`
		SELECT quiz_id, points, maxpoints, completed, currentquestionorder
		FROM user_quizzes_points
		JOIN public.quizzes q on q.id = user_quizzes_points.quiz_id
		WHERE user_id = $1 AND  q.language = $2`)

	args := []any{userId, language}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Hour)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	points := []*Point{}
	for rows.Next() {
		var point Point

		err := rows.Scan(
			&point.QuizID,
			&point.Points,
			&point.MaxPoints,
			&point.Completed,
			&point.CurrentQuestionOrder,
		)

		if err != nil {
			return nil, err
		}

		points = append(points, &point)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return points, nil
}
