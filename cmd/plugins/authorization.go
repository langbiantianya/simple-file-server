package plugins

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"simpleFileServer/cmd/common"
	"simpleFileServer/cmd/plugins/account"

	"github.com/gin-gonic/gin"
)

func Default(ctx *common.ServerContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		BasicAuth(c, ctx)
	}
}

func BasicAuth(c *gin.Context, ctx *common.ServerContext) bool {
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
	username, password, err := ParseAuthHeader(authHeader)
	if err != nil {
		// 如果解析失败，返回 400 Bad Request
		c.AbortWithStatus(http.StatusBadRequest)
		return false
	}
	// TODO 重构👇
	// 检查用户名和密码是否匹配
	if checkBcryptCredentials(username, password, ctx) || checkCredentials(username, password, ctx) {
		// 如果匹配，继续处理请求
		c.Next()
	} else {
		// 如果不匹配，返回 401 Unauthorized
		c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	return true
}

// 解析 Authorization 字段，提取用户名和密码
func ParseAuthHeader(authHeader string) (string, string, error) {
	const prefix = "Basic "

	// 检查 Authorization 字段的前缀是否为 "Basic "
	if len(authHeader) < len(prefix) || authHeader[:len(prefix)] != prefix {
		return "", "", fmt.Errorf("invalid auth header")
	}

	// 解码 Base64 编码的凭证部分
	credentials, err := base64.StdEncoding.DecodeString(authHeader[len(prefix):])
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

// // Base64 解码
// func b64Decode(s string) ([]byte, error) {
// 	return base64.StdEncoding.DecodeString(s)
// }

// 检查用户名和密码是否匹配
func checkCredentials(username, password string, ctx *common.ServerContext) bool {
	// 从存储的用户名和密码中查找匹配

	if username != ctx.RootUser {
		return false
	}

	// 检查密码是否匹配
	return password == ctx.Passwd
}

func checkBcryptCredentials(username, password string, ctx *common.ServerContext) bool {
	if ctx.MultipleUser {
		return account.DefaultVerifyPassword(ctx.Acctx, username, password)
	} else {
		return false
	}

}
