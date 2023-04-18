package main

import (
	"NGB-SE/model"
	"NGB-SE/view"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "The server is running")
	})
	model.ConnectDB() //use model.DB to operate
	view.SetRouters(e)
	e.Logger.Fatal(e.Start(":1323"))

}
