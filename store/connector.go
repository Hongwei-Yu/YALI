package store

import (
	"YALI/config"
	"YALI/store/relastorage"
	"YALI/store/timestorage"
	"database/sql"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"log"
)

var GTC timestorage.InfluxConnect // global time connector
var GTW timestorage.InfluxWriter  // global time writer

var GRC relastorage.SqliteConnect

type Connectors struct {
}

func InitInflux() {
	GTC = timestorage.InfluxConnect{
		Proto: config.Gconfig.GetString("DB.time.proto"),
		Host:  config.Gconfig.GetString("DB.time.host"),
		Port:  config.Gconfig.GetString("DB.time.port"),
		Token: config.Gconfig.GetString("DB.time.token"),
	}
	GTW = timestorage.InfluxWriter{
		Org:    config.Gconfig.GetString("DB.time.org"),
		Bucket: config.Gconfig.GetString("DB.time.buck"),
	}

	GTC.Client = influxdb2.NewClient(GTC.Proto+"://"+GTC.Host+":"+GTC.Port, GTC.Token)
	if GTC.Client == nil {
		log.Fatalln("timedb server connect error")
	}
	log.Println("timedb server connect sucsess")
	GTW.Writer = GTC.Client.WriteAPIBlocking(GTW.Org, GTW.Bucket)

}

func InitSqlite() {
	GRC = relastorage.SqliteConnect{
		Servername: config.Gconfig.GetString("DB.rela.servername"),
		DBname:     config.Gconfig.GetString("DB.rela.dbname"),
	}
	var err error
	GRC.Client, err = sql.Open(GRC.Servername, GRC.DBname)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("relationDB server connect sucess")
}

func InitDB() {
	InitSqlite()
	InitInflux()
}
