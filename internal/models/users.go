package models

import "database/sql"

type Users struct {
	ID       int
	Username string
}
type UserModel struct {
	DB *sql.DB
}
type UserModelInterface interface{
	Insert(username string) error
	GetByID(id int) (*Users, error)
	GetByUsername(username string) (*Users, error)
}
func (m *UserModel) Insert(username string) error {
	stmt:= `INSERT INTO users (username) VALUES (?)`
	_, err := m.DB.Exec(stmt, username)
	if err != nil {
		return err
	}
	return nil
}
func(m *UserModel) GetByID(id int) (*Users, error) {
	stmt := `SELECT user_id, username FROM users WHERE user_id = ?`
	row := m.DB.QueryRow(stmt, id)
	u := &Users{}
	err := row.Scan(&u.ID, &u.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}
func(m *UserModel) GetByUsername(username string) (*Users, error) {
	stmt := `SELECT user_id, username FROM users WHERE username = ?`
	row := m.DB.QueryRow(stmt, username)
	u := &Users{}
	err := row.Scan(&u.ID, &u.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return u, nil
}
