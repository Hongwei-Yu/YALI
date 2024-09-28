package model

import (
	"YALI/constant"
	"YALI/global"
	"YALI/kit"
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/valyala/fasthttp"
	"io"
	"math"
	"mime/multipart"
	"net/textproto"
	"net/url"
	"strings"
	"sync"
)

type VarForm struct {
	IsChecked   int64       `json:"is_checked" bson:"is_checked"`   // 是否选中
	Type        string      `json:"type" bson:"type"`               // 参数类型
	FileBase64  []string    `json:"fileBase64"`                     // 文件Base64
	Key         string      `json:"key" bson:"key"`                 // 参数名称
	Value       interface{} `json:"value" bson:"value"`             // 参数值
	NotNull     int64       `json:"not_null" bson:"not_null"`       // 是否必填
	Description string      `json:"description" bson:"description"` // 参数描述
	FieldType   string      `json:"field_type" bson:"field_type"`   // 参数类型
}

func (v *VarForm) ValueToByte() (by []byte) {
	if v.Value == nil {
		return
	}
	switch v.Type {
	case constant.StringType:
		by = []byte(v.Value.(string))
	case constant.TextType:
		by = []byte(v.Value.(string))
	case constant.ObjectType:
		by = []byte(v.Value.(string))
	case constant.ArrayType:
		by = []byte(v.Value.(string))
	case constant.NumberType:
		bytesBuffer := bytes.NewBuffer([]byte{})
		_ = binary.Write(bytesBuffer, binary.BigEndian, v.Value.(int))
		by = bytesBuffer.Bytes()
	case constant.IntegerType:
		bytesBuffer := bytes.NewBuffer([]byte{})
		_ = binary.Write(bytesBuffer, binary.BigEndian, v.Value.(int))
		by = bytesBuffer.Bytes()
	case constant.DoubleType:
		bytesBuffer := bytes.NewBuffer([]byte{})
		_ = binary.Write(bytesBuffer, binary.BigEndian, v.Value.(int64))
		by = bytesBuffer.Bytes()
	case constant.FileType:
		bits := math.Float64bits(v.Value.(float64))
		binary.LittleEndian.PutUint64(by, bits)
	case constant.BooleanType:
		buf := bytes.Buffer{}
		enc := gob.NewEncoder(&buf)
		_ = enc.Encode(v.Value.(bool))
		by = buf.Bytes()
	case constant.DateType:
		by = []byte(v.Value.(string))
	case constant.DateTimeType:
		by = []byte(v.Value.(string))
	case constant.TimeStampType:
		bytesBuffer := bytes.NewBuffer([]byte{})
		_ = binary.Write(bytesBuffer, binary.BigEndian, v.Value.(int64))
		by = bytesBuffer.Bytes()

	}
	return
}

type Header struct {
	Parameter []*VarForm `json:"parameter"`
}

func (header *Header) SetHeader(req *fasthttp.Request) {
	if header == nil || header.Parameter == nil {
		return
	}
	for _, v := range header.Parameter {

		if v.IsChecked != constant.Open || v.Value == nil {
			continue
		}
		if strings.EqualFold(v.Key, "content-type") {
			req.Header.SetContentType(v.Value.(string)) // 开启content-type头部
		}
		if strings.EqualFold(v.Key, "host") {
			req.SetHost(v.Value.(string))
			req.UseHostHeader = true // 开启host头部
		}
		req.Header.Set(v.Key, v.Value.(string))
	}
}

type Query struct {
	Parameter []*VarForm `json:"parameter"`
}

type Body struct {
	Mode      string     `json:"mode"`
	Raw       string     `json:"raw"`
	Parameter []*VarForm `json:"parameter"`
}

