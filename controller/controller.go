package controller

import (
	"NGB-SE/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	/*POST Uid;Name;Password*/
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	token, key, err := GetJwt(user.Uid)
	if err != nil {
		panic(err)
	}
	data := map[string]interface{}{
		"token": token,
	}
	user.Jwt_key = key
	model.DB.Create(&user)
	return c.JSON(http.StatusCreated, data)
	//	return c.String(http.StatusOK, "JWT:"+token)
}
