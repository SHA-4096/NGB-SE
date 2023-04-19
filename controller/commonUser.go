package controller

import (
	"NGB-SE/model"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	/*POST Uid;Name;Password*/
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	user.IsAdmin = "False"
	data := map[string]interface{}{
		"message": "注册成功",
	}
	model.DB.Create(&user)
	return c.JSON(http.StatusCreated, data)
	//	return c.String(http.StatusOK, "JWT:"+token)
}

func Login(c echo.Context) error {
	/*POST Uid;Password*/
	inData := new(model.User)
	c.Bind(inData)
	var user model.User
	model.DB.Where("Uid = ? AND Password= ?", inData.Uid, inData.Password).First(&user)
	if user.Uid == "" {
		outData := map[string]interface{}{
			"message": "帐号不存在或密码错误",
		}
		return c.JSON(http.StatusUnauthorized, outData)
	}
	token, key, err := GetJwt(user.Uid)
	if err != nil {
		panic(err)
	}
	user.JwtKey = key
	model.DB.Save(&user)
	outData := map[string]interface{}{
		"token":   token,
		"message": fmt.Sprintf("welcome:%s", user.Name),
	}
	return c.JSON(http.StatusOK, outData)

}

func verifyUser(c echo.Context) error {
	Uid := c.Param("Uid")
	user := new(model.User)
	model.DB.Where("Uid = ?", Uid).First(&user)
	if user.Uid == "" {
		return fmt.Errorf("您不是已注册用户")
	}
	tokenRaw := c.Request().Header.Get("Authorization")
	token := (strings.Split(tokenRaw, " "))
	if len(token) < 2 {
		return fmt.Errorf("你没有在请求头携带token")
	}
	fmt.Println("TOKEN IS:", token[1])
	key := user.JwtKey
	claims, err := DecodeJwt(token[1], key)
	if err != nil {
		return err
	}
	if claims["Uid"] != Uid {
		return fmt.Errorf("你的token无效")
	}
	return nil
}

func DeleteUser(c echo.Context) error {
	/*GET src = /user/delete/:Uid*/
	err := verifyUser(c)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}
	err = model.DB.Where("Uid = ?", c.Param("Uid")).Delete(&model.User{}).Error
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	} else {
		outData := map[string]interface{}{
			"message": fmt.Sprintf("用户%s已经注销", c.Param("Uid")),
		}
		return c.JSON(http.StatusOK, outData)
	}

}

func ModifyUser(c echo.Context) error {
	/*POST src = /user/modify/:Uid with json containing key&value*/
	inData := new(AdminModifyUserINData)
	err := verifyUser(c)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
	}
	err = c.Bind(inData)
	if err != nil {
		return c.String(http.StatusUnauthorized, err.Error())
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
