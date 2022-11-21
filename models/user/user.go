package models_user

import (
	database_mysql "github.com/sureshtamrakar/socials/database/mysql"
)

type Entity struct {
	Email   string `json:"email"`
	Uuid    string `json:"uuid"`
	Name    string `json:"name"`
	Gender  string `json:"gender"`
	Country string `json:"country"`
	Age     int    `json:"age"`
}

func GetUuid(email string) (value Entity, err error) {
	row := database_mysql.EsConn.Conn.QueryRow("SELECT uuid FROM `users` WHERE `email`=?", email)
	err = row.Scan(
		&value.Uuid,
	)
	return value, err
}

func Load(uuid string) (value Entity, err error) {
	row := database_mysql.EsConn.Conn.QueryRow("SELECT email,uuid,name,gender,country,age FROM `users` WHERE `uuid`=?", uuid)
	err = row.Scan(
		&value.Email,
		&value.Uuid,
		&value.Name,
		&value.Gender,
		&value.Country,
		&value.Age,
	)
	return value, err
}
