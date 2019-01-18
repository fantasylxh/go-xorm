package utils

import (
	"gameCityService/global"
)

func AjaxResponse(code int, data interface{}, err string) map[string]interface{} {
	json := make(map[string]interface{})
	json["code"] = code
	json["data"] = data
	json["msg"] = err
	return json
}

func AjaxSuccess(data interface{}) map[string]interface{} {
	return AjaxResponse(global.HTTP_OK, data, "")
}

func AjaxError(err string) map[string]interface{} {
	return AjaxResponse(global.HTTP_ERR, "", err)
}
