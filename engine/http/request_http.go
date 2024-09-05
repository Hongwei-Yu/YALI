package http

import (
	"YALI/engine/constant"
	"github.com/valyala/fasthttp"
	"time"
)

func (r RequestHttp) Send() (bool, int64, uint64, float64, float64, string, time.Time, time.Time) {
	var (
		isSucceed     = true
		errCode       = constant.NoError
		receivedBytes = float64(0)
		errMsg        = ""
	)

	if r.HttpApiSetup == nil {
		r.HttpApiSetup = new(HttpApiSetup)
	}

	resp, req, requestTime, sendBytes, err, _, startTime, endTime := r.Request()

	defer fasthttp.ReleaseResponse(resp) // 用完需要释放资源
	defer fasthttp.ReleaseRequest(req)

	if err != nil {
		isSucceed = false
		errCode = constant.RequestError
		errMsg = err.Error()
	}

	receivedBytes = float64(resp.Header.ContentLength()) / 1024
	if receivedBytes <= 0 {
		receivedBytes = float64(len(resp.Body())) / 1024
	}
	// 缺少断言判断

	return isSucceed, errCode, requestTime, sendBytes, receivedBytes, errMsg, startTime, endTime

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
		err = client.DoTimeout(req, resp, 3*time.Second) // 这个超时时间应该修改一下
	}
	endTime = time.Now()
	requestTime = uint64(time.Since(startTime))
	return
}

func (r RequestHttp) RespAssert() {

}
func (r RequestHttp) Extract() {

}
