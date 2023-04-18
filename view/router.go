package view

import (
	"NGB-SE/controller"

	"github.com/labstack/echo/v4"
)

func SetRouters(e *echo.Echo) {
	e.POST("/user/register", controller.Register)
	e.POST("/user/login", controller.Login)
	e.GET("/user/delete/:Uid", controller.DeleteUser)
	e.GET("/user/admin/delete/:Uid", controller.AdminDeleteUser)
}
