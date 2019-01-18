package controllers

import (
	"bytes"
	"encoding/base64"
	"gameCityService/global"
	"gameCityService/models/advert"
	"gameCityService/utils"
	"github.com/json-iterator/go"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func PublicsLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", "")
}

func PublicCheckCode(c echo.Context) error {
	img, key := utils.CreateCode()
	cookie := new(http.Cookie)
	cookie.Value = key
	cookie.Name = "idkey"
	c.SetCookie(cookie)
	buff, _ := base64.StdEncoding.DecodeString(img)
	return c.Stream(http.StatusOK, "image/png", bytes.NewBuffer(buff))
}

func PublicsRunLogin(c echo.Context) error {
	params := utils.AnalysisRequestParams(c.Request().Body)
	code := params["code"]
	key, _ := c.Cookie("idkey")
	if key != nil && utils.CheckCode(code.(string), key.Value) {
		//fmt.Println("66666")
	} else {
		//fmt.Println("验证码错误")
		//return c.JSON(http.StatusOK,utils.AjaxError("验证码错误"))
	}

	//host := c.Request().Header.Get("Host")
	//remoteAddr := strings.Split(c.Request().RemoteAddr,":")
	username := params["username"]
	password := params["password"]
	userInfo,err := new(advert.AdAdmin).GetAdminByUserName(username.(string))
	if err != nil{
		return c.JSON(http.StatusOK,utils.AjaxError(err.Error()))
	}
	if userInfo == nil{
		return c.JSON(http.StatusOK,utils.AjaxError("用户名不存在"))
	}
	if utils.CheckPasswordHash(password.(string),userInfo.Password){
		var json  = jsoniter.ConfigCompatibleWithStandardLibrary
		jsonByte,_ := json.Marshal(userInfo)
		token := utils.EncryptPassword(userInfo.Username + "&"+userInfo.Password,strconv.Itoa(userInfo.Id))
		result := make(map[string]string)
		result["token"] = token
		result["username"] = userInfo.Username
		result["introduction"] = userInfo.Introduction
		result["avatar"] = userInfo.Avatar;
		global.RedisConn.SetNX("Auth_"+token,string(jsonByte),60*60*time.Millisecond).Err()

		return c.JSON(http.StatusOK,utils.AjaxSuccess(result))
	}
	return c.JSON(http.StatusOK,utils.AjaxError("密码错误"))
}

func PublicsRunLogout(c echo.Context) error {
	authorization :=c.Request().Header.Get("Authorization")
	strs := strings.Split(authorization," ")
	if len(strs) == 2{
		global.RedisConn.Del("Auth_"+ strs[1])
	}
	return c.JSON(http.StatusOK,utils.AjaxSuccess("OK"))
}
