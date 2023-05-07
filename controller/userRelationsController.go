package controller

import (
	"NGB-SE/middleware"
	"NGB-SE/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

//
//GET src = /relation/:Uid/mkfirend/:FriendId
//传入token
//
func AddFriend(c echo.Context) error {
	//用户认证部分
	tokenRaw := c.Request().Header.Get("Authorization")
	_, err := middleware.VerifyUser(c.Param("Uid"), tokenRaw, false)
	if err != nil {
		errMsg := MsgStruct{
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, errMsg)
	}
	//发送加好友请求
	err = model.CreateFirendRequest(c.Param("Uid"), c.Param("FriendId"))
	if err != nil {
		msg := MsgStruct{
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, msg)
	}
	message := fmt.Sprintf("用户%s希望加你为好友", c.Param("Uid"))
	err = model.CreateMessage(c.Param("Uid"), c.Param("FriendId"), message)
	if err != nil {
		return err
	}
	msg := MsgStruct{
		Message: "加好友请求已经发送",
	}
	return c.JSON(http.StatusUnauthorized, msg)

}

//
//GET src = /relation/:Uid/agree/:FriendId
//同意好友请求，传入token
//
func AgreeFriendRequest(c echo.Context) error {
	//用户认证部分
	tokenRaw := c.Request().Header.Get("Authorization")
	_, err := middleware.VerifyUser(c.Param("Uid"), tokenRaw, false)
	if err != nil {
		errMsg := MsgStruct{
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, errMsg)
	}
	//确认对方有给自己发请求
	err = model.CheckFriendRequest(c.Param("FriendId"), c.Param("Uid"))
	if err != nil {
		msg := MsgStruct{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}
	//确认两人此前不是好友
	state, err := model.GetUserRelation(c.Param("Uid"), c.Param("FriendId"))
	if state&model.StateFriend != 0 {
		//此前已经是好友了
		msg := MsgStruct{
			Message: fmt.Sprintf("你和%s早就是好友了哦", c.Param("FriendId")),
		}
		return c.JSON(http.StatusOK, msg)
	}
	//此前不是好友
	err = model.CreateFriendRelation(c.Param("Uid"), c.Param("FriendId"))
	if err != nil {
		return err
	}

	msg := MsgStruct{
		Message: fmt.Sprintf("你和%s已经成为好友了", c.Param("FriendId")),
	}
	return c.JSON(http.StatusOK, msg)
}
