package main

import (
	"os"
	"simpleFileServer/cmd"
	"simpleFileServer/cmd/server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	workHome := os.Getenv("WORK_HOME")
	if workHome == "" {
		workHome = "/tmp/"
	}
	ctx := server.Default(workHome)
	cmd.SetupRouter(r, ctx)
	r.Run(":8080")
}
