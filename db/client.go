package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Client struct {
	DB *sql.DB
}

func Init(dsn string) (Client, error) {
	client := Client{}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return client, err
	}

	err = db.Ping()
	if err != nil {
		return client, err
	}

	client.DB = db
	return client, nil
}
