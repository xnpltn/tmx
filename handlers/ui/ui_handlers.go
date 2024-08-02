package handlers

import (
	"fmt"

	"spreadsheet/spdb/core"
	"spreadsheet/spdb/models"
	"spreadsheet/spdb/ui"
	"spreadsheet/spdb/ui/views"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// render view utility
func RenderView(c echo.Context, view templ.Component, layoutPath string) error {
	if c.Request().Header.Get("Hx-Request") == "true" {
		return view.Render(c.Request().Context(), c.Response())
	}
	return ui.Layout(layoutPath).Render(c.Request().Context(), c.Response())
}

// home page
func HomePage(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		var sheets []models.Sheet
		app.DB().Find(&sheets)
		return RenderView(c, views.HomeView(sheets), "/")
	}
}

// edit sheet page
func EditPage(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		var sheet models.Sheet
		app.DB().First(&sheet, c.Param("id"))
		return RenderView(c, views.EditView(sheet), fmt.Sprintf("/edit/%d", sheet.ID))
	}
}
