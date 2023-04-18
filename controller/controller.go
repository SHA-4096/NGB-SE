package controller

import (
	"NGB-SE/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	/*POST Uid;Name;Password*/
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

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
	model.DB.Where("Uid = ?", inData.Uid).First(&user)
	if user.Uid == "" {
		return c.String(http.StatusUnauthorized, "You're not a registered user")
	}
	token, key, err := GetJwt(user.Uid)
	if err != nil {
		panic(err)
	}
	user.Jwt_key = key
	model.DB.Save(&user)
	outData := map[string]interface{}{
		"token":   token,
		"message": fmt.Sprintf("welcome:%s", user.Name),
	}
	return c.JSON(http.StatusOK, outData)

}

func VerifyUser(c echo.Context) error {
	/*GET with token,src = /user/verify/:Uid */
	Uid := c.Param("Uid")
	user := new(model.User)
	model.DB.Where("Uid = ?", Uid).First(&user)
	if user.Uid == "" {
		return c.String(http.StatusUnauthorized, "您不是已注册用户")
	}
	tokenRaw := c.Request().Header.Get("Authorization")
	token := (strings.Split(tokenRaw, " "))[1]
	fmt.Println("TOKEN IS:", token)
	key := user.Jwt_key
	claims, err := DecodeJwt(token, key)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	if claims["Uid"] != Uid {
		return c.String(http.StatusForbidden, "你的token无效")
	}
	return c.String(http.StatusOK, "用户认证成功")

}
