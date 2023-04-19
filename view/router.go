package view

import (
	"NGB-SE/controller"

	"github.com/labstack/echo/v4"
)

func SetRouters(e *echo.Echo) {
	e.POST("/user/register", controller.Register)
	e.POST("/user/login", controller.Login)
	e.POST("/user/:Uid/logout", controller.LogOut)
	e.GET("/user/:Uid/delete", controller.DeleteUser)
	e.POST("/user/:Uid/modify", controller.ModifyUser)
	e.GET("/user/:Uid/refreshtoken", controller.RenewWithRefreshToken)
	e.GET("/user/admin/:AdminId/delete/:Uid", controller.AdminDeleteUser)
	e.POST("/user/admin/:AdminId/modify/:Uid", controller.AdminModifyUser)
}
