package rights

import (
	"github.com/gin-gonic/gin"
)

type FileOperations uint

const (
	R FileOperations = 1 << iota // 1
	W                            // 2
	D                            // 4
)

// 读
func Read(c *gin.Context) FileOperations {
	switch c.Request.Method {
	case "GET", "PROPFIND", "OPTIONS":
		return R
	default:
		return 0
	}
}

// 写
func Write(c *gin.Context) FileOperations {
	switch c.Request.Method {
	case "PUT", "POST", "PROPPATCH", "MOVE", "MKCOL", "COPY":
		return W
	default:
		return 0
	}
}

// 删
func Delete(c *gin.Context) FileOperations {
	switch c.Request.Method {
	case "DELETE":
		return D
	default:
		return 0
	}
}
