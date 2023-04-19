package controller

import (
	"NGB-SE/model"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/labstack/echo/v4"
)

type LogOutInData struct {
	Uid string
}

func encodeMethod(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func Register(c echo.Context) error {
	/*POST Uid;Name;Password*/
	//检查id是否被使用
	user := new(model.User)
	checkUser := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	err := model.DB.Where("Uid = ?", user.Uid).First(&checkUser).Error
	if err == nil {
		outData := map[string]interface{}{
			"message": "这个用户id已经被占用了哦",
		}
		return c.JSON(http.StatusBadRequest, outData)

	}
	user.IsAdmin = "False"
	user.Password = encodeMethod(user.Password)
	model.DB.Create(&user)

	data := map[string]interface{}{
		"message": "注册成功",
	}
	return c.JSON(http.StatusCreated, data)
	//	return c.String(http.StatusOK, "JWT:"+token)
}

func Login(c echo.Context) error {
	/*POST Uid;Password*/
	inData := new(model.User)
	c.Bind(inData)
	var user model.User
	model.DB.Where("Uid = ? AND Password= ?", inData.Uid, encodeMethod(inData.Password)).First(&user)
	if user.Uid == "" {
		outData := map[string]interface{}{
			"message": "帐号不存在或密码错误",
		}
		return c.JSON(http.StatusUnauthorized, outData)
	}
	jwtToken, key, err := GetJwt(user.Uid)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}
	user.JwtKey = key
	model.DB.Save(&user)
	refreshToken, err := GetRefreshJwt(user.Uid)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}
	outData := map[string]interface{}{
		"jwtToken":     jwtToken,
		"refreshToken": refreshToken,
		"message":      fmt.Sprintf("welcome:%s", user.Name),
	}
	return c.JSON(http.StatusOK, outData)

}

func LogOut(c echo.Context) error {
	/*GET src = /user/:Uid/logout with token*/
	_, err := verifyUser(c, false)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, outData)
	}
	//从数据库里面取出用户
	user := new(model.User)
	err = model.DB.Where("Uid = ?", c.Param("Uid")).First(&user).Error
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, outData)
	}
	//注销用户
	user.JwtKey = ""
	model.DB.Save(&user)
	outData := map[string]interface{}{
		"message": fmt.Sprintf("用户%s已经注销", c.Param("Uid")),
	}
	return c.JSON(http.StatusInternalServerError, outData)
}

func verifyUser(c echo.Context, isRefresh bool) (string, error) {
	/*内部函数，用来验证用户,需要:Uid的路径参数以及token,会返回一个token和error*/
	Uid := c.Param("Uid")
	user := new(model.User)
	model.DB.Where("Uid = ?", Uid).First(&user)
	if user.Uid == "" {
		return "", fmt.Errorf("您不是已注册用户")
	}
	tokenRaw := c.Request().Header.Get("Authorization")
	token := (strings.Split(tokenRaw, " "))
	if len(token) < 2 {
		return "", fmt.Errorf("你没有在请求头携带token")
	}
	fmt.Println("TOKEN IS:", token[1])
	var key string
	if isRefresh {
		key = refreshTokenKey
	} else {
		key = user.JwtKey
	}

	claims, err := DecodeJwt(token[1], key)
	if err != nil {
		return "", err
	}
	if claims["Uid"] != Uid {
		return "", fmt.Errorf("你的token无效")
	}
	return token[1], nil
}

func DeleteUser(c echo.Context) error {
	/*GET src = /user/delete/:Uid*/
	_, err := verifyUser(c, false)
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
			"message": fmt.Sprintf("用户%s已经删除", c.Param("Uid")),
		}
		return c.JSON(http.StatusOK, outData)
	}

}

func ModifyUser(c echo.Context) error {
	/*POST src = /user/modify/:Uid with json containing key&value*/
	inData := new(AdminModifyUserINData)
	_, err := verifyUser(c, false)
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

func RenewWithRefreshToken(c echo.Context) error {
	/*GET方法,更新jwtToken,src = /user/:Uid/refreshtoken,请求头携带refreshToken*/
	//检查用户
	refreshToken, err := verifyUser(c, true)
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, outData)
	}

	user := new(model.User)
	err = model.DB.Where("Uid = ?", c.Param("Uid")).First(&user).Error
	if err != nil {
		outData := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusUnauthorized, outData)
	}
	//分发新的token
	jwtToken, key, err := GetJwt(user.Uid)
	if err != nil {
		panic(err)
	}
	user.JwtKey = key
	model.DB.Save(&user)
	outData := map[string]interface{}{
		"jwtToken":     jwtToken,
		"refreshToken": refreshToken,
		"message":      fmt.Sprintf("welcome:%s", user.Name),
	}
	return c.JSON(http.StatusOK, outData)

}
