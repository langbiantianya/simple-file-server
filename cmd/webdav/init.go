package webdav

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"simpleFileServer/cmd/server"

	"github.com/gin-gonic/gin"
)

func InitWebDav(r *gin.Engine, b *server.SelectedPath) {
	r.Use(serve("/webdav", b.RootPath, func(c *gin.Context) bool {

		// 获取请求头中的 Authorization 字段
		authHeader := c.Request.Header.Get("Authorization")

		// 检查 Authorization 字段是否为空
		if authHeader == "" {
			// 如果为空，返回 401 Unauthorized
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return false
		}

		// 解析 Authorization 字段
		username, password, err := parseAuthHeader(authHeader)
		if err != nil {
			// 如果解析失败，返回 400 Bad Request
			c.AbortWithStatus(http.StatusBadRequest)
			return false
		}

		// 检查用户名和密码是否匹配
		if checkCredentials(username, password) {
			// 如果匹配，继续处理请求
			c.Next()
		} else {
			// 如果不匹配，返回 401 Unauthorized
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		return true
	}, func(req *http.Request, err error) {
		if err != nil {
			log.Default().Println(req.URL.Path, err)
		}
	}))
}

// 在这里定义一个 map，用于存储用户名和密码的键值对
var users = map[string]string{
	"username1": "password1",
	"username2": "password2",
}

// 解析 Authorization 字段，提取用户名和密码
func parseAuthHeader(authHeader string) (string, string, error) {
	const prefix = "Basic "

	// 检查 Authorization 字段的前缀是否为 "Basic "
	if len(authHeader) < len(prefix) || authHeader[:len(prefix)] != prefix {
		return "", "", fmt.Errorf("invalid auth header")
	}

	// 解码 Base64 编码的凭证部分
	credentials, err := b64Decode(authHeader[len(prefix):])
	if err != nil {
		return "", "", fmt.Errorf("invalid auth header")
	}

	// 将凭证部分拆分为用户名和密码
	pair := bytes.SplitN(credentials, []byte(":"), 2)
	if len(pair) != 2 {
		return "", "", fmt.Errorf("invalid auth header")
	}

	return string(pair[0]), string(pair[1]), nil
}

// Base64 解码
func b64Decode(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}

// 检查用户名和密码是否匹配
func checkCredentials(username, password string) bool {
	// 从存储的用户名和密码中查找匹配
	storedPassword, ok := users[username]
	if !ok {
		return false
	}

	// 检查密码是否匹配
	return password == storedPassword
}
