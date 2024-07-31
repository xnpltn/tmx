package handlers

import (
	"spreadsheet/spdb/core"
	"spreadsheet/spdb/ui"
	"spreadsheet/spdb/ui/views"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func RenderView(c echo.Context, view templ.Component, layoutPath string) error {
	if c.Request().Header.Get("Hx-Request") == "true" {
		return view.Render(c.Request().Context(), c.Response())
	}
	return ui.Layout("/").Render(c.Request().Context(), c.Response())
}

func HomePage(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		return RenderView(c, views.HomeView(), "/")
	}
}

func EditPage(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		return RenderView(c, views.EditView(), "/edit")
	}
}
