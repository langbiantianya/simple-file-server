package plugins

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitWebStatic(r *gin.Engine) *gin.Engine {
	r.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/static")
	})
	r.Static("/static", "./static/dist")
	r.StaticFS("/assets", http.Dir("./static/dist/assets"))
	r.StaticFile("/favicon.ico", "./static/dist/favicon.ico")
	return r
}
