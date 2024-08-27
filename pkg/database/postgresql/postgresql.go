package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

const (
	HOST        = "HOST"
	PORT        = "PORT"
	USER_DB     = "USER_DB"
	PASSWORD_DB = "PASSWORD_DB"
	DB_NAME     = "DB_NAME"
)

func OpenConnDB() (*sql.DB, error) {
	host := os.Getenv(HOST)
	port := os.Getenv(PORT)
	userDB := os.Getenv(USER_DB)
	passwordDB := os.Getenv(PASSWORD_DB)
	dbName := os.Getenv(DB_NAME)
	dsnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, userDB, passwordDB, dbName)

	db, err := sql.Open("postgres", dsnString)
	if err != nil {
		fmt.Printf("unable connect to db")
	}

	err = db.Ping()

	return db, err
}
