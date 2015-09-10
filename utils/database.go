package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DatabaseSession struct {
	*sql.DB
}

func NewSession() *DatabaseSession {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/traveller")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return &DatabaseSession{db}
}
