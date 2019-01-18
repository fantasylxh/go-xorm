package routers

import (
	"gameCityService/controllers/controller_user"
	"github.com/labstack/echo"
)

func AdminRegisterRouter(e *echo.Group) {
	//domain := global.Cfg.Section("domain")
	/*e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://" + domain.Key("DomainApi").MustString("localhost")},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))*/

	e.GET("/user/info", controller_user.Info)

}
