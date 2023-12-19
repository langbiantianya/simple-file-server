package web

import (
	"log"
	"simpleFileServer/cmd/server"
	"simpleFileServer/cmd/vo"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine, b *server.SelectedPath) {
	r.GET("/testPath/*paths", func(ctx *gin.Context) {
		dirs, err := listDir(ctx, b)
		if err != nil {
			log.Default().Fatalln(err)
			ctx.String(400, err.Error())
		}

		items := make([]vo.FileItem, 0)
		for _, d := range dirs {
			info, err := d.Info()
			if err != nil {
				log.Default().Fatalln(err)
				continue
			}
			items = append(items, vo.FileItem{
				Name:    info.Name(),
				Size:    info.Size(),
				Mode:    uint32(info.Mode()),
				ModTime: info.ModTime(),
				IsDir:   info.IsDir(),
			})
		}
		dir := vo.Dir{
			RootPath:  b.RootPath,
			Parent:    b.Parent,
			NowPath:   b.NowPath,
			FileItems: items,
		}
		ctx.JSON(200, dir)
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
