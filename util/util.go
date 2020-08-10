package util

import (
	_ "cloudDesktop/util/conf" // init conf

	"cloudDesktop/util/db"
	"cloudDesktop/util/log"
	"cloudDesktop/util/mc"
	"cloudDesktop/util/redis"
)

// GatherMetrics 收集一些被动指标
func GatherMetrics() {
	mc.GatherMetrics()
	redis.GatherMetrics()
}

// Reset all utils
func Reset() {
	log.Reset()
	db.Reset()
	mc.Reset()
}

// Stop all utils
func Stop() {
}
