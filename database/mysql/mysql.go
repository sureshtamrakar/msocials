package database_mysql

import (
	"database/sql"
	"fmt"

	"github.com/sureshtamrakar/socials/util"

	_ "github.com/go-sql-driver/mysql"
)

type Connection struct {
	Conn *sql.DB
}

var EsConn Connection

func init() {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s", util.Yamlvalue.DBUsername, util.Yamlvalue.DBPassword, util.Yamlvalue.DBHost, util.Yamlvalue.DBName),
	)
	if err != nil {
		return
	}

	EsConn.Conn = db

}
