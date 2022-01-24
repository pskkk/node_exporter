package collector

import (
	"fmt"
	"github.com/go-kit/log"
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
)

type testCollector struct {
	logger log.Logger
}

var testDesc = prometheus.NewDesc(
	prometheus.BuildFQName(namespace, "diy", "test_collectors"),
	"diy test node exporter collectors",
	nil, nil,
)

func init() {
	registerCollector("test", true, NewtestCollector)
}

func NewtestCollector(logger log.Logger) (Collector, error) {
	return &testCollector{logger: logger}, nil
}

func (t *testCollector) Update(ch chan<- prometheus.Metric) error {
	n := getRandomInt()
	fmt.Println(n, float64(n))
	ch <- prometheus.MustNewConstMetric(testDesc, prometheus.GaugeValue, float64(n))
	return nil
}

func getRandomInt() int {
	return rand.Intn(100)
}
