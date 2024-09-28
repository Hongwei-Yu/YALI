package service

import (
	"YALI/engine/http"
)

func DisposeApi(api *http.ApiHttp) (result bool, resp string) {

	// 发送请求
	result, _, _, _, _, resp, _, _ = api.Send()
	return result, resp
}
