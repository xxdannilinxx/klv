package db

import (
	"database/sql"
	"fmt"

	"github.com/xxdannilinxx/klv/utils"
)

// Connects to the database
func ConnectDB(Config utils.Config) *sql.DB {
	psglconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Config.POSTGRES_HOST, Config.POSTGRES_PORT, Config.POSTGRES_USER, Config.POSTGRES_PASSWORD, Config.POSTGRES_DB)
	db, err := sql.Open("postgres", psglconn)
	utils.CheckError(err)

	return db
}
