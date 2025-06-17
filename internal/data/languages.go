package data

import (
	"context"
	"database/sql"
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
