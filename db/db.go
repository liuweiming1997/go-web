package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sundayfun/go-web/logs"
)

var GlobalDB *sqlx.DB

func init() {
	conf, err := getDBConfig()
	logs.CheckErr(err)
	source := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true&sql_mode=ansi",
		conf.DB_User,
		conf.DB_Password,
		conf.DB_Host,
		conf.DB_Port,
		conf.DB_Name)
	fmt.Println(source)
	db, err := sqlx.Connect("mysql", source)
	logs.CheckErr(err)
	GlobalDB = db
}
