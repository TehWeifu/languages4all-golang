package data

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

type User struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	AndroidUid string `json:"android_uid"`
}

type UserRanking struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
	Total  int    `json:"total"`
}

type UserModel struct {
	DB *sql.DB
}

func (m UserModel) GetByNameAndAndroidUid(name, androidUid string) (*User, error) {
	query := `
		SELECT id, name, android_uid 
		FROM users
		WHERE name = $1 AND android_uid = $2`

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Hour)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, name, androidUid).Scan(
		&user.ID,
		&user.Name,
		&user.AndroidUid,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (m UserModel) Insert(user *User) error {
	query := `
		INSERT INTO users (android_uid, name) 
		VALUES ($1, $2)
		RETURNING id`

	args := []any{user.AndroidUid, user.Name}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Hour)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&user.ID)
	// TODO: do error triaging to discover if there is a duplicated name + android_uid
	return err
}

func (m UserModel) GetRanking(language int64, androidUid string) ([]*UserRanking, error) {
	query := `
		SELECT user_id, u.name, SUM(maxpoints) AS total 
		FROM user_quizzes_points
		JOIN users u on u.id = user_quizzes_points.user_id
		JOIN quizzes q on user_quizzes_points.quiz_id = q.id
		WHERE q.language = $1 AND (android_uid = $2 OR $2 = '')
		GROUP BY user_id, u.name 
		ORDER BY total
		LIMIT 50`

	args := []any{language, androidUid}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Hour)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usersRanking := []*UserRanking{}
	for rows.Next() {
		var userRanking UserRanking

		err := rows.Scan(
			&userRanking.UserID,
			&userRanking.Name,
			&userRanking.Total,
		)

		if err != nil {
			return nil, err
		}

		usersRanking = append(usersRanking, &userRanking)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return usersRanking, nil
}
