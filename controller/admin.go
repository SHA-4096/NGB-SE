package controller

import (
	"NGB-SE/model"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

type AdminModifyUserINData struct {
	Key   string
	Value string
}

func verifyAdmin(c echo.Context) error {
	AdminId := c.Param("AdminId")
	user := new(model.User)
	tokenRaw := c.Request().Header.Get("Authorization")
	token := (strings.Split(tokenRaw, " "))
	if len(token) < 2 {
		return fmt.Errorf("你没有在请求头携带token")
	}
	fmt.Println("TOKEN IS:", token[1])
	//检查管理员身份
	model.DB.Where("Uid = ?", AdminId).First(&user)
	key := user.JwtKey
	claims, err := DecodeJwt(token[1], key)
	if err != nil {
		return err
	}
	if claims["Uid"] != AdminId {
		return fmt.Errorf("你的token似乎和你的身份不符")
	}
	if user.IsAdmin != "True" {
		return fmt.Errorf("你不是管理员")
	}
	return nil
}

func AdminDeleteUser(c echo.Context) error {
	/*POST src = /user/admin/{adminID}/delete/{userID}*/
	err := verifyAdmin(c)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, outData)
	}
	//检查要删除的用户是否存在
	user := new(model.User)
	err = model.DB.Where("Uid = ?", c.Param("Uid")).First(&user).Error
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}
	err = model.DB.Where("Uid = ?", c.Param("Uid")).Delete(&model.User{}).Error
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
	/*POST src = /user/admin/modify/:Uid with json containing key&value*/
	inData := new(AdminModifyUserINData)
	//验证身份
	err := verifyAdmin(c)
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
	user := new(model.User)
	err = model.DB.Where("Uid = ?", c.Param("Uid")).First(&user).Error
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
		model.DB.Save(&user)
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
