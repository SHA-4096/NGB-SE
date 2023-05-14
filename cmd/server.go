package main

import (
	"NGB-SE/internal/view"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "The server is running")
	})
	view.SetRouters(e)
	e.Logger.Fatal(e.Start(":1323"))

}
