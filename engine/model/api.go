package model

import "github.com/google/uuid"

type Api struct {
	TargetId   string      `json:"target_id"`
	Uuid       uuid.UUID   `json:"uuid"`
	Name       string      `json:"name"`
	TeamId     string      `json:"team_id"`
	TargetType string      `json:"target_type"` // api/webSocket/tcp/grpc
	Debug      string      `json:"debug"`       // 是否开启Debug模式
	Request    RequestHttp `json:"request"`
}
