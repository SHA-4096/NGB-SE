package controller

import (
	"NGB-SE/model"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func verifyAdmin(c echo.Context) error {
	Uid := c.Param("Uid")
	user := new(model.User)
	model.DB.Where("Uid = ?", Uid).First(&user)
	if user.Uid == "" {
		return fmt.Errorf("您不是已注册用户")
	}
	if user.IsAdmin == "False" {
		return fmt.Errorf("您不是管理员")
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

func AdminDeleteUser(c echo.Context) error {
	/*GET src = /user/admin/delete/{userID}*/
	err := verifyAdmin(c)
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
