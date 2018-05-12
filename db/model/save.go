package model

import (
	"database/sql"

	"github.com/sundayfun/go-web/db"
	"github.com/sundayfun/go-web/util"
)

func SaveCompany(d *Company) {
	if d.Exists() {
		return
	}
	const sqlstr = `INSERT INTO homework.company (` +
		`user, money` +
		`) VALUES (` +
		`?, ?` +
		`)`
	if d.Money.Float64 > 50 {
		s := &Self{}
		s.Money = sql.NullFloat64{d.Money.Float64 - 50, true}
		d.Money = sql.NullFloat64{50, true}
		s.User = d.User
		SaveSelf(s)
	}
	// run query
	XOLog(sqlstr, d.User, d.Money)
	res, err := db.GlobalDB.Exec(sqlstr, d.User, d.Money)
	util.CheckErr(err)

	// retrieve id
	id, err := res.LastInsertId()
	util.CheckErr(err)

	// set primary key and existence
	d.AutoID = int(id)
	d._exists = true
	d._deleted = false
	// get the time
	d, err = CompanyByAutoID(db.GlobalDB, d.AutoID)
	util.CheckErr(err)
}

func SaveSelf(s *Self) {
	if s.Exists() {
		return
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO homework.self (` +
		`user, money` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, s.User, s.Money, s.Time)
	res, err := db.GlobalDB.Exec(sqlstr, s.User, s.Money)
	util.CheckErr(err)

	// retrieve id
	id, err := res.LastInsertId()
	util.CheckErr(err)

	// set primary key and existence
	s.AutoID = int(id)
	s._exists = true
	s._deleted = false

	s, err = SelfByAutoID(db.GlobalDB, s.AutoID)
	util.CheckErr(err)
}

func SumCompanyByUser(username string) (res float64, err error) {
	const sqlstr = `select sum(money) from homework.company where user = ?`
	XOLog(sqlstr, username)
	err = db.GlobalDB.QueryRow(sqlstr, username).Scan(&res)
	return res, err
}

func SumSelfByUser(username string) (res float64, err error) {
	const sqlstr = `select sum(money) from homework.self where user = ?`
	XOLog(sqlstr, username)
	err = db.GlobalDB.QueryRow(sqlstr, username).Scan(&res)
	return res, err
}
