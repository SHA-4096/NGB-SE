package view

import (
	"NGB-SE/controller"

	"github.com/labstack/echo/v4"
)

func SetRouters(e *echo.Echo) {
	//用户管理部分
	e.POST("/user/register", controller.Register)
	e.POST("/user/login", controller.Login)
	e.POST("/user/:Uid/logout", controller.LogOut)
	e.GET("/user/:Uid/delete", controller.DeleteUser)
	e.POST("/user/:Uid/modify", controller.ModifyUser)
	e.GET("/user/:Uid/refreshtoken", controller.RenewWithRefreshToken)
	e.GET("/user/admin/:AdminId/delete/:Uid", controller.AdminDeleteUser)
	e.POST("/user/admin/:AdminId/modify/:Uid", controller.AdminModifyUser)
	//节点管理部分
	e.POST("/nodes/:AdminId/create/zone", controller.CreateZone)
	e.POST("/nodes/:Uid/create/passage", controller.CreatePassage)
	e.GET("/nodes/get/zones", controller.QueryAllZones)
	e.GET("/nodes/get/passages/:ZoneId", controller.QueryAllPassageByZoneId)
	e.GET("/nodes/:Uid/delete/passage/:PassageId", controller.DeletePassageCommonUser)
	//用户操作管理部分
	e.GET("/view/passage/:PassageId/user/:Uid/like", controller.LikePassage)
	e.GET("/view/:Uid/likes/:FriendId", controller.QueryAllLikes)
}
