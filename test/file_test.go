package test

import (
	"os"
	"testing"
)

func TestPath(t *testing.T) {
	rootPath := "/home/lbty/goProject/simpleFileServer"
	dir, err := os.Open(rootPath)
	if err != nil {
		println(err)
	}
	defer dir.Close()
	entries, err := dir.ReadDir(0)
	if err != nil {
		println(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			println("目录:", entry.Name())
		} else {
			println("文件:", entry.Name())
		}
	}
}
