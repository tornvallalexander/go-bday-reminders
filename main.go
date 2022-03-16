package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"go-bday-reminders/api"
	db "go-bday-reminders/db/sqlc"
	"go-bday-reminders/utils"
	"log"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("could not load env vars:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("could not connect to db:", err)
	}
	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("could not start new server")
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("couldn't start server:", err)
	}
}
