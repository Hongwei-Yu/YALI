package main

import (
	"YALI/store"
	"context"
	"fmt"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"log"
	"time"
)

func main() {
	store.InitInflux()
	org := "YALI"
	bucket := "dev"
	writeAPI := store.Global.Client.WriteAPIBlocking(org, bucket)
	for value := 0; value < 10; value++ {
		tags := map[string]string{
			"tagname1": "tagvalue1",
		}
		fields := map[string]interface{}{
			"field1": value,
		}
		point := write.NewPoint("measurement1", tags, fields, time.Now())
		time.Sleep(1 * time.Second) // separate points by 1 second

		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			log.Fatal(err)
		}
	}
	queryAPI := store.Global.Client.QueryAPI(org)
	query := `from(bucket: "dev")
            |> range(start: -10m)`
	results, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		fmt.Println(results.Record().Values())
	}
	if err := results.Err(); err != nil {
		log.Fatal(err)
	}
}
