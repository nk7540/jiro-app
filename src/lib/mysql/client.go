package mysql

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Client - Mysql client
type Client struct {
	DB *sql.DB
}

// NewClient - connects to mysql server
func NewClient(ctx context.Context, user string, password string, host string, port string, database string) (*Client, error) {
	addr := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database

	db, err := sql.Open("mysql", addr)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(300 * time.Second)

	return &Client{db}, nil
}
