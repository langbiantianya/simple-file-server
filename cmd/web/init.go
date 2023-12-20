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
	api.POST("/upload", func(ctx *gin.Context) {
		err := uploadr(ctx, b)
		if err != nil {
			log.Default().Println(err)
			ctx.JSON(400, vo.Error{Error: err.Error()})
		}
		ctx.JSON(200, vo.Success{Data: "ok"})
	})
}
