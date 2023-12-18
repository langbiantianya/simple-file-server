package web

import (
	"simpleFileServer/cmd/server"

	"github.com/gin-gonic/gin"
)

func uploadr(c *gin.Context, b *server.SelectedPath) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	srcFile, err := file.Open()
	if err != nil {
		return err
	}
	defer srcFile.Close()
	return b.Touch(file.Filename, srcFile)
}
