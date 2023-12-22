package main

import (
	"os"
	"simpleFileServer/cmd"
	"simpleFileServer/cmd/common"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	ctx := &common.ServerContext{
		WorkHome: os.Getenv("WORK_HOME"),
		Passwd:   os.Getenv("PASSWD"),
	}
	if ctx.WorkHome == "" {
		ctx.WorkHome = "./"
	}
	if ctx.Passwd == "" {
		ctx.Passwd = "123456"
	}
	cmd.SetupRouter(r, ctx)
	r.Run(":8080")
}
