package test

import (
	"sys/timetool"
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	test := timetool.GetTime(time.Now())
	t.Log(test)
}
