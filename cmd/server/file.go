package server

import (
	"io/fs"
	"log"
	"os"
)

type Basic interface {
	Ls() SelectedPath
}

type SelectedPath struct {
	RootPath string
	Parent   string
	NowPath  string
	Entries  []fs.DirEntry
}

func (b *SelectedPath) Ls() *SelectedPath {
	dir, err := os.Open(b.RootPath)
	if err != nil {
		log.Default().Println(err)
	}
	defer dir.Close()
	entries, err := dir.ReadDir(0)
	if err != nil {
		println(err)
	}
	b.Entries = entries
	return b
}
