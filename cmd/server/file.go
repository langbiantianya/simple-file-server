package server

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
)

type Basic interface {
	Ls() *SelectedPath
	Mkdir(new string) *SelectedPath
	Touch(name string, file bytes.Reader) *SelectedPath
}

type SelectedPath struct {
	RootPath string
	Parent   string
	NowPath  string
	Entries  []fs.DirEntry
}

func (b *SelectedPath) Ls(err error) error {
	if err != nil {
		return err
	}
	dir, err := os.Open(b.RootPath)
	if err != nil {
		return err
	}
	defer dir.Close()
	entries, err := dir.ReadDir(0)
	if err != nil {
		return err
	}
	b.Entries = entries
	return nil
}

func (b *SelectedPath) Mkdir(new string) error {
	err := os.MkdirAll(fmt.Sprintf("%s/%s", b.NowPath, new), 0755)
	return b.Ls(err)
}

func (b *SelectedPath) Touch(name string, file io.Reader) error {
	dstPath := fmt.Sprintf("%s/%s", b.NowPath, name)
	dstFile, err := os.Create(dstPath)
	if err != nil {
		return b.Ls(err)
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, file)
	if err != nil {
		return b.Ls(err)
	}
	return b.Ls(err)
}
