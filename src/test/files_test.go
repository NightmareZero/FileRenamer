package test

import (
	"files"
	"testing"
)

func TestNewFile(t *testing.T) {
	a := files.NewFile("d:\\b.txt")
	t.Error("1", a)
}
