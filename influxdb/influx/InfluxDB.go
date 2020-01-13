package influx

import (
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"log"
	"time"
)

type DBClient interface {
	Build()
	Insert(tags *map[string]string, fields *map[string]interface{})
	Query(sql string) (res []client.Result, err error)
}

type InfluxDB struct {
	Url       string
	Name      string
	Pwd       string
	Db        string
	Mmt       string
	Client    client.Client
	bps       client.BatchPoints
	Precision string // 时间精度 s、ms、ns
}

func main() {
	dbClient := InfluxDB{
		Url:       "http://127.0.0.1:8086",
		Name:      "admin",
		Pwd:       "",
		Db:        "logprocess",
		Mmt:       "log",
		Precision: "ms",
	}
	dbClient.Build()

	tags := &map[string]string{"value": "100"}
	field := &map[string]interface{}{
		"t":    time.Now(),
		"data": "aaaaaa",
	}
	dbClient.Insert(tags, field)
	res, _ := dbClient.Query("select * from log")
	fmt.Println(res)
	fmt.Println(res[0])
}

// build influx DB client
func (influx *InfluxDB) Build() {
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     influx.Url,
		Username: influx.Name,
		Password: influx.Pwd,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	influx.Client = cli
}

//query
func (influx *InfluxDB) Query(cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: influx.Db,
	}
	if response, err := influx.Client.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

//Insert
func (influx *InfluxDB) Insert(tags *map[string]string, fields *map[string]interface{}) {
	if influx.bps == nil {
		influx.bps, _ = client.NewBatchPoints(client.BatchPointsConfig{
			Database:  influx.Db,
			Precision: influx.Precision,
		})
	}

	pt, err := client.NewPoint(
		influx.Mmt,
		*tags,
		*fields,
		time.Now(),
	)
	if err != nil {
		log.Fatal(err)
	}
	influx.bps.AddPoint(pt)
	//if len(influx.bps.Points()) == 1000 {
	if err := influx.Client.Write(influx.bps); err != nil {
		log.Fatal(err)
	}
	influx.bps, _ = client.NewBatchPoints(client.BatchPointsConfig{
		Database:  influx.Db,
		Precision: influx.Precision,
	})
	//}
}
