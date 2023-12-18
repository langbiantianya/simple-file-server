package main

import (
	"io/fs"
	"simpleFileServer/cmd"
	"simpleFileServer/cmd/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	ctx := &server.SelectedPath{
		RootPath: "./",
		Parent:   "./",
		NowPath:  "./",
		Entries:  []fs.DirEntry{},
	}
	cmd.SetupRouter(r, ctx)
	r.Run(":8080")
}
