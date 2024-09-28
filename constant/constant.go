package constant

// 返回 code 码
const (
	// NoError 没有错误
	NoError = int64(10000)
	// AssertError 断言错误
	AssertError = int64(10001)
	// RequestError 请求错误
	RequestError = int64(10002)
	// ServiceError 服务错误
	ServiceError = int64(10003)
)

type Reg struct {
	Key   string      `json:"key" bson:"key"`
	Value interface{} `json:"value" bson:"value"`
}

type Regex struct {
	Regs []*Reg `json:"regs" bson:"regs"`
}

const (
	Close = 0 // 关闭
	Open  = 1 // 开启
)

// HTTP协议
const (
	HTTP  = "http"  // http
	HTTPS = "https" // https

)

// 数据类型
const (
	StringType    = "String"
	TextType      = "Text"
	ObjectType    = "Object"
	ArrayType     = "Array"
	IntegerType   = "Integer"
	NumberType    = "Number"
	FloatType     = "Float"
	DoubleType    = "Double"
	FileType      = "File"
	FileUrlType   = "FileUrl"
	DateType      = "Date"
	DateTimeType  = "DateTime"
	TimeStampType = "TimeStampType"
	BooleanType   = "boolean"
	InterfaceMap  = "map[interface {}]interface {}"
)

// body 格式
const (
	NoneMode      = "none"
	FormMode      = "form-data"
	UrlencodeMode = "urlencoded"
	JsonMode      = "json"
	XmlMode       = "xml"
	JSMode        = "javascript"
	PlainMode     = "plain"
	HtmlMode      = "html"
	MysqlMode     = "sql"
)
const (
	RegExtract    = 0
	JsonExtract   = 1
	HeaderExtract = 2
	CodeExtract   = 3
)

// debug日志状态
const (
	All         = "all"
	OnlyError   = "only_error"
	OnlySuccess = "only_success"
	StopDebug   = "stop"
)
