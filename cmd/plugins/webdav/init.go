package webdav

import (
	"log"
	"net/http"
	"simpleFileServer/cmd/common"
	"simpleFileServer/cmd/plugins"
	"simpleFileServer/cmd/plugins/account"
	"simpleFileServer/cmd/plugins/rights"

	"github.com/gin-gonic/gin"
)

func InitWebDav(r *gin.Engine, ctx *common.ServerContext) {
	r.Use(serve("/webdav", ctx.WorkHome, func(c *gin.Context) bool {
		username, _, _ := plugins.ParseAuthHeader(c.Request.Header.Get("Authorization"))
		// TODO 👇要重构
		if ctx.MultipleUser {
			return plugins.BasicAuth(c, ctx) && account.VerifyRights(ctx.Acctx, username, func(operations rights.FileOperations) bool {
				return rights.Verify(c, operations)
			})
		} else {
			return plugins.BasicAuth(c, ctx)
		}

	}, func(req *http.Request, err error) {
		if err != nil {
			log.Default().Println(req.URL.Path, err)
		}
	}))
}
