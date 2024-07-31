package application

import (
	"spreadsheet/spdb/core"

	handlers "spreadsheet/spdb/handlers/ui"
)

func loadUIRoutes(app core.App) {
	app.Router().GET("/", handlers.HomePage(app))
	app.Router().GET("/edit/:id", handlers.EditPage(app))
}
