package handlers

import (
	"fmt"
	"net/http"

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

		// returns sheet with every accociated data
		app.DB().Preload("Rows", "sheet_id = ?", c.Param("id")).Preload("Rows.Cells").Preload("Titles", "sheet_id = ?", c.Param("id")).First(&sheet, c.Param("id"))
		if sheet.ID < 1 {
			return RenderView(c, views.NotFound(), fmt.Sprintf("/edit/%d", sheet.ID))
		}
		return RenderView(c, views.EditView(sheet), fmt.Sprintf("/edit/%d", sheet.ID))
	}
}

/*
<td

	id="data-cell"
	class="p-2 editable"
	contenteditable="true"
	hx-swap="none"
	hx-trigger="input"
	hx-post="/api/cell"
	hx-vals={ fmt.Sprintf(`js:{value: getCellData(event), rowId:%d, cellId:%d, sheetId:%d}`, row.ID, cell.ID, sheet.ID) }>
	  ${
	  fmt.Sprintf(" %v", cell.Value) }

</td>
*/
func AddNewRow(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		cellsHTMl := ""
		rowHTML := "<tr>%s</tr>"
		type reqParams struct {
			SheetId uint32 `json:"sheetId" param:"sheetId" query:"sheetId" form:"sheetId"`
		}
		var sheet models.Sheet
		var params reqParams
		var newRow models.Row
		c.Bind(&params)
		// app.DB().Preload("Titles", "sheet_id = ?", params.SheetId).Preload("Rows", "sheet_id", params.SheetId).Find(&sheet)
		app.DB().Preload("Rows", "sheet_id = ?", params.SheetId).Preload("Rows.Cells").Preload("Titles", "sheet_id = ?", params.SheetId).First(&sheet, params.SheetId)
		sheet.Rows = append(sheet.Rows, newRow)
		if res := app.DB().Save(&sheet); res.Error != nil {
			fmt.Println(res.Error)
		} else {
			fmt.Println("savedd successfully")
		}

		var afterSaveSheet models.Sheet
		app.DB().Preload("Titles", "sheet_id = ?", params.SheetId).Preload("Rows", "sheet_id", params.SheetId).Find(&afterSaveSheet)
		app.DB().Where("sheet_id = ?", sheet.ID).Last(&newRow)

		// Creating html string for each cell
		for i := 0; i < len(sheet.Titles); i++ {
			newRow.Cells = append(newRow.Cells, models.Cell{})
			app.DB().Save(&newRow)
			var lastCell models.Cell
			app.DB().Where("row_id = ?", newRow.ID).Last(&lastCell)
			cellsHTMl += fmt.Sprintf(`<td id="data-cell" class="p-2 editable" contenteditable="true" hx-swap="none" hx-trigger="input" hx-post="/api/cell" hx-vals="js:{value: getCellData(event), rowId:%d, cellId:%d, sheetId:%d}", "></td>`, newRow.ID, lastCell.ID, sheet.ID)
		}
		rowHTML = fmt.Sprintf(rowHTML, cellsHTMl)
		fmt.Println(rowHTML)
		fmt.Println(len(sheet.Titles))
		return c.HTML(http.StatusOK, rowHTML)
	}
}
