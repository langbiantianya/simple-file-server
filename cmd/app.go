package cmd

import (
	"simpleFileServer/cmd/common"
	"simpleFileServer/cmd/plugins/webdav"
	"simpleFileServer/cmd/web"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, c *common.ServerContext) *gin.Engine {
	web.InitRoute(r, c)
	webdav.InitWebDav(r, c)
	return r
}

// BasicAuth 是一个中间件函数，用于进行基本身份验证
func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{})
}
