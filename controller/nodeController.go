package controller

import (
	"NGB-SE/middleware"
	"NGB-SE/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

//
//新建分区 POST方法，要求有token,json携带ZoneName
//src = /nodes/:Uid/create/zone
//
func CreateZone(c echo.Context) error {
	inData := new(CreateNodeInData)
	c.Bind(inData)
	tokenRaw := c.Request().Header.Get("Authorization")
	err := middleware.VerifyAdmin(c.Param("AdminId"), tokenRaw)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, outData)
	}
	var node model.Nodes
	node.AuthorId = c.Param("Uid")
	node.NodeType = "zone"
	node.SelfId = model.GetRandomId()
	node.NodeName = inData.ZoneName
	err = model.CreateNode(&node)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}
	outData := map[string]interface{}{
		"message": "分区创建成功",
		"ZoneId:": node.SelfId,
	}
	return c.JSON(http.StatusOK, outData)
}

//
//新建分区 POST方法，要求有token,json携带Content,ZoneId
//需要管理员权限
//返回PassageId和ZoneId
//src = /nodes/:Uid/create/passage
//
func CreatePassage(c echo.Context) error {
	inData := new(CreateNodeInData)
	c.Bind(inData)
	tokenRaw := c.Request().Header.Get("Authorization")
	_, err := middleware.VerifyUser(c.Param("Uid"), tokenRaw, false)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, outData)
	}
	var node model.Nodes
	node.AuthorId = c.Param("Uid")
	node.NodeType = "passage"
	node.Content = inData.Content
	node.FatherNodeId = inData.ZoneId
	node.SelfId = model.GetRandomId()
	zoneName, err := model.GetNodeName(inData.ZoneId)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)

	}
	err = model.CreateNode(&node)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}
	outData := map[string]interface{}{
		"message":   "分区创建成功",
		"PassageId": node.SelfId,
		"ZoneName":  zoneName,
	}
	return c.JSON(http.StatusOK, outData)
}

//
//GET src = /nodes/get/zones
//返回所有的分区id和名称
//
func QueryAllZones(c echo.Context) error {
	res, err := model.GetAllZones()
	zones := []ZoneNameStruct{}
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}
	//得到的数据里面提取出ZoneId和ZoneName
	for _, nodes := range res {
		var tempZone ZoneNameStruct
		tempZone.ZoneId = nodes.SelfId
		tempZone.ZoneName = nodes.NodeName
		zones = append(zones, tempZone)
	}
	return c.JSON(http.StatusOK, zones)
}

//
//GET src = /nodes/get/passages/:ZoneId
//返回所有的分区id和名称
//
func QueryAllPassageByZoneId(c echo.Context) error {
	res, err := model.GetAllPassageByZones(c.Param("ZoneId"))
	passages := []PassageStruct{}
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}
	//得到的数据里面提取出PassageId和Content
	for _, nodes := range res {
		var tempPassage PassageStruct
		tempPassage.PassageId = nodes.SelfId
		tempPassage.PassageContent = nodes.Content
		tempPassage.BelongZoneId = nodes.FatherNodeId
		tempPassage.BelongZoneName, _ = model.GetNodeName(nodes.FatherNodeId)
		passages = append(passages, tempPassage)
	}
	return c.JSON(http.StatusOK, passages)
}

//
//GET src = /view/passage/:PassageId/user/:Uid/like
//用户的点赞功能，需要token
//
func LikePassage(c echo.Context) error {
	tokenRaw := c.Request().Header.Get("Authorization")
	_, err := middleware.VerifyUser(c.Param("Uid"), tokenRaw, false)
	if err != nil {
		errMsg := MsgStruct{
			Message: err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, errMsg)
	}
	node, err := model.GetSingleNode(c.Param("PassageId"))
	if err != nil {
		msg := MsgStruct{
			Message: err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, msg)
	}
	if node.NodeType == "zone" {
		msg := MsgStruct{
			Message: "这是一个分区而不是文章",
		}
		return c.JSON(http.StatusBadRequest, msg) //这里不确定是不是BadRequest
	}
	node.Likes++
	model.SaveNode(&node)
	msg := MsgStruct{
		Message: fmt.Sprintf("点赞成功，现在文章%s有%d个赞", c.Param("PassageId"), node.Likes),
	}
	return c.JSON(http.StatusOK, msg)
}
