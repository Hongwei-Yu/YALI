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
	Global = timestorage.InfluxConnect{Proto: "http", Host: "192.168.0.110", Port: "8086",
		Token: "zdng0zJ7Wc69NbuD7lcot6cToX_UacEmspIu4oKGS368_sdkZqq8PjChHRsMOZrrvBu270sKSW5rRR4uVafDYQ==",
	}
	Global.Client = influxdb2.NewClient(Global.Proto+"://"+Global.Host+":"+Global.Port, Global.Token)
	if Global.Client == nil {
		log.Fatalln("connect error")
	}
}
