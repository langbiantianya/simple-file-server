package cmd

import (
	"simpleFileServer/cmd/server"
	"simpleFileServer/cmd/web"
	"simpleFileServer/cmd/webdav"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, b *server.SelectedPath) *gin.Engine {
	r.GET("/", func(c *gin.Context) {
		c.String(200, "")
	})
	web.InitRoute(r, b)
	webdav.InitWebDav(r, b)
	return r
}
