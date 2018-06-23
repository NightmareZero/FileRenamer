package test

import (
	"sys/ostl"
	"testing"
)

func Test1(t *testing.T) {
	if !ostl.IsWindows() {
		t.Error()
	}
}
