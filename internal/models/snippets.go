package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
)

type Snippet struct {
	ID      int       `db:"id"`
	Title   string    `db:"title"`
	Content string    `db:"content"`
	Created time.Time `db:"created"`
	Expires time.Time `db:"expires"`
}

type SnippetModel struct {
	DB *sqlx.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires) 
          VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL '1 day' * $3)
          RETURNING id`

	var id int
	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > CURRENT_TIMESTAMP AND id = $1`

	var s Snippet
	err := m.DB.Get(&s, stmt, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}

	return &s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > CURRENT_TIMESTAMP ORDER BY created DESC LIMIT 10`

	var snippets []*Snippet
	err := m.DB.Select(&snippets, stmt)
	if err != nil {
		return nil, err
	}

	return snippets, nil
}
