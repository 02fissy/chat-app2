package models

import (
	"database/sql"
	"time"
)

type Messages struct{
	ID int
	UserID int
	Username string
	Content string
	CreatedAt time.Time
	RoomID int
}
type MessageModel struct {
	DB *sql.DB
}
type MessageModelInterface interface{
	Insert(roomID, userID int, content string) error
	GetByRoomID(roomID int) ([]Messages, error)
}
func (m *MessageModel) Insert(roomID, userID int, content string) error {
	stmt := `
		INSERT INTO messages
        (room_id, user_id, content)
        VALUES (?, ?, ?)
	`
	_, err := m.DB.Exec(stmt, roomID, userID, content)
	if err != nil{
		return err
	}
	return nil
}
func (m *MessageModel) GetByRoomID(roomID int) ([]Messages, error) {

	stmt := `
		SELECT
			m.user_id,
			u.username,
			m.content
		FROM messages m
		JOIN users u ON u.user_id = m.user_id
		WHERE m.room_id = ?
		ORDER BY m.id ASC
	`

	rows, err := m.DB.Query(stmt, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []Messages

	for rows.Next() {
		var msg Messages

		err := rows.Scan(
			&msg.UserID,
			&msg.Username,
			&msg.Content,
		)
		if err != nil {
			return nil, err
		}

		messages = append(messages, msg)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}