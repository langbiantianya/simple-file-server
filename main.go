package main

import (
	"io/fs"
	"simpleFileServer/cmd"
	"simpleFileServer/cmd/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	selectedPath := &server.SelectedPath{
		RootPath: "",
		Parent:   "",
		NowPath:  "",
		Entries:  []fs.DirEntry{},
	}
	cmd.SetupRouter(r, selectedPath)
	r.Run(":8080")
}
