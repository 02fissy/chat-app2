package models

import (
	"database/sql"
	"time"
	"errors"
)
 var ErrNoRecord = errors.New("models: no matching record")

type Rooms struct {
	ID        int
	Name      string
	CreatedAt time.Time
}
type RoomModel struct {
	DB *sql.DB
}
type RoomModelInterface interface{
	Insert(name string) error
	Get(id int) (*Rooms, error)
	GetByName(name string) (int, error)
}
func (m *RoomModel) Insert(name string) error {
	stmt := `INSERT INTO rooms (name) VALUES (?)`
	_, err := m.DB.Exec(stmt, name)
	if err != nil {
		return err
	}
	return nil
}
func (m *RoomModel) Get(id int) (*Rooms, error) {

	stmt := `
		SELECT
			id,
			name,
			created_at
		FROM rooms
		WHERE id = ?
	`

	r := &Rooms{}

	err := m.DB.QueryRow(stmt, id).Scan(
		&r.ID,
		&r.Name,
		&r.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}

	return r, nil
}
func (m *RoomModel) GetByName(name string) (int, error) {
	var id int
	stmt := `SELECT id FROM rooms WHERE name = ?`
	err := m.DB.QueryRow(stmt, name).Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, ErrNoRecord
		}
		return 0, err
}
	return id, nil
}