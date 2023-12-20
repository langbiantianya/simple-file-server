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
	passwd := os.Getenv("PASSWD")
	if workHome == "" {
		workHome = "./"
	}
	if passwd == "" {
		passwd = "123456"
	}
	ctx := server.Default(workHome, passwd)
	cmd.SetupRouter(r, ctx)
	r.Run(":8080")
}
