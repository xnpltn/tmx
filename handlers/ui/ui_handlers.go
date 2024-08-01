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

func RenderView(c echo.Context, view templ.Component, layoutPath string) error {
	if c.Request().Header.Get("Hx-Request") == "true" {
		return view.Render(c.Request().Context(), c.Response())
	}
	return ui.Layout(layoutPath).Render(c.Request().Context(), c.Response())
}

func HomePage(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		var sheets []models.Sheet
		res := app.DB().Find(&sheets)
		if res.Error != nil {
			fmt.Println(res.Error)
		}
		fmt.Println(len(sheets))
		return RenderView(c, views.HomeView(sheets), "/")
	}
}

func EditPage(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		sheetID := c.Param("id")
		var cells []models.Cell
		app.DB().Where("sheet_id = ?", sheetID).Find(&cells)
		var sheet models.Sheet
		app.DB().Find(&sheet, sheetID)
		sheet.Cells = cells
		return RenderView(c, views.EditView(sheet), fmt.Sprintf("/edit/%d", sheet.ID))
	}
}
