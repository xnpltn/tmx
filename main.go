package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"

	"speadshets/spdb/application"
	"speadshets/spdb/database"
)

var (
	port     uint64
	sqliteDb string
)

// init function runs before main
func init() {
	flag.Uint64Var(&port, "port", 6969, "port for server to listen on")
	if os.Getenv("SQLITE_DB") != "" {
		sqliteDb = os.Getenv("SQLITE_DB")
	} else {
		sqliteDb = "store.db"
	}
	database.ConnectDB(sqliteDb)
}

func main() {
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	app := application.New(port)
	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
