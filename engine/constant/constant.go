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