func (b *Body) SetBody(req *fasthttp.Request) string {
	if b == nil {
		return ""
	}
	switch b.Mode {
	case constant.NoneMode:
	case constant.FormMode:
		req.Header.SetContentType("multipart/form-data")
		// 新建一个缓冲，用于存放文件内容

		if b.Parameter == nil {
			b.Parameter = []*VarForm{}
		}

		bodyBuffer := &bytes.Buffer{}
		bodyWriter := multipart.NewWriter(bodyBuffer)
		contentType := bodyWriter.FormDataContentType()
		//var fileTypeList []string
		for _, value := range b.Parameter {

			if value.IsChecked != constant.Open {
				continue
			}
			if value.Key == "" {
				continue
			}

			switch value.Type {
			case constant.FileType:
				if value.FileBase64 == nil || len(value.FileBase64) < 1 {
					continue
				}
				for _, base64Str := range value.FileBase64 {
					by, fileType := kit.Base64DeEncode(base64Str, constant.FileType)
					log.Debug(fmt.Sprintf("机器ip:%s, fileType:    ", global.GlobalEngine.Host), fileType)
					if by == nil {
						continue
					}
					//fileWriter, err := bodyWriter.CreateFormFile(value.Key, value.Value.(string))
					h := make(textproto.MIMEHeader)
					h.Set("Content-Type", fileType)
					h.Set("Content-Disposition",
						fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
							value.Key, value.Value.(string)))
					fileWriter, err := bodyWriter.CreatePart(h)
					if err != nil {
						log.Error(fmt.Sprintf("机器ip:%s, CreateFormFile失败：%s ", global.GlobalEngine.Host, err.Error()))
						continue
					}
					file := bytes.NewReader(by)
					_, err = io.Copy(fileWriter, file)
					if err != nil {
						continue
					}
				}
			case constant.FileUrlType:
				val, ok := value.Value.(string)
				if !ok {
					continue
				}
				if strings.HasPrefix(val, "https://") || strings.HasPrefix(val, "http://") {
					strList := strings.Split(val, "/")
					if len(strList) < 1 {
						continue
					}
					fileTypeList := strings.Split(strList[len(strList)-1], ".")
					if len(fileTypeList) < 1 {
						continue
					}
					fc := &fasthttp.Client{}
					loadReq := fasthttp.AcquireRequest()
					defer loadReq.ConnectionClose()
					// set url
					loadReq.Header.SetMethod("GET")
					loadReq.SetRequestURI(val)
					loadResp := fasthttp.AcquireResponse()
					defer loadResp.ConnectionClose()
					if err := fc.Do(loadReq, loadResp); err != nil {
						log.Error(fmt.Sprintf("机器ip:%s, 下载body上传文件错误：", global.GlobalEngine.Host), err)
						continue
					}
					if loadResp.StatusCode() != 200 {
						continue
					}

					if loadResp.Body() == nil {
						continue
					}
					h := make(textproto.MIMEHeader)
					h.Set("Content-Type", fileTypeList[len(fileTypeList)-1])
					h.Set("Content-Disposition",
						fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
							value.Key, strList[len(strList)-1]))
					fileWriter, err := bodyWriter.CreatePart(h)
					if err != nil {
						log.Error(fmt.Sprintf("机器ip:%s, CreateFormFile失败：%s ", global.GlobalEngine.Host, err.Error()))
						continue
					}
					file := bytes.NewReader(loadResp.Body())
					_, err = io.Copy(fileWriter, file)
					if err != nil {
						continue
					}
				}

			default:
				filedWriter, err := bodyWriter.CreateFormField(value.Key)
				by := value.ValueToByte()
				filed := bytes.NewReader(by)
				_, err = io.Copy(filedWriter, filed)
				if err != nil {
					//log.Logger.Error(fmt.Sprintf("机器ip:%s, CreateFormFile失败： %s", global.GlobalEngine.Host, err.Error()))
					continue
				}
			}

		}
		bodyWriter.Close()
		req.Header.SetContentType(contentType)
		if bodyBuffer.Bytes() != nil && bodyBuffer.Len() != 68 {
			req.SetBody(bodyBuffer.Bytes())
		}
		return bodyBuffer.String()
	case constant.UrlencodeMode:

		req.Header.SetContentType("application/x-www-form-urlencoded")
		args := url.Values{}

		for _, value := range b.Parameter {
			if value.IsChecked != constant.Open || value.Key == "" || value.Value == nil {
				continue
			}
			args.Add(value.Key, value.Value.(string))

		}
		req.SetBodyString(args.Encode())
		return args.Encode()

	case constant.XmlMode:
		req.Header.SetContentType("application/xml")
		req.SetBodyString(b.Raw)
		return b.Raw
	case constant.JSMode:
		req.Header.SetContentType("application/javascript")
		req.SetBodyString(b.Raw)
		return b.Raw
	case constant.PlainMode:
		req.Header.SetContentType("text/plain")
		req.SetBodyString(b.Raw)
		return b.Raw
	case constant.HtmlMode:
		req.Header.SetContentType("text/html")
		req.SetBodyString(b.Raw)
		return b.Raw
	case constant.JsonMode:
		req.Header.SetContentType("application/json")
		req.SetBodyString(b.Raw)
		return b.Raw
	}
	return ""
}

