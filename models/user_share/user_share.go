package models_user_share

import (
	"fmt"

	database_mysql "github.com/sureshtamrakar/socials/database/mysql"
	models_user_post "github.com/sureshtamrakar/socials/models/user_post"
)

type Entity struct {
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func Create(id int, uuid string) error {
	stmt, err := database_mysql.EsConn.Conn.Prepare("INSERT INTO post_shares (post_id, user_uuid) VALUES (?,?)")
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

func LoadAll(User_Uuid string) (contents []models_user_post.Entity, err error) {
	query := fmt.Sprintf(`SELECT author,title,description FROM posts INNER JOIN post_shares ON posts.id = post_shares.post_id WHERE(post_shares.user_uuid= "%s")`, User_Uuid)
	fmt.Println(query)
	stmt, err := database_mysql.EsConn.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	for stmt.Next() {
		var content models_user_post.Entity

		err = stmt.Scan(
			&content.Author,
			&content.Title,
			&content.Description,
		)

		contents = append(contents, content)
	}

	return contents, err
}
