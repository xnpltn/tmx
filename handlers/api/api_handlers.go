package api

import (
	"fmt"
	"net/http"

	"spreadsheet/spdb/core"
	"spreadsheet/spdb/models"

	"github.com/labstack/echo/v4"
)

type NewSpreadSheetParams struct {
	Collection string `json:"collection" form:"collection"`
	SheetName  string `json:"sheet" form:"sheetName"`
}

// saves cell data as the cell changes
func Save(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		var cellParams models.Cell
		var cell models.Cell

		var sheet models.Sheet
		err := c.Bind(&cellParams)
		if err != nil {
			fmt.Println(err)
		}
		if res := app.DB().Where("name = ?", cellParams.Name).Where("sheet_id = ?", cellParams.SheetID).First(&cell); res.Error == nil {

			cell.Value = cellParams.Value
			cell.RowNumber = cellParams.RowNumber
			cell.Name = cellParams.Name
			cell.ColumnNumber = cellParams.ColumnNumber

			res := app.DB().Save(&cell)
			if res.Error != nil {
				fmt.Println(res.Error)
			} else {
				fmt.Println("Updated Successfully")
			}

		} else {
			app.DB().First(&sheet, cellParams.SheetID)
			sheet.Cells = append(sheet.Cells, cellParams)
			res := app.DB().Save(&sheet)
			if res.Error != nil {
				fmt.Println(res.Error)
			}
		}
		return c.JSON(http.StatusOK, `"hello":"wolrd"`)
	}
}

// Fetch data on open of spreadsheet
func Fetch(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		var sheet models.Sheet
		var cells []models.Cell
		app.DB().Where("id = ?", c.Param("spreadsheetid")).First(&sheet)
		app.DB().Where("sheet_id = ?", sheet.ID).Find(&cells)

		return c.JSON(http.StatusOK, cells)
	}
}

// create new spreasheet
func CreateNewSpreadSheet(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("HX-Redirect", "/")
		var params NewSpreadSheetParams
		var spreadsheet models.Spreadsheet
		var sheet models.Sheet
		if err := c.Bind(&params); err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusOK, "Something Went Wrong")
		}
		if res := app.DB().Where("name = ?", params.Collection).First(&spreadsheet); res.Error == nil {
			if res := app.DB().Where("name = ?", params.SheetName).Where("spreadsheet_id = ?", spreadsheet.ID).First(&sheet); res.Error == nil {
				return c.JSON(http.StatusOK, "Sheet with that name already exist")
			} else {
				spreadsheet.Sheets = []models.Sheet{{Name: params.SheetName}}
				res := app.DB().Save(&spreadsheet)
				if res.Error != nil {
					fmt.Println(res.Error)
					return c.JSON(http.StatusOK, "Something went wrong")
				} else {
					return c.HTML(http.StatusOK, "Created Successfully")
				}
			}
		} else {
			newSpreadsheet := models.Spreadsheet{Name: params.Collection}
			newSpreadsheet.Sheets = []models.Sheet{{Name: params.SheetName}}
			res := app.DB().Save(&newSpreadsheet)
			if res.Error != nil {
				fmt.Println(res.Error)
				return c.JSON(http.StatusOK, "Something went wrong")
			} else {
				return c.HTML(http.StatusOK, "Created Successfully")
			}
		}
	}
}
