package mocks

import (
	"database/sql"
	"snippetbox.ogidi.net/internal/models"
	"time"
)

var mockSnippets = models.Snippet{
	ID:      1,
	Title:   "An old silent pond",
	Content: "An old silent pond..",
	Created: time.Now(),
	Expires: time.Now(),
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

func (m *SnippetModel) Get(id int) (models.Snippet, error) {
	switch id {
	case 1:
		return mockSnippets, nil
	default:
		return models.Snippet{}, models.ErrorNoRecord
	}
}

func (m *SnippetModel) Latest() ([]models.Snippet, error) {
	return []models.Snippet{}, nil
}
