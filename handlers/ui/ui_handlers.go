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
		// var cells []models.Cell
		var rows []models.Row
		app.DB().First(&sheet, c.Param("id"))
		app.DB().Where("sheet_id = ?", sheet.ID).Find(&sheet.Rows)
		app.DB().Where("sheet_id = ?", sheet.ID).Find(&sheet.Tittles)
		app.DB().Where("sheet_id = ?", sheet.ID).Find(&rows)
		for i := 0; i < len(rows); i++ {
			app.DB().Where("row_id = ?", rows[i].ID).Find(&rows[i].Cells)
		}
		sheet.Rows = append(sheet.Rows, rows...)
		return RenderView(c, views.EditView(sheet), fmt.Sprintf("/edit/%d", sheet.ID))
	}
}
