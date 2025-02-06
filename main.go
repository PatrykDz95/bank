package main

import (
	"bank/api"
	db "bank/db/sqlc"
	"bank/util"
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	server.Start(config.ServerAddcress)
}
