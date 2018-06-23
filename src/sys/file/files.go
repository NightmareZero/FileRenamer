package file

// 文件操作

import (
	"os"
	"path"
	"strings"
	"sys/ostl"
)

// File ...
type File struct {
	path string // 完整路径
	Dir  string // 目录路径
	Name string // 文件名
	Ext  string // 文件扩展名
}

// NewFile ... Construcotr
func NewFile(filepath string) *File {
	res := new(File)
	res.setPath(filepath)
	res.Split()
	return res
}

func (that *File) setPath(filepath string) {
	that.path = strings.Replace(filepath, "\\", "/", -1)
}

// Split ... 分析路径
func (that *File) Split() {
	that.Dir, that.Name = path.Split(that.path)
	that.Ext = path.Ext(that.path)
	that.Name = strings.Replace(that.Name, that.Ext, "", -1)
}

// GetPath ... 获取设置的路径
func (that *File) GetPath() string {
	rev := that.path
	if ostl.IsWindows() {
		return strings.Replace(rev, "/", "\\", -1)
	}
	return rev
}

// GetNewPath ... 获取填入的新路径
func (that *File) GetNewPath() string {
	rev := that.Dir + that.Name + that.Ext
	if ostl.IsWindows() {
		return strings.Replace(rev, "/", "\\", -1)
	}
	return rev
}

// IsExist ...
// 判断文件是否存在
func (that *File) IsExist() bool {
	_, err := os.Stat(that.path)
	if err != nil && os.IsExist(err) {
		return false
	}
	return true
}

// Rename ... 重命名一个文件 .
func Rename(file *File) (err error) {
	err = nil
	err = os.Rename(file.GetPath(), file.GetNewPath())
	file.setPath(file.GetNewPath())
	return
}
