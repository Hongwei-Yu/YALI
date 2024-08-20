package timestorage

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type InfluxConnect struct {
	Proto  string
	Host   string
	Port   string
	Token  string
	Client influxdb2.Client
}
