package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

//	func main() {
//		store.InitInflux()
//		org := "YALI"
//		bucket := "YALI_DEV"
//		writeAPI := store.GTC.Client.WriteAPIBlocking(org, bucket)
//		for value := 0; value < 10; value++ {
//			tags := map[string]string{
//				"tagname1": "tagvalue1",
//			}
//			fields := map[string]interface{}{
//				"field1": value,
//			}
//			point := write.NewPoint("measurement1", tags, fields, time.Now())
//			time.Sleep(1 * time.Second) // separate points by 1 second
//
//			if err := writeAPI.WritePoint(context.Background(), point); err != nil {
//				log.Fatal(err)
//			}
//		}
//		queryAPI := store.GTC.Client.QueryAPI(org)
//		query := `from(bucket: "YALI_DEV")
//	           |> range(start: -10m)
//	           |> filter(fn: (r) => r._measurement == "measurement1")`
//		results, err := queryAPI.Query(context.Background(), query)
//		if err != nil {
//			log.Fatal(err)
//		}
//		for results.Next() {
//			fmt.Println(results.Record().Values())
//		}
//		if err := results.Err(); err != nil {
//			log.Fatal(err)
//		}
//	}
func main() {
	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}

}
