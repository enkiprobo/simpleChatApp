package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	getUserQuery string = `
		SELECT
			user_id,
			username,
			password
		FROM 
			user
		WHERE
			username = ?

	`
	getChatDetailQuery string = `
		SELECT
			cd_id,
			message,
			chat_create_date
		WHERE
			chatroom = ?
		ORDER BY
			chat_create_date
			
	`
	insertChatQuery string = `
		SELECT
			
			
	`
)

type (
	User struct {
		ID       int    `db:"user_id" json:"id"`
		Username string `db:"username" json:"username"`
		Password string `db:"password" json:"password"`
	}
)

var (
	chatDB *sql.DB
)

func InitDB() error {
	var err error

	chatDB, err = sql.Open("sqlite3", "etc/database/chatapp.db")
	if err != nil {
		return err
	}

	return nil
}

func GetUser(uname string) (User, error) {
	var id int
	var username, password string

	row := chatDB.QueryRow(getUserQuery, uname)

	err := row.Scan(&id, &username, &password)
	if err != nil {
		return User{}, err
	}

	return User{
		ID:       id,
		Username: username,
		Password: password,
	}, nil
}
