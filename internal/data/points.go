package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Point struct {
	ID                   int64 `json:"-"`
	QuizID               int64 `json:"quiz_id"`
	UserID               int64 `json:"-"`
	Points               int   `json:"points"`
	MaxPoints            int   `json:"maxPoints"`
	Completed            int   `json:"completed"`
	CurrentQuestionOrder int   `json:"currentQuestionOrder"`
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

func (m *PointModel) Upsert(point *Point) error {
	query := `
		SELECT id, maxpoints, completed
		FROM user_quizzes_points
		WHERE quiz_id = $1 AND user_id = $2`

	args := []any{point.QuizID, point.UserID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Hour)
	defer cancel()

	var ID int64
	var maxPoints int
	var completed int

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&ID, &maxPoints, &completed)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			// Insert the new score record
			point.MaxPoints = point.Points
			if point.Completed == 1 {
				point.Points = 0
			}

			query = `
				INSERT INTO user_quizzes_points (user_id, quiz_id, points, completed, currentquestionorder, maxpoints) 
				VALUES ($1, $2, $3, $4, $5, $6)`

			args := []any{point.UserID, point.QuizID, point.Points, point.Completed, point.CurrentQuestionOrder, point.MaxPoints}

			ctx, cancel = context.WithTimeout(context.Background(), 3*time.Hour)
			defer cancel()

			_, err := m.DB.ExecContext(ctx, query, args...)
			return err

		default:
			return err
		}
	}

	// Update the existing score record
	point.MaxPoints = max(point.Points, maxPoints)

	if point.Completed == 1 {
		point.Points = 0
	}

	if completed == 1 {
		point.Completed = 1
	}
	point.ID = ID

	query = `
		UPDATE user_quizzes_points
		SET points = $1, maxpoints = $2, completed = $3, currentquestionorder = $4
		WHERE id = $5`

	args = []any{point.Points, point.MaxPoints, point.Completed, point.CurrentQuestionOrder, point.ID}

	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Hour)
	defer cancel()

	_, err = m.DB.ExecContext(ctx, query, args...)
	return err
}
