package common

import (
	"github.com/test/go/support/log"
)

const TestAmountPrecision = 7

func CreateLogger(serviceName string) *log.Entry {
	return log.DefaultLogger.WithField("service", serviceName)
}
