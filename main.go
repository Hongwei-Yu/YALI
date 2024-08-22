package main

import (
	"YALI/config"
	"YALI/store"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

func main() {
	config.InitConfig()
	store.InitDB()
	for i := 0; i < 1000; i++ {
		go testbingfawrite(i)
	}
	time.Sleep(100 * time.Second)
}

func testbingfawrite(i int) {
	//start := time.Now()
	err := store.GTW.Write(influxdb2.NewPoint("test", map[string]string{"ce": "shi"}, map[string]interface{}{"start": i}, time.Now()))
	if err != nil {
		return
	}
}
