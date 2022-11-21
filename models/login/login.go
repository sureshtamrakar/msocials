package models_login

import (
	database_mysql "github.com/sureshtamrakar/socials/database/mysql"
)

type Entity struct {
	Email    string `json:"email"`
	Uuid     string `json:"uuid"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Country  string `json:"country"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

func Login(email string) (login Entity, err error) {
	row := database_mysql.EsConn.Conn.QueryRow("SELECT email, password FROM `users` WHERE `email`=?", email)
	err = row.Scan(
		&login.Email,
		&login.Password,
	)
	return login, err
}

func Create(email, uuid, name, gender, country, password string, age int) error {
	stmt, err := database_mysql.EsConn.Conn.Prepare("INSERT INTO users (email, uuid, name, gender, country, age, password) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(email, uuid, name, gender, country, age, password)
	if err != nil {
		return err
	}
	stmt.Close()
	return nil
}
