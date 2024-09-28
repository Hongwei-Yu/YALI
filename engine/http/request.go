package http

import (
	"YALI/engine/model"
)

type ApiHttp struct {
	*model.Api
	*RequestHttp `json:"request"`
}

// RequestHttp 请求体 结构
type RequestHttp struct {
	PreUrl       string                     `json:"pre_url"` // 前置URL例如网址
	URL          string                     `json:"url"`     // 请求地址uri
	Method       string                     `json:"method"`  // 方法 GET/POST/PUT
	Debug        string                     `json:"debug"`
	Parameter    []*model.VarForm           `json:"parameter"`
	Header       *model.Header              `json:"header"` // Headers
	Query        *model.Query               `json:"query"`
	Body         *model.Body                `json:"body"`
	Auth         *model.Auth                `json:"auth"`
	Cookie       *model.Cookie              `json:"cookie"`
	HttpApiSetup *HttpApiSetup              `json:"http_api_setup"`
	Assert       []*model.AssertionText     `json:"assert"` // 验证的方法(断言)
	Regex        []*model.RegularExpression `json:"regex"`  // 正则表达式
}
type HttpApiSetup struct {
	ClientName          string `json:"client_name"`
	IsRedirects         int64  `json:"is_redirects"`           // 是否跟随重定向 0: 是   1：否
	RedirectsNum        int    `json:"redirects_num"`          // 重定向次数>= 1; 默认为3
	ReadTimeOut         int64  `json:"read_time_out"`          // 请求读取超时时间
	WriteTimeOut        int64  `json:"write_time_out"`         // 响应读取超时时间
	KeepAlive           bool   `json:"keep_alive"`             // 是否保持长连接
	MaxIdleConnDuration int64  `json:"max_idle_conn_duration"` // 最大空闲连接时长
	MaxConnPerHost      int    `json:"max_conn_per_host"`      // 最大连接数
	UserAgent           bool   `json:"user_agent"`             // 是否携带User-Agent
	MaxConnWaitTimeout  int64  `json:"max_conn_wait_timeout"`  // 最大等待时长
}
