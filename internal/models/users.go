package models

import (
	"database/sql"
	"errors"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"modernc.org/sqlite"
)

type Users struct {
	ID       int
	Username string
	PhoneNumber string
	PasswordHash  []byte
}
type UserModel struct {
	DB *sql.DB
}
type UserModelInterface interface{
	Insert(name, phone_no, password_hash string) error
	GetByID(id int) (*Users, error)
	GetByUsername(username string) (*Users, error)
	Authenticate(name, password_hash string) (int64, error)
}
func(m *UserModel) Insert(name, phone_no, password_hash string) error{
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password_hash), 12)
	if err != nil {
		return err
	}
	stmt := `INSERT INTO users(username, phone_number, password_hash) 
	VALUES (?, ?, ?)`
	_, err = m.DB.Exec(stmt, name, phone_no, passwordHash)
	if err != nil{
		   var sqliteError *sqlite.Error
        if errors.As(err, &sqliteError) {
            if strings.Contains(sqliteError.Error(), "users_uc_phone") {
                return ErrDuplicatePhone
            }
        }
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
func(m *UserModel) Authenticate(name,  password_hash string) (int64, error) {
	var id int64
	var passwordHash []byte
	stmt := `SELECT user_id, password_hash FROM users WHERE username = ?`
	err := m.DB.QueryRow(stmt, name).Scan(&id, &passwordHash)
	if err != nil{
		if errors.Is(err, sql.ErrNoRows) {
            return 0, ErrInvalidCredentials
        } else {
            return 0, err
        }
	}
	err = bcrypt.CompareHashAndPassword(passwordHash, []byte(password_hash))
    if err != nil {
        if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
            return 0, ErrInvalidCredentials
        } else {
            return 0, err
        }
    }
	
	return id, nil
 }