package controller

import (
	"NGB-SE/middleware"
	"NGB-SE/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreateNodeInData struct {
	ZoneName string
	Content  string
}

//
//新建分区 POST方法，要求有token,json携带ZoneName
//src = /nodes/:Uid/Create/Zone
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
	model.CreateNode(&node)

}

//
//新建分区 POST方法，要求有token,json携带Content,ZoneName
//src = /nodes/:Uid/Create/passage
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
}
