package web

import (
	"simpleFileServer/cmd/server"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine, b *server.SelectedPath) {
	r.POST("/upload", func(ctx *gin.Context) {
		uploadr(ctx, b)
	})
}
