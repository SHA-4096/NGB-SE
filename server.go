package main

import (
	"NGB-SE/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	db := model.ConnectDB()
	user := model.User{
		Name:     "test",
		Uid:      "112233345",
		Password: "233344",
	}
	db.Create(&user)
	fmt.Println("AAA")
	e.Logger.Fatal(e.Start(":1323"))
}
