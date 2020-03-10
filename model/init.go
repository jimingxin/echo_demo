package model

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
)

var db *sqlx.DB
var Err error
func init()  {
	db, Err = sqlx.Open(`mysql`, `root:root123@tcp(127.0.0.1:3306)/echo_demo?charset=utf8&parseTime=true`)
	db.Ping()
}
