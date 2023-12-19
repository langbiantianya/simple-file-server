package web

import (
	"log"
	"simpleFileServer/cmd/server"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine, b *server.SelectedPath) {
	r.GET("/testPath/*paths", func(ctx *gin.Context) {
		err := listDir(ctx, b)
		if err != nil {
			log.Default().Fatalln(err)
			ctx.String(400, err.Error())
		}
	})
	r.POST("/upload", func(ctx *gin.Context) {
		err := uploadr(ctx, b)
		if err != nil {
			log.Default().Fatalln(err)
			ctx.String(400, err.Error())
		}
		ctx.String(200, "ok")
	})
}
