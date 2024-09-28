package http

import (
	"YALI/constant"
	"YALI/engine/model"
	"YALI/log"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/valyala/fasthttp"
	"strings"
	"sync"
	"time"
)

func (r RequestHttp) Send() (bool, int64, uint64, float64, float64, string, time.Time, time.Time) {
	// 定义变量
	var (
		// 表示请求是否成功
		isSucceed = true
		// 表示错误代码
		errCode = constant.NoError
		// 表示接收的字节数
		receivedBytes = float64(0)
		// 表示错误信息
		errMsg = ""
	)
	// 如果 HttpApiSetup 为 nil，则创建一个新的 HttpApiSetup 实例
	if r.HttpApiSetup == nil {
		r.HttpApiSetup = new(HttpApiSetup)
	}
	// 调用 Request 方法，获取响应、请求、请求时间、发送字节数、错误信息等
	resp, req, requestTime, sendBytes, err, _, startTime, endTime := r.Request()

	// 延迟调用 ReleaseResponse 方法，释放响应资源
	defer fasthttp.ReleaseResponse(resp)
	// 延迟调用 ReleaseRequest 方法，释放请求资源
	defer fasthttp.ReleaseRequest(req)

	// 如果发生错误
	if err != nil {
		// 设置请求失败
		isSucceed = false
		// 设置错误代码为请求错误
		errCode = constant.RequestError
		// 设置错误信息为错误的具体内容
		errMsg = err.Error()
	}

	// 计算接收的字节数
	receivedBytes = float64(resp.Header.ContentLength()) / 1024
	// 如果接收的字节数小于等于 0，则使用响应体的长度来计算
	if receivedBytes <= 0 {
		receivedBytes = float64(len(resp.Body())) / 1024
	}
	/**
	 * documentation comment：
	 * @Description:  This function sends an HTTP request and returns whether the request was successful.
	 * @return isSucceed Indicates whether the request was successful.
	 * @return errCode Error code, in this function indicates request error.
	 * @return receivedBytes Indicates the number of bytes received.
	 * @return errMsg Error message.
	 */
	// 返回请求是否成功、错误代码、接收字节数、错误信息等

	// 缺少断言
	var regex = &constant.Regex{}

	r.Withdraw(regex, &sync.Map{}, resp)

	return isSucceed, errCode, requestTime, sendBytes, receivedBytes, errMsg, startTime, endTime

}

func (r RequestHttp) Request() (resp *fasthttp.Response, req *fasthttp.Request, requestTime uint64, sendBytes float64, err error, str string, startTime, endTime time.Time) {
	var client *fasthttp.Client
	req = fasthttp.AcquireRequest()
	resp = fasthttp.AcquireResponse()
	client = FastClient(r.HttpApiSetup, r.Auth)

	r.Header.SetHeader(req)
	r.Cookie.SetCookie(req)

	urls := strings.Split(r.URL, "//")
	if !strings.EqualFold(urls[0], constant.HTTP) && !strings.EqualFold(urls[0], constant.HTTPS) {
		r.URL = constant.HTTP + "//" + r.URL
	}
	urlQuery := req.URI().QueryArgs()
	if r.Query.Parameter != nil {
		for _, v := range r.Query.Parameter {
			if v.IsChecked != constant.Open {
				continue
			}
			if !strings.Contains(r.URL, v.Key) {
				by := v.ValueToByte()
				urlQuery.AddBytesV(v.Key, by)
				r.URL = r.URL + fmt.Sprintf("&%s=%s", v.Key, string(v.ValueToByte()))
			}
		}
	}
	req.SetRequestURI(r.URL)
	// set body
	str = r.Body.SetBody(req)

	startTime = time.Now()
	if r.HttpApiSetup.IsRedirects == 0 {
		req.SetTimeout(30 * time.Second)
		err = client.DoRedirects(req, resp, r.HttpApiSetup.RedirectsNum)

	} else {
		err = client.DoTimeout(req, resp, 3*time.Second) // 这个超时时间应该修改一下
	}
	endTime = time.Now()
	requestTime = uint64(time.Since(startTime))
	if r.Debug == "all" {
		log.Logger.Info(fmt.Sprintf("请求时间：%d", requestTime))
		log.Logger.Info(fmt.Sprintf("响应体：%s", string(resp.Body())))
	}
	return resp, req, requestTime, float64(req.Header.ContentLength()), err, str, startTime, endTime
}

func (r RequestHttp) RespAssert() {

}
func (r RequestHttp) Withdraw(regex *constant.Regex, globalValue *sync.Map, resp *fasthttp.Response) {
	if r.Regex == nil {
		return
	}
	for _, regular := range r.Regex {
		if regular.IsChecked != constant.Open {
			continue
		}
		reg := new(constant.Reg)
		value := regular.Extract(resp, globalValue)
		if value == nil {
			continue
		}
		reg.Key = regular.Var
		reg.Value = value
		regex.Regs = append(regex.Regs, reg)
	}
}

// FastClient 获取fasthttp客户端
func FastClient(httpApiSetup *HttpApiSetup, auth *model.Auth) (fc *fasthttp.Client) {
	tr := &tls.Config{InsecureSkipVerify: true}
	if auth != nil || auth.Bidirectional != nil {
		switch auth.Type {
		case Bidirectional:
			tr.InsecureSkipVerify = false
			if auth.Bidirectional.CaCert != "" {
				if strings.HasPrefix(auth.Bidirectional.CaCert, "https://") || strings.HasPrefix(auth.Bidirectional.CaCert, "http://") {
					client := &fasthttp.Client{}
					loadReq := fasthttp.AcquireRequest()
					defer loadReq.ConnectionClose()
					// set url
					loadReq.Header.SetMethod("GET")
					loadReq.SetRequestURI(auth.Bidirectional.CaCert)
					loadResp := fasthttp.AcquireResponse()
					defer loadResp.ConnectionClose()
					if err := client.Do(loadReq, loadResp); err != nil {
						//log.Error(fmt.Sprintf("机器ip:%s, 下载crt文件失败：", global.GlobalHost), err)
					}
					if loadResp != nil && loadResp.Body() != nil {
						caCertPool := x509.NewCertPool()
						if caCertPool != nil {
							caCertPool.AppendCertsFromPEM(loadResp.Body())
							tr.ClientCAs = caCertPool
						}
					}
				}
			}
		case Unidirectional:
			tr.InsecureSkipVerify = false
		}
	}
	fc = &fasthttp.Client{
		TLSConfig: tr,
	}
	if httpApiSetup.ClientName != "" {
		fc.Name = httpApiSetup.ClientName
	}
	if httpApiSetup.UserAgent {
		fc.NoDefaultUserAgentHeader = false
	}
	if httpApiSetup.MaxIdleConnDuration != 0 {
		fc.MaxIdleConnDuration = time.Duration(httpApiSetup.MaxIdleConnDuration) * time.Second
	} else {
		fc.MaxIdleConnDuration = time.Duration(0) * time.Second
	}
	if httpApiSetup.MaxConnPerHost != 0 {
		fc.MaxConnsPerHost = httpApiSetup.MaxConnPerHost
	}

	fc.MaxConnWaitTimeout = time.Duration(httpApiSetup.MaxConnWaitTimeout) * time.Second
	fc.WriteTimeout = time.Duration(httpApiSetup.WriteTimeOut) * time.Millisecond
	fc.ReadTimeout = time.Duration(httpApiSetup.ReadTimeOut) * time.Millisecond

	return fc
}
