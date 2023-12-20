package web

import (
	"fmt"
	"log"
	"mime/multipart"
	"regexp"
	"simpleFileServer/cmd/server"
	"simpleFileServer/cmd/vo"

	"github.com/gin-gonic/gin"
)

func flushed(c *gin.Context, b *server.SelectedPath) error {
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
	return b.Ls(nil)
}

func ls(b *server.SelectedPath) vo.Dir {
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
	return vo.Dir{
		RootPath:  b.RootPath,
		Parent:    b.Parent,
		NowPath:   b.NowPath,
		FileItems: items,
	}
}

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
		c.JSON(200, vo.Success{Data: ls(b)})

	} else {
		c.File(b.NowPath)
	}
	return nil
}

func upload(c *gin.Context, b *server.SelectedPath) error {

	file, err := c.FormFile("file")
	if file != nil && err == nil {
		touch(c, b, file)
	}

	path, flg := c.GetPostForm("path")
	if flg {
		err = mkdir(c, b, path)
	}
	return err

}

func mkdir(c *gin.Context, b *server.SelectedPath, path string) error {
	flushed(c, b)
	return b.Mkdir(path)
}

func touch(c *gin.Context, b *server.SelectedPath, file *multipart.FileHeader) error {
	flushed(c, b)
	srcFile, err := file.Open()
	if err != nil {
		return err
	}
	defer srcFile.Close()
	return b.Touch(file.Filename, srcFile)
}
