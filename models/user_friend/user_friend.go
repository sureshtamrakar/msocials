package models_user_friend

import (
	"fmt"

	database_mysql "github.com/sureshtamrakar/socials/database/mysql"
	models_user "github.com/sureshtamrakar/socials/models/user"
	models_user_post "github.com/sureshtamrakar/socials/models/user_post"
)

func Create(UserUuid, RequestUuid string) error {
	stmt, _ := database_mysql.EsConn.Conn.Prepare("INSERT INTO user_friends (user_uuid, friend_uuid) VALUES (?,?)")
	_, err := stmt.Exec(UserUuid, RequestUuid)
	stmt.Close()
	return err
}

func List(UserUuid string) []models_user.Entity {
	aa, _ := ListA(UserUuid)
	bb, _ := ListB(UserUuid)
	aa = append(aa, bb...)
	return aa

}

func ListA(UserUuid string) (contents []models_user.Entity, err error) {
	query := fmt.Sprintf(`SELECT email,uuid,name,gender,country,age FROM users where users.uuid in (SELECT user_friends.friend_uuid from user_friends WHERE user_friends.user_uuid = "%s")`, UserUuid)
	stmt, err := database_mysql.EsConn.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	for stmt.Next() {
		var content models_user.Entity

		err = stmt.Scan(
			&content.Email,
			&content.Uuid,
			&content.Name,
			&content.Gender,
			&content.Country,
			&content.Age,
		)

		contents = append(contents, content)
	}

	return contents, err
}

func ListB(UserUuid string) (contents []models_user.Entity, err error) {
	query := fmt.Sprintf(`SELECT email,uuid,name,gender,country,age FROM users where users.uuid in (SELECT user_friends.user_uuid from user_friends WHERE user_friends.friend_uuid = "%s")`, UserUuid)
	stmt, err := database_mysql.EsConn.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	for stmt.Next() {
		var content models_user.Entity

		err = stmt.Scan(
			&content.Email,
			&content.Uuid,
			&content.Name,
			&content.Gender,
			&content.Country,
			&content.Age,
		)

		contents = append(contents, content)
	}

	return contents, err
}

func LoadAll(UserUuid string) (contents []models_user_post.Response, err error) {
	query := fmt.Sprintf(`select u.uuid from users u inner join (select * from user_friends where "%s" in (user_friends.user_uuid, user_friends.friend_uuid)) user_friends on u.uuid in (user_friends.user_uuid, user_friends.friend_uuid)`, UserUuid)
	fmt.Println(query)
	stmt, err := database_mysql.EsConn.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	for stmt.Next() {
		var content models_user_post.Response

		err = stmt.Scan(
			&content.UserUuid,
		)

		contents = append(contents, content)
	}

	return contents, err
}
