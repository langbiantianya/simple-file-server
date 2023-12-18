package cmd

import (
	"simpleFileServer/cmd/server"
	"simpleFileServer/cmd/web"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, b *server.SelectedPath) *gin.Engine {
	web.InitRoute(r, b)
	return r
}
