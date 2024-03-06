package mysql

import (
	"database/sql"
	"errors"

	// Import the models package that we just created. You need to prefix this with
	// whatever module path you set up back in chapter 02.02 (Project Setup and Enabling
	// Modules) so that the import statement looks like this:
	// "{your-module-path}/pkg/models".
	"github.com/oynaToys/pkg/models"
)

// Define a FeedbacksModel type which wraps a sql.DB connection pool.
type FeedbacksModel struct {
	DB *sql.DB
}

func (m *FeedbacksModel) InsertFeedback(name string, content string, stars string) (int, error) {
	stmt := `INSERT INTO feedbacks (name, content, stars)
VALUES(?, ?, ?)`
	result, err := m.DB.Exec(stmt, name, content, stars)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *FeedbacksModel) ShowFeedbacks() ([]*models.Feedback, error) {
	stmt := `SELECT id, name, content, stars FROM feedbacks
ORDER BY id ASC`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	snippets := []*models.Feedback{}
	for rows.Next() {
		s := &models.Feedback{}
		err = rows.Scan(&s.ID, &s.Name, &s.Content, &s.Stars)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return snippets, nil
}

func (m *FeedbacksModel) GetFeedback(id int) (*models.Feedback, error) {
	stmt := `SELECT id, name, content, stars FROM feedbacks
WHERE id = ?`
	row := m.DB.QueryRow(stmt, id)
	s := &models.Feedback{}
	err := row.Scan(&s.ID, &s.Name, &s.Content, &s.Stars)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}
