package application

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"spreadsheet/spdb/core"
	"spreadsheet/spdb/database"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// app should implement core.App
var _ core.App = (*app)(nil)

type app struct {
	router      *echo.Echo
	db          *gorm.DB
	port        uint64
	staticFiles embed.FS
}

// creates new app
func New(port uint64, staticFiles embed.FS) core.App {
	return &app{
		router:      echo.New(),
		db:          database.DB(),
		port:        port,
		staticFiles: staticFiles,
	}
}

// statrts the app
func (a *app) Start(ctx context.Context) error {
	server := http.Server{
		Handler: a.router,
		Addr:    fmt.Sprintf(":%d", a.port),
	}
	static, err := fs.Sub(a.staticFiles, "static")
	if err != nil {
		fmt.Println(err)
		return err
	}

	// serve static files
	a.router.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static", http.FileServer(http.FS(static)))))

	// application routes
	loadUIRoutes(a)
	loadAPIRoutes(a)

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
		// graceful termination
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

// shuts down the app gracefully waiting 10 seconds to finish running tasks
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
