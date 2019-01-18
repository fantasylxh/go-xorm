package utils

import (
	"github.com/json-iterator/go"
	"io"
	"io/ioutil"
)

func AnalysisRequestParams(r io.ReadCloser) (params map[string] interface{})  {
	body, err := ioutil.ReadAll(r)
	if err == nil{
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		json.Unmarshal(body,&params)
	}
	return params
}
