package application

import (
	"spreadsheet/spdb/core"

	"spreadsheet/spdb/handlers/api"
	handlers "spreadsheet/spdb/handlers/ui"
)

func loadUIRoutes(app core.App) {
	app.Router().GET("/", handlers.HomePage(app))
	app.Router().GET("/edit/:id", handlers.EditPage(app))
}

func loadAPIRoutes(app core.App) {
	apiRoutes := app.Router().Group("/api")
	apiRoutes.POST("/save", api.Save(app))
	apiRoutes.POST("/new", api.CreateNewSpreadSheet(app))
	apiRoutes.GET("/cells/:spreadsheetid", api.Fetch(app))
}
