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
	apiRoutes.POST("/savesheet", api.SaveSheetData(app))
	// apiRoutes.GET("/cells/:spreadsheetid", api.Fetch(app))
	apiRoutes.POST("/newsheet", api.CreateNewSpreadSheet(app))
	apiRoutes.POST("/delete-sheet/:id", api.DeleteSheet(app))
	apiRoutes.POST("/cell", api.SaveCellData(app))
}
