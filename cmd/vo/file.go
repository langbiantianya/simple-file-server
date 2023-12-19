package vo

import (
	"time"
)

type Dir struct {
	RootPath  string
	Parent    string
	NowPath   string
	FileItems []FileItem
}
type FileItem struct {
	Name    string
	Size    int64
	Mode    uint32
	ModTime time.Time
	IsDir   bool
}
