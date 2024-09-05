package model

type Plan struct {
	PlanID      string      `json:"plan_id" bson:"plan_id"`
	PlanName    string      `json:"plan_name" bson:"plan_name"`
	PlanDesc    string      `json:"plan_desc" bson:"plan_desc"` // nm,都到执行了还要啥desc啊
	PlanType    string      `json:"plan_type" bson:"plan_type"`
	ProjectId   string      `json:"project_id" bson:"project_id"`
	CreateUser  string      `json:"create_user" bson:"create_user"`
	StartTimer  string      `json:"start_timer" bson:"start_timer"`
	EngineNum   int         `json:"engine_num" bson:"engine_num"`
	PlanConfig  *PlanConfig `json:"plan_config" bson:"plan_config"`
	Scene       interface{} `json:"scane" bson:"scane"`
	SceneConfig interface{} `json:"scene_config" bson:"scene_config"`
}
type PlanConfig struct {
	Mode       string      `json:"mode" bson:"mode"` //cnm,搞几种模式呢？
	ModeConfig *ModeConfig `json:"mode_config" bson:"mode_config"`
}
type ModeConfig struct {
	RoundNum         int64 `json:"round_num"`         // 轮次
	Concurrency      int64 `json:"concurrency"`       // 并发数
	StartConcurrency int64 `json:"start_concurrency"` // 起始并发数
	Step             int64 `json:"step"`              // 并发步长
	StepRunTime      int64 `json:"step_run_time"`     // 步长持续时间
	MaxConcurrency   int64 `json:"max_concurrency"`   // 最大并发数
	Duration         int64 `json:"duration"`          // 稳定持续市场
}
