package storage

import (
	"auth-service/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DB struct {
	conn *sql.DB
}

func NewDB(cf *config.Config) (*DB, error) {
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", cf.DB_USER, cf.DB_PASSWORD, cf.DB_HOST, cf.DB_PORT, cf.DB_NAME)
	conn, err := sql.Open("postgres", dbConn)
	if err != nil {
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		return nil, err
	}
	return &DB{conn: conn}, nil
}
