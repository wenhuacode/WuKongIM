package trace

import (
	"go.opentelemetry.io/otel"
)

var (
	meter = otel.Meter("metrics")
)

type metrics struct {
	cluster IClusterMetrics
	opts    *Options
}

func newMetrics(opts *Options) *metrics {
	return &metrics{
		cluster: newClusterMetrics(opts),
		opts:    opts,
	}
}

// System 系统监控
func (d *metrics) System() ISystemMetrics {

	return nil
}

// App  应用监控
func (d *metrics) App() IAppMetrics {
	return nil
}

// Cluster 分布式监控
func (d *metrics) Cluster() IClusterMetrics {
	return d.cluster
}

// DB 数据库监控
func (d *metrics) DB() IDBMetrics {
	return nil
}
