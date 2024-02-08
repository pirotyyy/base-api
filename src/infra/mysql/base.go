package mysql

import (
	"database/sql"
	"os"
)

type DBConnection struct {
	Conn *sql.DB
}

func NewDBConnection() (*DBConnection, error) {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PW")
	dbName := os.Getenv("DB_NAME")
	dbAddr := os.Getenv("DB_ADDR")
	conn, err := sql.Open("mysql", dbUser+":"+dbPass+"@("+dbAddr+")/"+dbName)
	if err != nil {
		return nil, err
	}

	return &DBConnection{Conn: conn}, nil
}
