package http

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/valyala/fasthttp"
	"strings"
	"time"
)

type VarForm struct {
	IsChecked   int64       `json:"is_checked" bson:"is_checked"`
	Type        string      `json:"type" bson:"type"`
	FileBase64  []string    `json:"fileBase64"`
	Key         string      `json:"key" bson:"key"`
	Value       interface{} `json:"value" bson:"value"`
	NotNull     int64       `json:"not_null" bson:"not_null"`
	Description string      `json:"description" bson:"description"`
	FieldType   string      `json:"field_type" bson:"field_type"`
}

type Header struct {
	Parameter []*VarForm `json:"parameter"`
}

type Query struct {
	Parameter []*VarForm `json:"parameter"`
}

type Body struct {
	Mode      string     `json:"mode"`
	Raw       string     `json:"raw"`
	Parameter []*VarForm `json:"parameter"`
}

type Auth struct {
	Type          string    `json:"type" bson:"type"`
	Bidirectional *TLS      `json:"bidirectional"`
	KV            *KV       `json:"kv" bson:"kv"`
	Bearer        *Bearer   `json:"bearer" bson:"bearer"`
	Basic         *Basic    `json:"basic" bson:"basic"`
	Digest        *Digest   `json:"digest"`
	Hawk          *Hawk     `json:"hawk"`
	Awsv4         *AwsV4    `json:"awsv4"`
	Ntlm          *Ntlm     `json:"ntlm"`
	Edgegrid      *Edgegrid `json:"edgegrid"`
	Oauth1        *Oauth1   `json:"oauth1"`
}
type Cookie struct {
	Parameter []*VarForm `json:"parameter"`
}
type HttpApiSetup struct {
	ClientName          string `json:"client_name"`
	IsRedirects         int64  `json:"is_redirects"`   // 是否跟随重定向 0: 是   1：否
	RedirectsNum        int    `json:"redirects_num"`  // 重定向次数>= 1; 默认为3
	ReadTimeOut         int64  `json:"read_time_out"`  // 请求读取超时时间
	WriteTimeOut        int64  `json:"write_time_out"` // 响应读取超时时间
	KeepAlive           bool   `json:"keep_alive"`
	MaxIdleConnDuration int64  `json:"max_idle_conn_duration"`
	MaxConnPerHost      int    `json:"max_conn_per_host"`
	UserAgent           bool   `json:"user_agent"`
	MaxConnWaitTimeout  int64  `json:"max_conn_wait_timeout"`
}
type AssertionText struct {
	IsChecked    int    `json:"is_checked"`
	ResponseType int8   `json:"response_type"`
	Compare      string `json:"compare"`
	Var          string `json:"var"`
	Val          string `json:"val"`
	Index        int    `json:"index"`
}
type RegularExpression struct {
	IsChecked int         `json:"is_checked"`
	Type      int         `json:"type"`
	Var       string      `json:"var"`
	Express   string      `json:"express"`
	Index     int         `json:"index"`
	Val       interface{} `json:"val"`
}
type TLS struct {
	CaCert string `json:"ca_cert"`
}
type KV struct {
	Key   string      `json:"key" bson:"key"`
	Value interface{} `json:"value" bson:"value"`
}
type Bearer struct {
	Key string `json:"key" bson:"key"`
}
type Basic struct {
	UserName string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}
type Digest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Realm     string `json:"realm"`
	Nonce     string `json:"nonce"`
	Algorithm string `json:"algorithm"`
	Qop       string `json:"qop"`
	Nc        string `json:"nc"`
	Cnonce    string `json:"cnonce"`
	Opaque    string `json:"opaque"`
}
type Hawk struct {
	AuthID             string `json:"authId"`
	AuthKey            string `json:"authKey"`
	Algorithm          string `json:"algorithm"`
	User               string `json:"user"`
	Nonce              string `json:"nonce"`
	ExtraData          string `json:"extraData"`
	App                string `json:"app"`
	Delegation         string `json:"delegation"`
	Timestamp          string `json:"timestamp"`
	IncludePayloadHash int    `json:"includePayloadHash"`
}
type Ntlm struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Domain              string `json:"domain"`
	Workstation         string `json:"workstation"`
	DisableRetryRequest int    `json:"disableRetryRequest"`
}
type Edgegrid struct {
	AccessToken   string `json:"accessToken"`
	ClientToken   string `json:"clientToken"`
	ClientSecret  string `json:"clientSecret"`
	Nonce         string `json:"nonce"`
	Timestamp     string `json:"timestamp"`
	BaseURi       string `json:"baseURi"`
	HeadersToSign string `json:"headersToSign"`
}
type Oauth1 struct {
	ConsumerKey          string `json:"consumerKey"`
	ConsumerSecret       string `json:"consumerSecret"`
	SignatureMethod      string `json:"signatureMethod"`
	AddEmptyParamsToSign int    `json:"addEmptyParamsToSign"`
	IncludeBodyHash      int    `json:"includeBodyHash"`
	AddParamsToHeader    int    `json:"addParamsToHeader"`
	Realm                string `json:"realm"`
	Version              string `json:"version"`
	Nonce                string `json:"nonce"`
	Timestamp            string `json:"timestamp"`
	Verifier             string `json:"verifier"`
	Callback             string `json:"callback"`
	TokenSecret          string `json:"tokenSecret"`
	Token                string `json:"token"`
}
type AwsV4 struct {
	AccessKey          string `json:"accessKey"`
	SecretKey          string `json:"secretKey"`
	Region             string `json:"region"`
	Service            string `json:"service"`
	SessionToken       string `json:"sessionToken"`
	AddAuthDataToQuery int    `json:"addAuthDataToQuery"`
}

// 获取fasthttp客户端
func fastClient(httpApiSetup *HttpApiSetup, auth *Auth) (fc *fasthttp.Client) {
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
						//log.Logger.Error(fmt.Sprintf("机器ip:%s, 下载crt文件失败：", middlewares.LocalIp), err)
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
