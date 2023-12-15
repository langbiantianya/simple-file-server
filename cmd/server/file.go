package server

import "io/fs"

type Basic struct {
	RootPath string
}

type SelectedPath struct {
	Basic   Basic
	Parent  string
	Entries []fs.DirEntry
}
