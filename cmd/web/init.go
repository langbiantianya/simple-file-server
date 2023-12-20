package web

import (
	"log"
	"simpleFileServer/cmd/server"
	"simpleFileServer/cmd/vo"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine, b *server.SelectedPath) {
	api := r.Group("/api")
	api.GET("/*paths", func(ctx *gin.Context) {
		err := listDir(ctx, b)
		if err != nil {
			log.Default().Println(err)
			ctx.JSON(400, vo.Error{Error: err.Error()})
		}
	})
	api.POST("/*paths", func(ctx *gin.Context) {
		err := upload(ctx, b)
		if err != nil {
			log.Default().Println(err)
			ctx.JSON(400, vo.Error{Error: err.Error()})
		}
	})
	api.DELETE("/*paths", func(ctx *gin.Context) {

	})
}
