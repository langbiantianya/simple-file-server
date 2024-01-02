package main

import (
	"simpleFileServer/cmd"
	"simpleFileServer/cmd/common"
	"simpleFileServer/cmd/plugins"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	plugins.InitWebStatic(r)
	cmd.SetupRouter(r, common.InitContext())
	r.Run(":8080")
}
