package plugins

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitWebStatic(r *gin.Engine) *gin.Engine {
	r.GET("/", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "text/html")
		ctx.File("./static/index.html")
		// ctx.Redirect(http.StatusMovedPermanently, "/static")
	})
	r.Static("/static", "./static")
	r.StaticFS("/assets", http.Dir("./static/assets"))
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	return r
}
