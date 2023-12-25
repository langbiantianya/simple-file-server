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
	web.InitRoute(r, b)
	webdav.InitWebDav(r, b, c)
	return r
}

// BasicAuth 是一个中间件函数，用于进行基本身份验证
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{})
}
