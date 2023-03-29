package cron

import (
	"testing"
	"time"
)

func TestInitCron(t *testing.T) {
	InitCron()
	time.Sleep(time.Minute * 5)
}
