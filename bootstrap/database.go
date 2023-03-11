package bootstrap

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func newMysqlDatabase(env *Env) *sql.DB {
	db, err := sql.Open("mysql", env.DBUser+":"+env.DBPass+"@tcp("+env.DBHost+":"+env.DBPort+")/"+env.DBName+"?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
