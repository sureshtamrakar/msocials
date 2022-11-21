package models_user_post

import (
	"fmt"

	database_mysql "github.com/sureshtamrakar/socials/database/mysql"
)

type Response struct {
	UserUuid string `json:"user_uuid"`
}

type Entity struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func Create(author, uuid, title, description string) error {
	stmt, err := database_mysql.EsConn.Conn.Prepare("INSERT INTO posts (author, user_uuid, title, description) VALUES (?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(author, uuid, title, description)
	if err != nil {
		return err
	}
	stmt.Close()
	return nil
}

func LoadAll(UserUuid string) (contents []Entity, err error) {
	query := fmt.Sprintf(`select distinct p.author,title,description from posts p inner join (select * from user_friends  where "%s" in (user_friends.user_uuid, user_friends.friend_uuid)) user_friends on p.user_uuid in (user_friends.user_uuid, user_friends.friend_uuid)`, UserUuid)
	stmt, err := database_mysql.EsConn.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	for stmt.Next() {
		var content Entity

		err = stmt.Scan(
			&content.Author,
			&content.Title,
			&content.Description,
		)

		contents = append(contents, content)
	}

	return contents, err
}
func Load(id int) (value Response, err error) {
	row := database_mysql.EsConn.Conn.QueryRow("SELECT user_uuid FROM `posts` WHERE `id`=?", id)
	err = row.Scan(
		&value.UserUuid,
	)
	return value, err
}

func Like(id int, uuid string) error {
	stmt, err := database_mysql.EsConn.Conn.Prepare("INSERT INTO post_likes (post_id, user_uuid) VALUES (?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(id, uuid)
	if err != nil {
		return err
	}
	stmt.Close()
	return nil
}

func LikeList(UserUuid string) (contents []Entity, err error) {
	query := fmt.Sprintf(`SELECT author,title,description FROM posts INNER JOIN post_likes ON posts.id = post_likes.post_id WHERE(post_likes.user_uuid= "%s")`, UserUuid)
	stmt, err := database_mysql.EsConn.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	for stmt.Next() {
		var content Entity

		err = stmt.Scan(
			&content.Author,
			&content.Title,
			&content.Description,
		)

		contents = append(contents, content)
	}

	return contents, err
}

func LoadPost(User_Uuid string) (contents []Entity, err error) {
	stmt, err := database_mysql.EsConn.Conn.Query("SELECT `author`,`title`,`description` FROM `posts` WHERE `user_uuid`=?", User_Uuid)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	for stmt.Next() {
		var content Entity

		err = stmt.Scan(
			&content.Author,
			&content.Title,
			&content.Description,
		)

		contents = append(contents, content)
	}

	return contents, err
}
