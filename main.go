package main

import (
	"github.com/ovasquezbrito/simplebank/api"
	db "github.com/ovasquezbrito/simplebank/db/sqlc"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@172.17.0.2:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)


func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Imposible conectarse a la Base de Dato: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}