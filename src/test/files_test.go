package test

import (
	"sys/file"
	"testing"
)

func TestNewFile(t *testing.T) {
	a := file.NewFile("d:\\b.txt")
	t.Error("1", a)
}

func TestRename(t *testing.T) {
	fil := file.NewFile("d:\\b\\b.txt")
	fil.Dir = "d:\\"
	err := file.Rename(fil)
	if err != nil {
		t.Error(err)
	}
}
