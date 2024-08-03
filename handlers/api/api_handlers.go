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

type NewSpreadSheetParams struct {
	Collection string `json:"collection" form:"collection"`
	SheetName  string `json:"sheet" form:"sheetName"`
}

/*
recieved:  [{"name":"John Doe","amountChart":"[Chart Placeholde","totalMoney":"$123,456.78","platform":"eBay","orderTimes":15,"regi":"2022-4"},{"name":"Jane Smith","amountChart":"[Chart Placeholder]","totalMoney":"$654,321.00","platform":"Etsy","orderTimes":30,"regi":"2023-2"}]

*/

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

// saving cell data
func SaveCellData(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		// sample data
		type editCellParams struct {
			CellValue string `json:"value" param:"value" query:"value" form:"value"`
			Name      string `json:"name" param:"name" query:"name" form:"name"`
			RowID     uint32 `json:"rowId" param:"rowId" query:"rowId" form:"rowId"`
			CellID    uint32 `json:"cellId" param:"cellId" query:"cellId" form:"cellId"`
			SheetID   uint32 `json:"sheetId" param:"sheetId" query:"sheetId" form:"sheetId"`
		}

		var value editCellParams
		c.Bind(&value)
		fmt.Println(value)
		return c.JSON(http.StatusOK, map[string]string{"me": "you"})
	}
}
