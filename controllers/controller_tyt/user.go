package controller_tyt

import (
	"gameCityService/models/tyt"
	"gameCityService/utils"
	"github.com/labstack/echo"
	"net/http"
)

func LoginByUserName(c echo.Context) error  {
	username := c.Request().PostFormValue("username")
	password := c.Request().PostFormValue("password")
	user := new(tyt.User).Login(username,password)
	if user == nil{

	}
	return c.JSON(http.StatusOK,utils.AjaxSuccess(user))
}