// Auth 认证方式
type Auth struct {
	Type          string    `json:"type" bson:"type"`     // 认证方式
	Bidirectional *TLS      `json:"bidirectional"`        // 双向认证
	KV            *KV       `json:"kv" bson:"kv"`         // 参数
	Bearer        *Bearer   `json:"bearer" bson:"bearer"` // Bearer
	Basic         *Basic    `json:"basic" bson:"basic"`   // Basic
	Digest        *Digest   `json:"digest"`               // Digest
	Hawk          *Hawk     `json:"hawk"`                 // Hawk
	Awsv4         *AwsV4    `json:"awsv4"`                // AWS V4
	Ntlm          *Ntlm     `json:"ntlm"`                 // NTLM
	Edgegrid      *Edgegrid `json:"edgegrid"`             // 边缘网关
	Oauth1        *Oauth1   `json:"oauth1"`               // OAuth1
}

type Cookie struct {
	Parameter []*VarForm `json:"parameter"`
}

func (cookie *Cookie) SetCookie(req *fasthttp.Request) {
	if cookie == nil || cookie.Parameter == nil {
		return
	}
	for _, v := range cookie.Parameter {
		if v.IsChecked != constant.Open || v.Value == nil || v.Key == "" {
			continue
		}
		req.Header.SetCookie(v.Key, v.Value.(string))
	}
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

func (re RegularExpression) Extract(resp *fasthttp.Response, globalVar *sync.Map) (value interface{}) {
	re.Var = strings.TrimSpace(re.Var)
	name := kit.VariablesMatch(re.Var)
	if name == "" {
		return
	}
	re.Express = strings.TrimSpace(re.Express)
	keys := kit.FindAllDestStr(re.Express, "{{(.*?)}}")
	if keys != nil {
		for _, key := range keys {
			if len(key) < 2 {
				continue
			}
			realVar := kit.ParsFunc(key[1])
			if realVar != key[1] {
				re.Express = strings.Replace(re.Express, key[0], realVar, -1)
				continue
			}
			if v, ok := globalVar.Load(key[1]); ok {
				if v == nil {
					continue
				}
				re.Express = strings.Replace(re.Express, key[0], v.(string), -1)
			}
		}
	}
	switch re.Type {
	case constant.RegExtract:
		if re.Express == "" {
			value = ""
			globalVar.Store(name, value)
			return
		}
		value = kit.FindAllDestStr(string(resp.Body()), re.Express)
		if value == nil || len(value.([][]string)) < 1 {
			value = ""
		} else {
			value = value.([][]string)[0][1]
		}
		globalVar.Store(name, value)
	case constant.JsonExtract:
		value = kit.JsonPath(string(resp.Body()), re.Express)
		globalVar.Store(name, value)
	case constant.HeaderExtract:
		if re.Express == "" {
			value = ""
			globalVar.Store(name, value)
			return
		}
		value = kit.MatchString(resp.Header.String(), re.Express, re.Index)
		globalVar.Store(name, value)
	case constant.CodeExtract:
		value = resp.StatusCode()
		globalVar.Store(name, value)
	}
	return
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
