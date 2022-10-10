package influxd

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
)

func TestInit(t *testing.T) {
	// Create a new client using an InfluxDB server base URL and an authentication token
	client := influxdb2.NewClient("http://localhost:8086", "PYroYIGd4X90gLiKK0F_VKVTDXt3BxjAJtUNJfk9eFDD7kiSMSoOHocBT0Cqz_B_ZYT0UySO52TsiJGSYWY9NA==")
	// 获取写客户端
	writeAPI := client.WriteAPIBlocking("maxno", "test")
	// 全部参数创建一行数据
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45.0},
		time.Now())
	// measurement对应表，第一次创建数据时自动建表
	if e := writeAPI.WritePoint(context.Background(), p); e != nil {
		panic(e.Error())
	}
	// fluent风格创建数据
	p = influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 23.2).
		AddField("max", 45.0).
		SetTime(time.Now())
	if e := writeAPI.WritePoint(context.Background(), p); e != nil {
		panic(e.Error())
	}

	// 以数据行直接创建数据
	line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 23.5, 45.0)
	if e := writeAPI.WriteRecord(context.Background(), line); e != nil {
		panic(e.Error())
	}

	// 获取查询客户端
	queryAPI := client.QueryAPI("maxno")
	// Get parser flux query result
	result, err := queryAPI.Query(context.Background(), `from(bucket:"test")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`)
	if err == nil {
		// Use Next() to iterate over query result lines
		for result.Next() {
			// Observe when there is new grouping key producing new table
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			// read result
			fmt.Printf("row: %s\n", result.Record().String())
		}
		if result.Err() != nil {
			fmt.Printf("Query error: %s\n", result.Err().Error())
		}
	} else {
		panic(err.Error())
	}
	// Ensures background processes finishes
	client.Close()
}
