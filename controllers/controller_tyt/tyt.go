package controller_tyt

import (
	"crypto/md5"
	"fmt"
	"gameCityService/models/tyt"
	"gameCityService/utils"
	"github.com/labstack/echo"
	"net/http"
)

func Share(c echo.Context) error {
	const APPKEY string = "xaE5zGZOlwTVGmAMHAmbxt0lPOBcs9SGSOwsFKIF"
	uid := c.Request().PostFormValue("uid")               // 用户id
	share_code := c.Request().PostFormValue("share_code") // 邀请码
	sign := c.Request().PostFormValue("sign")             // 签名字符串
	// 签名验证
	signstr := uid + share_code + APPKEY // 签名字符串
	//_ = signstr
	signdata := []byte(signstr)
	has := md5.Sum(signdata)
	signstr = fmt.Sprintf("%x", has) //  MD5 加密签名字符串

	if uid == "" {
		return c.JSON(http.StatusOK, utils.AjaxError("uid 不能为空"))
	} else if share_code == "" {
		return c.JSON(http.StatusOK, utils.AjaxError("邀请码不能为空"))
	} else if sign != signstr {
		return c.JSON(http.StatusOK, utils.AjaxError("无效的签名"+signstr)) // todo remove signstr
	}

	new(tyt.User).UpdateVipEndTime(uid, share_code)
	return c.JSON(http.StatusOK, utils.AjaxSuccess("success"))
}
