package main

import (
	"NGB-SE/internal/middleware"
	"NGB-SE/internal/util"
	"NGB-SE/internal/view"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	view.SetRouters(e)
	util.EmailClientInit()
	middleware.RabbitMQInit()
	e.Logger.Fatal(e.Start(":8080"))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "The server is running")
	})

}
