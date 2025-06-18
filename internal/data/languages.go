package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Language struct {
	ID            int64  `json:"id"`
	Code          string `json:"code"`
	Name          string `json:"name"`
	ImageResource string `json:"image_resource"`
	ChooseTitle   string `json:"choose_title"`
}

type LanguageModel struct {
	DB *sql.DB
}

func (m *LanguageModel) GetAll() ([]*Language, error) {
	query := fmt.Sprintf(`
		SELECT id, code, name, image_resource, choose_title
		FROM languages`)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Hour)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var languages []*Language
	for rows.Next() {
		var language Language

		err := rows.Scan(
			&language.ID,
			&language.Code,
			&language.Name,
			&language.ImageResource,
			&language.ChooseTitle,
		)

		if err != nil {
			return nil, err
		}

		languages = append(languages, &language)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return languages, nil
}

func (m *LanguageModel) GetById(id int64) (*Language, error) {
	query := fmt.Sprintf(`
		SELECT id, code, name, image_resource, choose_title
		FROM languages
		WHERE id = $1`)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Hour)
	defer cancel()

	var language Language

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&language.ID,
		&language.Code,
		&language.Name,
		&language.ImageResource,
		&language.ChooseTitle,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &language, nil
}
