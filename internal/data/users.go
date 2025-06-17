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
