package web

import (
	"log"
	"simpleFileServer/cmd/server"

	"github.com/gin-gonic/gin"
)

func uploadr(c *gin.Context, b *server.SelectedPath) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(400, "未找到文件")
		return
	}
	srcFile, err := file.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()
	b.Touch(file.Filename, srcFile)
}
