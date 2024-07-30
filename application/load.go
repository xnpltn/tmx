package application

import (
	"speadshets/spdb/core"
	handlers "speadshets/spdb/handlers/ui"
)

func loadUIRoutes(app core.App) {
	app.Router().GET("/", handlers.HomePage(app))
}
