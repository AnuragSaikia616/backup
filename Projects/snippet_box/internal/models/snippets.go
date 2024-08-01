package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(Title string, content string, expires int) (int, error) {
	qry := "INSERT INTO snippets (Title, content, created, expires) VALUES (?,?,UTC_TIMESTAMP(),DATE_ADD(UTC_TIMESTAMP(),INTERVAL ? DAY))"
	result, err := m.DB.Exec(qry, Title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	qry := "SELECT id,Title,content,created,expires FROM snippets WHERE UTC_TIMESTAMP() < expires AND id = ?"

	// this struct stores the snippet which is the be returned
	s := &Snippet{}

	// Execute the query
	row := m.DB.QueryRow(qry, id)

	// Error evaluation
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// if the query returns no rows, the error is of ErrNoRecord type
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		} else {
			return nil, err
		}
	}

	// if everything OK
	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	qry := "SELECT id,Title,content,created,expires FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10"
	LatestSnippetsList := []*Snippet{}

	rows, err := m.DB.Query(qry)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		s := &Snippet{}
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		LatestSnippetsList = append(LatestSnippetsList, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return LatestSnippetsList, nil
}
