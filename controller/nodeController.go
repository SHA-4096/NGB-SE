package controller

import (
	"NGB-SE/middleware"
	"NGB-SE/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreateNodeInData struct {
	ZoneName string
	ZoneId   string
	Content  string
}

type ZoneNameStruct struct {
	ZoneId   string
	ZoneName string
}

//
//新建分区 POST方法，要求有token,json携带ZoneName
//src = /nodes/:Uid/create/zone
//
func CreateZone(c echo.Context) error {
	inData := new(CreateNodeInData)
	c.Bind(inData)
	_, err := middleware.VerifyUser(c, false)
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
//返回PassageId和ZoneId
//src = /nodes/:Uid/create/passage
//
func CreatePassage(c echo.Context) error {
	inData := new(CreateNodeInData)
	c.Bind(inData)
	_, err := middleware.VerifyUser(c, false)
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
