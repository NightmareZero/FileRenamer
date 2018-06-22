package files

// 文件操作

import (
	"os"
	"path"
	"strings"
)

// File ...
// .
type File struct {
	truePath    string // 完整路径
	dirPath     string // 目录路径
	filename    string // 文件名
	fileExt     string // 文件扩展名
	filenameNew string // 新文件名
}

// NewFile ...
// .
func NewFile(filepath string) *File {
	res := new(File)
	res.truePath = strings.Replace(filepath, "\\", "/", -1)
	res.dirPath, res.filename = path.Split(res.truePath)
	res.fileExt = path.Ext(res.truePath)
	res.filename = strings.Replace(res.filename, res.fileExt, "", -1)
	return res
}

// IsExist ...
// 判断文件是否存在
func (that *File) IsExist() bool {
	_, err := os.Stat(that.truePath)
	if err != nil && os.IsExist(err) {
		return false
	}
	return true
}

// GetTruePath ...
// 获取输入名称
func (that *File) GetTruePath() string {
	return that.truePath
}
