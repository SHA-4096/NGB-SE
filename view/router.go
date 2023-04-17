package view

import (
	"NGB-SE/controller"

	"github.com/labstack/echo/v4"
)

func SetRouters(e *echo.Echo) {
	e.POST("/register", controller.Register)
}
