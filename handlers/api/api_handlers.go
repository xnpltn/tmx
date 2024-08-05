package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"spreadsheet/spdb/core"
	"spreadsheet/spdb/models"

	"github.com/labstack/echo/v4"
)

type editCellParams struct {
	CellValue string `json:"value" param:"value" query:"value" form:"value"`
	Name      string `json:"name" param:"name" query:"name" form:"name"`
	RowID     uint32 `json:"rowId" param:"rowId" query:"rowId" form:"rowId"`
	CellID    uint32 `json:"cellId" param:"cellId" query:"cellId" form:"cellId"`
	SheetID   uint32 `json:"sheetId" param:"sheetId" query:"sheetId" form:"sheetId"`
}

// saves cell data as the cell changes
func SaveSheetData(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		b, err := io.ReadAll(c.Request().Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("recieved: ", string(b))
		return c.JSON(http.StatusOK, map[string]string{"ok": "ok"})
	}
}

// Fetch data on open of spreadsheet
func Fetch(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

// create new spreasheet
func CreateNewSpreadSheet(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		var sheet models.Sheet

		if b, err := io.ReadAll(c.Request().Body); err == nil {
			if err := json.Unmarshal(b, &sheet); err != nil {
				fmt.Println("Error Unmarshaling: ", err)
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "something went wrong"})
			}
		}

		fmt.Println(sheet.Titles)
		/*

			        <option value="Text">Text</option>
			        <option value="CheckBox">CheckBox</option>
			        <option value="Number">Number</option>
			        <option value="Status">Status</option>
			        <option value="Tag">Tag</option>
			        <option value="Date">Date</option>
			        <option value="Label">Label</option>

			const (
				Text DataType = iota
				CheckBox
				Status
				Tag
				Date
				Number
				Label
			)
		*/
		for _, title := range sheet.Titles {
			switch title.DataTypeString {
			case "Text":
				title.DataType = models.Text
			case "CheckBox":
				title.DataType = models.CheckBox
			case "Number":
				title.DataType = models.Number
			case "Status":
				title.DataType = models.Status
			case "Tag":
				title.DataType = models.Tag
			case "Date":
				title.DataType = models.Date
			case "Label":
				title.DataType = models.Label
			default:
				title.DataType = models.Text
			}
		}
		if res := app.DB().Create(&sheet); res.Error != nil {
			fmt.Println("error saving: ", res.Error)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "something went wrong"})
		}

		return c.JSON(http.StatusOK, map[string]string{"error": "something went wrong"})
	}
}

func DeleteSheet(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

// saving cell data, runns on every edit of the cell
func SaveCellData(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		var params editCellParams
		c.Bind(&params)

		if params.CellID == 0 {
			return c.JSON(http.StatusBadRequest, "error")
		} else {
			var cell models.Cell
			app.DB().Where("id = ? AND row_id = ? ", params.CellID, params.RowID).First(&cell, params.CellID)
			fmt.Println(cell.ID == uint(params.CellID))
			cell.Value = params.CellValue
			if res := app.DB().Save(&cell); res.Error != nil {
				fmt.Println("error saving")
				fmt.Println(res.Error)
			} else {
				fmt.Println("saved successfully")
			}

		}
		return c.JSON(http.StatusOK, map[string]string{"me": "you"})
	}
}
