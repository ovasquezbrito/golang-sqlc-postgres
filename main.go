package main

import (
	"github.com/ovasquezbrito/simplebank/util"
	"github.com/ovasquezbrito/simplebank/api"
	db "github.com/ovasquezbrito/simplebank/db/sqlc"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Imposible conectarse a la Base de Dato: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}