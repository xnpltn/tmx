package api

import (
	"fmt"
	"net/http"

	"spreadsheet/spdb/core"
	"spreadsheet/spdb/models"

	"github.com/labstack/echo/v4"
)

type NewSpreadSheet struct {
	Collection string `json:"collection" form:"collection"`
	SheetName  string `json:"sheet" form:"sheetName"`
}

func Save(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		// b, err := io.ReadAll(c.Request().Body)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		var spreadsheet models.Spreadsheet
		var sheet models.Sheet
		var cell models.Cell
		err := c.Bind(&cell)
		if err != nil {
			fmt.Println(err)
		}
		app.DB().Find(&sheet, cell.SheetID)
		app.DB().Find(&spreadsheet, sheet.SpreadsheetID)
		sheet.Cells = append(sheet.Cells, cell)
		spreadsheet.Sheets = append(spreadsheet.Sheets, sheet)
		res := app.DB().Save(&spreadsheet)
		if res.Error != nil {
			fmt.Println(res.Error)
		} else {
			fmt.Println("Updated Successfully")
		}
		return c.JSON(http.StatusOK, `"hello":"wolrd"`)
	}
}

func Fetch(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(200, "")
	}
}

func CreateNewSpreadSheet(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("HX-Redirect", "/")
		var params NewSpreadSheet
		var spreadsheet models.Spreadsheet
		var sheet models.Sheet
		if err := c.Bind(&params); err != nil {
			fmt.Println(err)
			return c.JSON(http.StatusOK, "Something Went Wrong")
		}
		if res := app.DB().Where("name = ?", params.Collection).First(&spreadsheet); res.Error == nil {
			if res := app.DB().Where("name = ?", params.SheetName).Where("spreadsheet_id = ?", spreadsheet.ID).First(&sheet); res.Error == nil {
				return c.JSON(200, "Sheet with that name already exist")
			} else {
				spreadsheet.Sheets = []models.Sheet{{Name: params.SheetName}}
				res := app.DB().Save(&spreadsheet)
				if res.Error != nil {
					fmt.Println(res.Error)
					return c.JSON(200, "Something went wrong")
				} else {
					return c.HTML(200, "Created Successfully")
				}
			}
		} else {
			newSpreadsheet := models.Spreadsheet{Name: params.Collection}
			newSpreadsheet.Sheets = []models.Sheet{{Name: params.SheetName}}
			res := app.DB().Save(&newSpreadsheet)
			if res.Error != nil {
				fmt.Println(res.Error)
				return c.JSON(200, "Something went wrong")
			} else {
				return c.HTML(200, "Created Successfully")
			}
		}
	}
}
