package server

import (
	"database/sql"
	"fmt"

	"Skavengerr/coins-crud/pkg/config"
	"Skavengerr/coins-crud/util"
)

func ConnectDB(cfg *config.Config) *sql.DB {
	dbInfo := fmt.Sprintf("user=%s  dbname=%s sslmode=disable password=%s",
		cfg.User, cfg.DbName, cfg.Password)
	db, err := sql.Open("postgres", dbInfo)

	util.CheckErr(err)

	if err := db.Ping(); err != nil {
		util.CheckErr(err)
	}

	return db
}
