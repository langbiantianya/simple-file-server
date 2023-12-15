package main

import "simpleFileServer/cmd/server"

func main() {
	root := server.SelectedPath{
		RootPath: "./",
		Parent:   "./",
		NowPath:  "./",
		Entries:  nil,
	}
	res := root.Ls().Mkdir("测试目录创建")
	for _, entry := range res.Entries {
		if entry.IsDir() {
			println("目录:", entry.Name())
		} else {
			println("文件:", entry.Name())
		}
	}

}
