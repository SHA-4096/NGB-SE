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
	return c.JSON(http.StatusOK, msg)

}

//
//GET src = /relation/:Uid/mkfollow/:FollowId
//关注某个用户，需要token
//
func AddFollow(c echo.Context) error {
	//用户认证部分
	tokenRaw := c.Request().Header.Get("Authorization")
	_, err := middleware.VerifyUser(c.Param("Uid"), tokenRaw, false)
	if err != nil {
		errMsg := MsgStruct{
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, errMsg)
	}
	//检查要关注的用户是否存在
	_, err = model.QueryUid(c.Param("FollowId"))
	if err != nil {
		errMsg := MsgStruct{
			Message: "你要关注的用户不存在",
		}
		return c.JSON(http.StatusOK, errMsg)
	}
	//检查用户关系
	relation, err := model.GetUserRelation(c.Param("Uid"), c.Param("FollowId"))
	if err != nil {
		return err
	}
	if relation&model.StateFollow != 0 {
		msg := MsgStruct{
			Message: "已经关注该用户",
		}
		return c.JSON(http.StatusOK, msg)
	}
	//添加关注关系
	err = model.CreateFollowRelation(c.Param("Uid"), c.Param("FollowId"))
	if err != nil {
		return err
	}
	model.CreateMessage(c.Param("Uid"), c.Param("FollowId"), fmt.Sprintf("%s关注了你", c.Param("Uid")))
	msg := MsgStruct{
		Message: "添加关注成功",
	}
	return c.JSON(http.StatusOK, msg)
}

//
//GET src = /relation/:Uid/unfollow/:FollowId
//取消关注
//需要token
//
func UnFollow(c echo.Context) error {
	//用户认证部分
	tokenRaw := c.Request().Header.Get("Authorization")
	_, err := middleware.VerifyUser(c.Param("Uid"), tokenRaw, false)
	if err != nil {
		errMsg := MsgStruct{
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, errMsg)
	}
	//检查要关注的用户是否存在
	_, err = model.QueryUid(c.Param("FollowId"))
	if err != nil {
		errMsg := MsgStruct{
			Message: "你要取消关注的用户不存在",
		}
		return c.JSON(http.StatusOK, errMsg)
	}
	//检查用户关系
	relation, err := model.GetUserRelation(c.Param("Uid"), c.Param("FollowId"))
	if err != nil {
		errMsg := MsgStruct{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, errMsg)
	}
	if relation&model.StateFollow == 0 {
		msg := MsgStruct{
			Message: "还没有关注该用户",
		}
		return c.JSON(http.StatusOK, msg)
	}
	//删除关注关系
	err = model.DeleteFollowRelation(c.Param("Uid"), c.Param("FollowId"))
	if err != nil {
		return err
	}
	model.CreateMessage(c.Param("Uid"), c.Param("FollowId"), fmt.Sprintf("%s取消了对你的关注", c.Param("Uid")))
	msg := MsgStruct{
		Message: "取消关注成功",
	}
	return c.JSON(http.StatusOK, msg)
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
	state, _ := model.GetUserRelation(c.Param("Uid"), c.Param("FriendId"))
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

//
//src = /relation/query/follows/:Uid
//查询用户的关注列表
//
func QueryAllFollows(c echo.Context) error {
	relations, _ := model.GetAllFollows(c.Param("Uid"))
	return c.JSON(http.StatusOK, relations)
}
