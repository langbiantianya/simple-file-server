package web

import (
	"fmt"
	"io/fs"
	"regexp"
	"simpleFileServer/cmd/server"

	"github.com/gin-gonic/gin"
)

func listDir(c *gin.Context, b *server.SelectedPath) ([]fs.DirEntry, error) {
	paths := c.Params
	path, _ := paths.Get("paths")
	path = path[1:]
	b.NowPath = fmt.Sprintf("%s%s", b.RootPath, path)
	pattern := "^(.*/)[^/]+/?$"
	regex := regexp.MustCompile(pattern)
	matches := regex.FindStringSubmatch(b.NowPath)
	if len(matches) > 1 {
		b.Parent = matches[1]
	}
	err := b.Ls(nil)
	return b.Entries, err
}
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
