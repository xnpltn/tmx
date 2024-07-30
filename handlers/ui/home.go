package handlers

import (
	"speadshets/spdb/core"
	"speadshets/spdb/ui"
	"speadshets/spdb/ui/views"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

/*
	func (a *Application) RenderView(c echo.Context, view templ.Component, layoutPath string) error {
		if c.Request().Header.Get("Hx-Request") == "true" {
			return view.Render(c.Request().Context(), c.Response().Writer)
		}
		_, err := utls.CheckCredentials(c, a.db)
		if err != nil {
			return templates.Layout(layoutPath, false).Render(c.Request().Context(), c.Response().Writer)
		}
		return templates.Layout(layoutPath, true).Render(c.Request().Context(), c.Response().Writer)
	}
*/
func RenderView(c echo.Context, view templ.Component, layoutPath string) error {
	if c.Request().Header.Get("Hx-Request") == "true" {
		return view.Render(c.Request().Context(), c.Response())
	}
	return ui.Layout("/").Render(c.Request().Context(), c.Response())
}

// func HomePage(app core.App) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		return ui.Layout().Render(c.Request().Context(), c.Response().Writer)
// 	}
// }

func HomePage(app core.App) echo.HandlerFunc {
	return func(c echo.Context) error {
		return RenderView(c, views.HomeView(), "/")
	}
}
