package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	connStr := "user=postgres dbname=go_user_api sslmode=disable password=root"
	return sql.Open("postgres", connStr)
}
