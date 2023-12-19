package web

import (
	"fmt"
	"log"
	"regexp"
	"simpleFileServer/cmd/server"
	"simpleFileServer/cmd/vo"

	"github.com/gin-gonic/gin"
)

func listDir(c *gin.Context, b *server.SelectedPath) error {
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
	if b.IsDir() {
		err := b.Ls(nil)
		if err != nil {
			return err
		}
		items := make([]vo.FileItem, 0)
		for _, d := range b.Entries {
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
		c.JSON(200, vo.Success{Data: dir})

	} else {
		c.File(b.NowPath)
	}
	return nil
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
