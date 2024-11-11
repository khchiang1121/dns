package forward

import (
	"github.com/coredns/coredns/plugin"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Variables declared for monitoring.
var (
	RequestCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "forward",
		Name:      "requests_total",
		Help:      "Counter of requests made per upstream.",
	}, []string{"server", "zones", "view", "to", "type"})
	RcodeCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "forward",
		Name:      "responses_total",
		Help:      "Counter of responses received per upstream.",
	}, []string{"server", "zones", "view", "rcode", "to"})
	RequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: plugin.Namespace,
		Subsystem: "forward",
		Name:      "request_duration_seconds",
		Buckets:   plugin.TimeBuckets,
		Help:      "Histogram of the time each request took.",
	}, []string{"server", "zones", "view", "to", "rcode", "type"})
	HealthcheckCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "forward",
		Name:      "healthcheck_total",
		Help:      "Counter of the number of healthchecks.",
	}, []string{"server", "zones", "view", "to"})
	HealthcheckFailureCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "forward",
		Name:      "healthcheck_failures_total",
		Help:      "Counter of the number of failed healthchecks.",
	}, []string{"server", "zones", "view", "to"})
	HealthcheckBrokenCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "forward",
		Name:      "healthcheck_broken_total",
		Help:      "Counter of the number of complete failures of the healthchecks.",
	}, []string{"server", "zones", "view"})
	HealthcheckBrokenSkipForwardCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "forward",
		Name:      "healthcheck_broken_skip_forward_total",
		Help:      "Counter of requests not forwarded due to skipForward being enabled when all upstream servers are marked as down.",
	}, []string{"server", "zones", "view"})
	MaxConcurrentRejectCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "forward",
		Name:      "max_concurrent_rejects_total",
		Help:      "Counter of the number of queries rejected because the concurrent queries were at maximum.",
	}, []string{"server", "zones", "view"})
	ConnCacheHitsCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "forward",
		Name:      "conn_cache_hits_total",
		Help:      "Counter of connection cache hits per upstream and protocol.",
	}, []string{"to", "proto"})
	ConnCacheMissesCount = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: plugin.Namespace,
		Subsystem: "forward",
		Name:      "conn_cache_misses_total",
		Help:      "Counter of connection cache misses per upstream and protocol.",
	}, []string{"to", "proto"})
)
