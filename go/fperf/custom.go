package main_perf

import (
	"fmt"
	"github.com/fperf/fperf"
	"math/rand"
	"time"
)

// fperf-build ./
// ./fperf -cpu 8 -connection 100 -goroutine 8 -tick 2s -N 12 demo https://baidu.com.com
// ./fperf -cpu 16 -connection 100 -goroutine 16 -tick 2s -N 12 demo https://baidu.com.com
func init() {
	fperf.Register("demo", newDemoClient, "This is a demo client discription")
}

type demoClient struct{}

func newDemoClient(flag *fperf.FlagSet) fperf.Client {
	return &demoClient{}
}

func (c *demoClient) Dial(addr string) error {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)))
	fmt.Println("Dial to", addr, time.Now().Unix())
	// todo send to influxdb and display by grafana
	return nil
}
