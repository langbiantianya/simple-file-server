package web

import (
	"log"
	"simpleFileServer/cmd/common"
	"simpleFileServer/cmd/plugins"
	"simpleFileServer/cmd/server"
	"simpleFileServer/cmd/vo"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine, ctx *common.ServerContext) {
	api := r.Group("/api")
	api.GET("/*paths", func(c *gin.Context) {
		err := listDir(c, server.Default(ctx.WorkHome))
		if err != nil {
			log.Default().Println(err)
			c.JSON(400, vo.Error{Error: err.Error()})
		}
	})
	api.POST("/*paths", func(c *gin.Context) {
		err := upload(c, server.Default(ctx.WorkHome))
		if err != nil {
			log.Default().Println(err)
			c.JSON(400, vo.Error{Error: err.Error()})
		}
	})
	api.DELETE("/*paths", plugins.Default(ctx), func(c *gin.Context) {
		err := remove(c, server.Default(ctx.WorkHome))
		if err != nil {
			log.Default().Println(err)
			c.JSON(400, vo.Error{
				Error: err.Error(),
			})
		}
	})
}
