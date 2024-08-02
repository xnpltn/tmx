package main

import (
	"context"
	"embed"
	"flag"
	"log"
	"os"
	"os/signal"

	"spreadsheet/spdb/application"
	"spreadsheet/spdb/database"
)

var (
	PORT      uint64
	SQLITE_DB string
)

// static files

//go:embed all:static
var staticAssets embed.FS

// init function runs before main
func init() {
	flag.Uint64Var(&PORT, "port", 6969, "port for server to listen on")
	if os.Getenv("SQLITE_DB") != "" {
		SQLITE_DB = os.Getenv("SQLITE_DB")
	} else {
		SQLITE_DB = "store.db"
	}
	database.ConnectDB(SQLITE_DB)
}

func main() {
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	app := application.New(PORT, staticAssets)
	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
}
