package server

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

func Default(rootPath string) *SelectedPath {
	return &SelectedPath{
		// Passwd:   passwd,
		RootPath: rootPath,
		Parent:   "/",
		NowPath:  "/",
		Entries:  []fs.DirEntry{},
	}
}

type SelectedPath struct {
	Passwd   string
	RootPath string
	Parent   string
	NowPath  string
	Entries  []fs.DirEntry
}

func (b *SelectedPath) Ls(err error) error {
	if err != nil {
		return err
	}
	dir, err := os.Open(filepath.Clean(b.NowPath))
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
	re := regexp.MustCompile(`^(./)|^(/)`)
	new = re.ReplaceAllString(new, "")
	err := os.MkdirAll(filepath.Clean(fmt.Sprintf("%s/%s", b.NowPath, new)), 0755)
	return b.Ls(err)
}

func (b *SelectedPath) Touch(name string, file io.Reader) error {
	dstPath := filepath.Clean(fmt.Sprintf("%s/%s", b.NowPath, name))
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

func (b *SelectedPath) IsDir() bool {
	info, _ := os.Stat(filepath.Clean(b.NowPath))
	return info.IsDir()
}

func (b *SelectedPath) Remove() error {
	path := filepath.Clean(b.NowPath)
	if b.IsDir() {
		if path == "./" || path == "/" || path == filepath.Clean(b.RootPath) {
			return fmt.Errorf("禁止删除该文件夹")
		} else {
			return os.RemoveAll(path)
		}
	} else {
		return os.Remove(path)
	}
}
