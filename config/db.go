package config

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)


const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "my-secret-pw-23"
	dbname   = "db_go_sql"
)
var (
	DB *sql.DB
	err error
)
func Connect() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err = sql.Open("postgres", psqlInfo) // untuk memvalidasi argumen-argumen yang diberikan
	if err != nil {
		panic(err)
	}

	return DB
}