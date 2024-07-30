package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"speadshets/spdb/core"
	"speadshets/spdb/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// app should implement core.App
var _ core.App = (*app)(nil)

type app struct {
	router *echo.Echo
	db     *gorm.DB
	port   uint64
}

// creates new app
func New(port uint64) *app {
	return &app{
		router: echo.New(),
		db:     database.DB(),
		port:   port,
	}
}

// statrts the app
func (a *app) Start(ctx context.Context) error {
	server := http.Server{
		Handler: a.router,
		Addr:    fmt.Sprintf(":%d", a.port),
	}
	errChan := make(chan error, 1)
	go func() {
		log.Println("Server listening on port: ", a.port)
		fmt.Println(a.db)
		err := server.ListenAndServe()
		if err != nil {
			errChan <- err
		}
	}()
	select {
	case <-errChan:
		err := <-errChan
		return err
	case <-ctx.Done():
		err := a.Shutdown(&server)
		if err != nil {
			log.Println("Failed to gracefully shutdown the server")
			log.Println(" Cause error: ", err.Error())
			log.Println("Forcefull shutting down the server")
			os.Exit(1)
		}
	}
	return nil
}

// shuts down the app
func (a *app) Shutdown(server *http.Server) error {
	log.Println("gracefully shutting down ...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return server.Shutdown(ctx)
}

// exposes the db instance of the app
func (a *app) DB() *gorm.DB {
	return a.db
}

// exposes the router of the app
func (a *app) Router() *echo.Echo {
	return a.router
}
