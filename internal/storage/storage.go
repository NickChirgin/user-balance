package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/nickchirgin/user-balance/internal/helpers"
)

type Storage struct {
	Db *sql.DB
}

	var (
		Client   *sql.DB
		username = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASSWORD")
		schema   = os.Getenv("DB_NAME")
		host = os.Getenv("DB_HOST")
	)
func init() {
		connInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host,	username, password, schema)
		var err error
		Client, err = sql.Open("pgx", connInfo)
		if err != nil {
			helpers.CheckErr(err)
		}
		if err = Client.Ping(); err != nil {
			helpers.CheckErr(err)
		} 
		log.Println("Database ready to accept connections")
}