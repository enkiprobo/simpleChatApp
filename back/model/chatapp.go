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
	getChatFriends string = `
		SELECT
			user_id,
			username,
			password
		FROM
			user
		WHERE
			username <> ?
	`
	getChatDetailQuery string = `
		SELECT
			cd_id,
			message,
			username,
			create_time
		FROM 
			chat_detail
				INNER JOIN 
					user
				ON 
					chat_detail.user_id = user.user_id
		WHERE
			chatroom = ?
		ORDER BY
			create_time 
	`
	getChatRoomQuery string = `
		SELECT
			cr_id 
		FROM
			chat_room
		WHERE
			(user1 = ? AND user2 = ?)
			OR
			(user1 = ? AND user2 = ?)
	`
	insertChatQuery string = `
		INSERT INTO
			chat_detail (message, chatroom, user_id)
		VALUES
			(?,?,?)
	`
)

type (
	User struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Password string `json:"-"`
	}
	ChatDetail struct {
		ID         int    `json:"chat_id"`
		Username   string `json:"message_author"`
		Message    string `json:"message"`
		ChatCreate string `json:"create_date"`
	}
)

var (
	ChatDB *sql.DB
)

func InitDB() error {
	var err error

	ChatDB, err = sql.Open("sqlite3", "etc/database/chatapp.db")
	if err != nil {
		return err
	}

	return nil
}

func GetUser(uname string) (User, error) {
	var id int
	var username, password string

	row := ChatDB.QueryRow(getUserQuery, uname)

	row.Scan(&id, &username, &password)

	return User{
		ID:       id,
		Username: username,
		Password: password,
	}, nil
}

func GetUsers(uname string) ([]User, error) {

	var friends = []User{}

	rows, err := ChatDB.Query(getChatFriends, uname)
	if err != nil {
		return friends, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var username, password string

		err := rows.Scan(&id, &username, &password)
		if err != nil {
			return friends, err
		}

		friend := User{
			ID:       id,
			Username: username,
			Password: password,
		}

		friends = append(friends, friend)
	}

	return friends, nil
}

func GetChatRoom(uname1, uname2 int) int {
	var id int

	row := ChatDB.QueryRow(getChatRoomQuery, uname1, uname2, uname2, uname1)

	row.Scan(&id)

	return id
}

func InsertChat(message string, chatroom, userid1 int) (int64, error) {

	var chatDetailID int64

	r, err := ChatDB.Exec(insertChatQuery, message, chatroom, userid1)
	if err != nil {
		return chatDetailID, err
	}

	chatDetailID, err = r.LastInsertId()
	if err != nil {
		return chatDetailID, err
	}

	return chatDetailID, nil
}

func GetChatDetail(chatroom int) ([]ChatDetail, error) {

	chatDetails := []ChatDetail{}

	rows, err := ChatDB.Query(getChatDetailQuery, chatroom)
	if err != nil {
		return chatDetails, err
	}
	defer rows.Close()

	for rows.Next() {
		var chatDetail ChatDetail
		rows.Scan(&chatDetail.ID, &chatDetail.Message, &chatDetail.Username, &chatDetail.ChatCreate)

		chatDetails = append(chatDetails, chatDetail)
	}

	return chatDetails, nil
}
