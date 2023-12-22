package cmd

import (
	"simpleFileServer/cmd/common"
	"simpleFileServer/cmd/server"
	"simpleFileServer/cmd/web"
	"simpleFileServer/cmd/webdav"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, c *common.ServerContext) *gin.Engine {

	b := server.Default(c.WorkHome, c.Passwd)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "")
	})
	web.InitRoute(r, b)
	webdav.InitWebDav(r, b)
	return r
}
