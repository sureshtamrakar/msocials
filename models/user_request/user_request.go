package models_user_request

import (
	database_mysql "github.com/sureshtamrakar/socials/database/mysql"
)

type Entity struct {
	User_Uuid string `json:"user_uuid"`
}

func Create(User_Uuid, Request_Uuid string) error {

	stmt, _ := database_mysql.EsConn.Conn.Prepare("INSERT INTO user_requests_rel (user_uuid, request_uuid) VALUES (?,?)")
	_, err := stmt.Exec(User_Uuid, Request_Uuid)
	if err != nil {
		return err
	}
	stmt.Close()
	return nil

}

func List(User_Uuid string) (contents []Entity, err error) {
	stmt, err := database_mysql.EsConn.Conn.Query("SELECT `User_Uuid` FROM `user_requests_rel` WHERE `request_uuid`=?", User_Uuid)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	for stmt.Next() {
		var content Entity

		err = stmt.Scan(
			&content.User_Uuid,
		)

		contents = append(contents, content)
	}

	return contents, err
}

//check if friend request is sent and available in DB
func Load(UserUuid, RequestUuid string) (count int, err error) {
	var val int
	err = database_mysql.EsConn.Conn.QueryRow("SELECT COUNT(*) FROM `user_requests_rel` WHERE `user_uuid`=? AND `request_uuid`=? AND `status`= 0", RequestUuid, UserUuid).Scan(&val)
	if err != nil {

		return val, err
	}
	return val, nil
}

func Update(UserUuid, RequestUuid string) (err error) {
	rows, err := database_mysql.EsConn.Conn.Query("UPDATE `user_requests_rel` SET `request_uuid` = ?, `user_uuid` = ?, `status` =? WHERE `request_uuid`=? AND `user_uuid` =?", UserUuid, RequestUuid, 1, UserUuid, RequestUuid)
	if err != nil {
		return err
	}
	rows.Close()
	return err
}
