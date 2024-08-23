package main

import (
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/valyala/fasthttp"
)

func main() {
	for i := 0; i < 10000; i++ {
		go Post("http://www.baidu.com")
	}
	select {}
}
func Post(url string) {
	req := fasthttp.AcquireRequest()   //获取Request连接池中的连接
	defer fasthttp.ReleaseRequest(req) // 用完需要释放资源
	// 默认是application/x-www-form-urlencoded
	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")
	req.SetRequestURI(url)
	// requestBody := []byte(`{"request":"test"}`)
	// req.SetBody(requestBody)
	resp := fasthttp.AcquireResponse()             //获取Response连接池中的连接
	defer fasthttp.ReleaseResponse(resp)           // 用完需要释放资源
	if err := fasthttp.Do(req, resp); err != nil { //发送请求
		log.Error(err)
		return
	}
	b := resp.Body()
	// resp.Body()
	log.Info(string(b))
}
