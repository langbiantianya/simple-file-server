package webdav

import (
	"log"
	"net/http"
	"simpleFileServer/cmd/server"

	"github.com/gin-gonic/gin"
)

func InitWebDav(r *gin.Engine, b *server.SelectedPath) {
	r.Use(serve("/webdav", b.RootPath, func(c *gin.Context) bool {
		return true
	}, func(req *http.Request, err error) {
		log.Default().Println(req.URL.Path, err)
	}))
}
