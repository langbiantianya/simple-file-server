package webdav

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/webdav"
)

func Serve(prefix string, rootDir string,
	validator func(c *gin.Context) bool,
	logger func(req *http.Request, err error)) gin.HandlerFunc {
	w := webdav.Handler{
		Prefix:     prefix,
		FileSystem: webdav.Dir(rootDir),
		LockSystem: webdav.NewMemLS(),
		Logger:     logger,
	}
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, w.Prefix) {
			if validator != nil && !validator(c) {
				c.AbortWithStatus(403)
				return
			}
			c.Status(200) // 200 by default, which may be override later
			w.ServeHTTP(c.Writer, c.Request)
			c.Abort()
		}
	}
}
