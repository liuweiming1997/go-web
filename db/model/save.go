package model

import (
	"github.com/sundayfun/go-web/db"
	"github.com/sundayfun/go-web/logs"
)

func SaveOut(d *Demo) {
	const sqlstr = `INSERT INTO homework.demo (` +
		`user, money` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, d.User, d.Money)
	res, err := db.GlobalDB.Exec(sqlstr, d.User, d.Money)
	logs.CheckErr(err)

	// retrieve id
	id, err := res.LastInsertId()
	logs.CheckErr(err)

	// set primary key and existence
	d.AutoID = int(id)
	d._exists = true
	d._deleted = false
	// get the time
	d, err = DemoByAutoID(db.GlobalDB, d.AutoID)
	logs.CheckErr(err)
}
