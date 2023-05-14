package controller

import (
	"NGB-SE/internal/middleware"
	"NGB-SE/internal/model"
	"fmt"
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
)

type AdminModifyUserINData struct {
	Key   string
	Value string
}

func AdminDeleteUser(c echo.Context) error {
	/*POST src = /user/admin/{adminID}/delete/{userID}*/
	tokenRaw := c.Request().Header.Get("Authorization")
	err := middleware.VerifyAdmin(c.Param("AdminId"), tokenRaw)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, outData)
	}
	//检查要删除的用户是否存在
	err = model.DeleteUid(c.Param("Uid"))
	if err != nil {
		outData := map[string]interface{}{
			"message": "找不到用户",
		}
		return c.JSON(http.StatusInternalServerError, outData)
	} else {
		outData := map[string]interface{}{
			"message": fmt.Sprintf("用户%s已经删除", c.Param("Uid")),
		}
		return c.JSON(http.StatusOK, outData)
	}

}

func AdminModifyUser(c echo.Context) error {
	/*POST src = /user/:AdminId/modify/:Uid with json containing key&value*/
	inData := new(AdminModifyUserINData)
	//验证身份
	tokenRaw := c.Request().Header.Get("Authorization")
	err := middleware.VerifyAdmin(c.Param("AdminId"), tokenRaw)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, outData)
	}
	//获取数据
	err = c.Bind(inData)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}
	user, err := model.QueryUid(c.Param("Uid"))
	//查询出错时
	if user.Uid == "" {
		outData := map[string]interface{}{
			"message": fmt.Sprintf("数据库出了问题耶，错误信息：%s", err),
		}
		return c.JSON(http.StatusBadRequest, outData)
	}
	//找到用户时
	fmt.Print(inData)
	refUser := reflect.ValueOf(user).Elem()
	fieldValue := refUser.FieldByName(inData.Key)
	if fieldValue.IsValid() {
		fieldValue.SetString(inData.Value)
		model.SaveUser(user)
		outData := map[string]interface{}{
			"message": fmt.Sprintf("用户%s的%s值被修改为%s", c.Param("Uid"), inData.Key, inData.Value),
		}
		return c.JSON(http.StatusOK, outData)
	} else {
		outData := map[string]interface{}{
			"message": "要修改的键值不存在",
		}
		return c.JSON(http.StatusBadRequest, outData)
	}

}
