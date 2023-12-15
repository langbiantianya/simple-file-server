package main

import "os"

func main() {
	rootPath := "./"
	dir, err := os.Open(rootPath)
	if err != nil {
		println(err)
	}
	defer dir.Close()
	entries, err := dir.ReadDir(0)
	if err != nil {
		println(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			println("目录:", entry.Name())
		} else {
			println("文件:", entry.Name())
		}
	}
	// r := cmd.SetupRouter()
	// // Listen and Server in 0.0.0.0:8080
	// err := r.Run(":8080")
	// if err != nil {
	// 	return
	// }
}
