package routers

import (
	"gameCityService/controllers/controller_advert"
	"gameCityService/controllers/controller_tyt"
	"github.com/labstack/echo"
)

func ApiRegisterRouter(e *echo.Echo) {
	e.POST("/gg/banner",controller_advert.Banner)
	e.POST("/gg/zhanshi",controller_advert.Zhanshi)
	e.POST("/gg/advertClick",controller_advert.AdvertClick)
	e.POST("/gg/share",controller_tyt.Share)
}

