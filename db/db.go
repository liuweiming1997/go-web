package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sundayfun/go-web/tool"
)

var GlobalDB *sqlx.DB

func init() {
	conf, err := getDBConfig()
	tool.CheckErr(err)

	source := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true&sql_mode=ansi",
		conf.DB_User,
		conf.DB_Password,
		conf.DB_Host,
		conf.DB_Port,
		conf.DB_Name)
	db, err := sqlx.Connect("mysql", source)
	tool.CheckErr(err)
	GlobalDB = db
}
