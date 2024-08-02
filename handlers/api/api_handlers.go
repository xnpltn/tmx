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

// saves cell data as the cell changes
func Save(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
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
				return c.JSON(http.StatusBadRequest, map[string]string{"errir": "something went wrong"})
			}
		}

		if res := app.DB().Create(&sheet); res.Error != nil {
			fmt.Println("error saving: ", res.Error)
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "something went wrong"})
		}

		return c.JSON(http.StatusOK, map[string]string{"errir": "something went wrong"})
	}
}

func DeleteSheet(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
