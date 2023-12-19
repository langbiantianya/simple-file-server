package main

import (
	"simpleFileServer/cmd"
	"simpleFileServer/cmd/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	ctx := server.Default("./")
	cmd.SetupRouter(r, ctx)
	r.Run(":8080")
}
