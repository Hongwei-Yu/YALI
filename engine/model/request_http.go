package model

import (
	"github.com/valyala/fasthttp"
	"time"
)

type RequestHttp struct {
	PreUrl       string               `json:"pre_url"`
	URL          string               `json:"url"`
	Method       string               `json:"method"` // 方法 GET/POST/PUT
	Debug        string               `json:"debug"`
	Parameter    []*VarForm           `json:"parameter"`
	Header       *Header              `json:"header"` // Headers
	Query        *Query               `json:"query"`
	Body         *Body                `json:"body"`
	Auth         *Auth                `json:"auth"`
	Cookie       *Cookie              `json:"cookie"`
	HttpApiSetup *HttpApiSetup        `json:"http_api_setup"`
	Assert       []*AssertionText     `json:"assert"` // 验证的方法(断言)
	Regex        []*RegularExpression `json:"regex"`  // 正则表达式
}

func (r RequestHttp) Send() {

}

func (r RequestHttp) Request() (resp *fasthttp.Response, req *fasthttp.Request, requestTime uint64, sendBytes float64, err error, str string, startTime, endTime time.Time) {
	var client *fasthttp.Client
	req = fasthttp.AcquireRequest()
	resp = fasthttp.AcquireResponse()
	client = fastClient(r.HttpApiSetup, r.Auth)
	startTime = time.Now()
	if r.HttpApiSetup.IsRedirects == 0 {
		req.SetTimeout(30 * time.Second)
		err = client.DoRedirects(req, resp, r.HttpApiSetup.RedirectsNum)
	} else {
		err = client.DoTimeout(req, resp, 3*time.Second)
	}
	endTime = time.Now()
	requestTime = uint64(time.Since(startTime))
	return
}
