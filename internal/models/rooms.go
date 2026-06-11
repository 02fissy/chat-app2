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
	GetOrCreate(name string) (int, error)
	GetAll()  ([]Rooms, error)
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
func (m *RoomModel) GetOrCreate(name string) (int, error) {

	var id int

	err := m.DB.QueryRow(
		"SELECT id FROM rooms WHERE name = ?",
		name,
	).Scan(&id)

	if err == nil {
		return id, nil
	}

	if err != sql.ErrNoRows {
		return 0, err
	}

	result, err := m.DB.Exec(
		"INSERT INTO rooms(name) VALUES(?)",
		name,
	)
	if err != nil {
		return 0, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastID), nil
}
func (m *RoomModel) GetAll() ([]Rooms, error) {

	stmt := `
		SELECT id, name, created_at
		FROM rooms
		ORDER BY created_at DESC
	`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []Rooms

	for rows.Next() {

		var room Rooms

		err := rows.Scan(
			&room.ID,
			&room.Name,
			&room.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		rooms = append(rooms, room)
	}

	return rooms, nil
}