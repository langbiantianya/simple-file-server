package web

import (
	"log"
	"simpleFileServer/cmd/server"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine, b *server.SelectedPath) {
	r.GET("/testPath/*paths", func(ctx *gin.Context) {
		paths := ctx.Params
		p, s := paths.Get("paths")
		ctx.JSON(200, gin.H{
			"paths": p,
			"s":     s,
		})
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
