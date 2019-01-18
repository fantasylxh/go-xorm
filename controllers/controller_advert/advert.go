package controller_advert

import (
	"gameCityService/models/advert"
	"gameCityService/utils"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func Banner(c echo.Context) error  {
	banner := new(advert.AdBanner).GetBanner()
	return c.JSON(http.StatusOK,utils.AjaxSuccess(banner));
}

func Zhanshi(c echo.Context) error  {
	ad_id := c.Request().PostFormValue("ad_id")
	device := c.Request().PostFormValue("device")
	ad_id_int ,err := strconv.Atoi(ad_id)
	if err != nil{
		return c.JSON(http.StatusOK,utils.AjaxError(err.Error()))
	}
	device_int,err := strconv.Atoi(device)
	if err != nil{
		return c.JSON(http.StatusOK,utils.AjaxError(err.Error()))
	}
	id := new(advert.AdZhanshi).CreateData(ad_id_int,device_int)
	if len(id) <= 0{
		return c.JSON(http.StatusOK,utils.AjaxError(""))
	}
	return c.JSON(http.StatusOK,utils.AjaxSuccess(id))
}

func AdvertClick(c echo.Context) error  {
	zs_id := c.Request().PostFormValue("zs_id")
	new(advert.AdZhanshi).AdvertClick(zs_id)
	return c.JSON(http.StatusOK,utils.AjaxSuccess(""))
}
