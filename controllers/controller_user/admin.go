package controller_user

import (
	"gameCityService/utils"
	"github.com/labstack/echo"
	"net/http"
)

func Info(c echo.Context) error  {
	data := make(map[string]interface{})
	roles := make([]string,0)
	roles = append(roles,"editor","develop")
	data["roles"] = roles
	return c.JSON(http.StatusOK,utils.AjaxSuccess(data))
	//new(advert.AdAdmin).GetAdminByUserName()
}
