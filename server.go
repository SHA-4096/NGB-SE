package main

import (
	"NGB-SE/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	model.ConnectDB()
	e.Logger.Fatal(e.Start(":1323"))
}
