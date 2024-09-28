package model

const (
	ConcurrentModel = 1 // 并发数模式
	LadderModel     = 2 // 阶梯模式
	ErrorRateModel  = 3 // 错误率模式
	RTModel         = 4 // 响应时间模式
	RpsModel        = 5 // 每秒请求数模式
	RoundModel      = 6 // 轮次模式

)

// ConfigTask 任务配置
type ConfigTask struct {
	DebugMode string `json:"debug_mode"` // 开启debug模式： all: 所有日志; only_error: 仅错误日志；  only_success：仅成功日志：stop： 停止debug;
	TaskType  int64  `json:"task_type"`  // 任务类型：0. 普通任务； 1. 定时任务； 2. cicd任务
	//ControlMode int64    `json:"control_mode"` // 控制模式： 0. 集中模式； 1. 单独模式
	Mode     int64    `json:"mode"`      // 压测模式 1:并发模式，2:阶梯模式，3:错误率模式，4:响应时间模式，5:每秒请求数模式
	Remark   string   `json:"remark"`    // 备注
	CronExpr string   `json:"cron_expr"` // 定时任务表达式
	ModeConf ModeConf `json:"mode_conf"` // 任务配置
}

type ModeConf struct {
	RoundNum         int64 `json:"round_num"`         // 轮次
	Concurrency      int64 `json:"concurrency"`       // 并发数
	StartConcurrency int64 `json:"start_concurrency"` // 起始并发数
	Step             int64 `json:"step"`              // 并发步长
	StepRunTime      int64 `json:"step_run_time"`     // 步长持续时间
	MaxConcurrency   int64 `json:"max_concurrency"`   // 最大并发数
	Duration         int64 `json:"duration"`          // 持续市场
}
