package timestorage

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type InfluxConnect struct {
	Proto  string
	Host   string
	Port   string
	Token  string
	Client influxdb2.Client
}

type InfluxWriter struct {
	Org    string
	Bucket string
	Writer api.WriteAPIBlocking
}
