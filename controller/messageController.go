package controller

import (
	"NGB-SE/middleware"
	"NGB-SE/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

//
//GET src = /message/:Uid
//获取收到的消息，传入token
//
func QueryMessage(c echo.Context) error {
	model.GetMessage(c.Param("Uid"))
	//用户认证部分
	tokenRaw := c.Request().Header.Get("Authorization")
	_, err := middleware.VerifyUser(c.Param("Uid"), tokenRaw, false)
	if err != nil {
		errMsg := MsgStruct{
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, errMsg)
	}
	//获取信息
	inStationMessage, err := model.GetMessage(c.Param("Uid"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, inStationMessage)
}
