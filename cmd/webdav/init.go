package webdav

import (
	"log"
	"net/http"
	"simpleFileServer/cmd/common"
	"simpleFileServer/cmd/plugins"

	"github.com/gin-gonic/gin"
)

func InitWebDav(r *gin.Engine, ctx *common.ServerContext) {
	r.Use(serve("/webdav", ctx.WorkHome, func(c *gin.Context) bool {
		return plugins.BasicAuth(c, ctx)
	}, func(req *http.Request, err error) {
		if err != nil {
			log.Default().Println(req.URL.Path, err)
		}
	}))
}
