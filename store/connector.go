package store

import (
	"YALI/store/timestorage"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"log"
)

var GTC timestorage.InfluxConnect // global time connector
var GTW timestorage.InfluxWriter  // global time writer

type Connectors struct {
}

func InitInflux() {
	GTC = timestorage.InfluxConnect{Proto: "http", Host: "192.168.0.110", Port: "8086",
		Token: "zdng0zJ7Wc69NbuD7lcot6cToX_UacEmspIu4oKGS368_sdkZqq8PjChHRsMOZrrvBu270sKSW5rRR4uVafDYQ==",
	}
	GTW = timestorage.InfluxWriter{Org: "YALI", Bucket: "YALI_DEV"}

	GTC.Client = influxdb2.NewClient(GTC.Proto+"://"+GTC.Host+":"+GTC.Port, GTC.Token)
	if GTC.Client == nil {
		log.Fatalln("connect error")
	}
	GTW.Writer = GTC.Client.WriteAPIBlocking(GTW.Org, GTW.Bucket)
}
