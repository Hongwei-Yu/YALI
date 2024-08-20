package store

import (
	"YALI/store/timestorage"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"log"
)

var Global timestorage.InfluxConnect

type Connectors struct {
}

func InitInflux() {
	Global = timestorage.InfluxConnect{Proto: "http", Host: "172.24.3.148", Port: "8086",
		Token: "RKwNxaDr_9Oc35asdXNBVUEe7eLzdQEAUtuW46E1thCD501dc5GaAlM2FxifrAJ9SldH07sGray9750eu7mM_g==",
	}
	Global.Client = influxdb2.NewClient(Global.Proto+"://"+Global.Host+":"+Global.Port, Global.Token)
	if Global.Client == nil {
		log.Fatalln("connect error")
	}
}
