package files

// File ...
// .
type File struct {
	TruePath string // 完整路径
	DirPath  string // 目录路径
	Name     string // 文件名
	Ext      string // 文件扩展名
	NameNew  string // 新文件名
}

// NewFile ...
// .
func NewFile(path string) *File {
	res := new(File)
	res.TruePath = path

	return res
}
