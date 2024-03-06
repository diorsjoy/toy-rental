package mysql

import (
	"database/sql"
	"errors"
	"github.com/oynaToys/pkg/models"
)

type ToysModel struct {
	DB *sql.DB
}

func (m *ToysModel) Latest() ([]*models.Toy, error) {
	// Write the SQL statement we want to execute.
	stmt := `SELECT id, name, description, tokens FROM toys
ORDER BY id ASC`
	// Use the Query() method on the connection pool to execute our
	// SQL statement. This returns a sql.Rows resultset containing the result of
	// our query.
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	toys := []*models.Toy{}
	for rows.Next() {
		s := &models.Toy{}

		err = rows.Scan(&s.ID, &s.Name, &s.Description, &s.Tokens)
		if err != nil {
			return nil, err
		}
		toys = append(toys, s)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	// If everything went OK then return the Snippets slice.
	return toys, nil
}

func (m *ToysModel) GetToys() ([]*models.Toy, error) {
	stmt := `SELECT id, name, description, tokens FROM toys`
	rows, err := m.DB.Query(stmt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	toys := []*models.Toy{}

	for rows.Next() {
		t := &models.Toy{}

		err := rows.Scan(&t.ID, &t.Name, &t.Description, &t.Tokens)
		if err != nil {
			return nil, err
		}
		toys = append(toys, t)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return toys, nil

}

func (m *ToysModel) GetToy(id int) (*models.Toy, error) {
	stmt := `SELECT id, name, description, tokens FROM toys
WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	t := &models.Toy{}

	err := row.Scan(&t.ID, &t.Name, &t.Description, &t.Tokens)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return t, nil

}

func (m *ToysModel) InsertToy(name, desc, tokens string) (int, error) {
	stmt := `INSERT INTO toys (name, description, tokens)
VALUES(?, ?, ?)`

	result, err := m.DB.Exec(stmt, name, desc, tokens)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}
