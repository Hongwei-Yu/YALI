package model

import (
	uuid "github.com/satori/go.uuid"
)

// Api 结构体
type Api struct {
	TargetId   string    `json:"target_id"`
	Uuid       uuid.UUID `json:"uuid"`
	Name       string    `json:"name"`
	TeamId     string    `json:"team_id"`
	TargetType string    `json:"target_type"` // api/webSocket/tcp/grpc
	Debug      string    `json:"debug"`       // 是否开启Debug模式
	//Request    interface{} `json:"request"`

}

/*
 "target_id": 1900,
  "uuid": "c859604c-d473-4efb-bfd2-795649e69cff",
  "name": "echo",
  "team_id": 175,
  "target_type": "api",
  "method": "POST",
  "request":{}
  "parameters": null,
  "assert": [
    {
      "is_checked": 1,
      "response_type": 2,
      "compare": "gte",
      "var": "errcode",
      "val": "0"
    }
  ],
  "timeout": 0,
  "regex": [
    {
      "is_checked": 1,
      "type": 1,
      "var": "host",
      "express": "errcode",
      "val": ""
    }
  ],
  "debug": "all",
  "connection": 0,
  "variable": null
*/
