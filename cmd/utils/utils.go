package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/gin-gonic/gin"
)

func ParseBasicAuthHeader(c *gin.Context) (string, string, error) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		return "", "", fmt.Errorf("解析用户失败")
	}
	username, password, err := parseAuthHeader(authHeader)
	if err != nil {
		return "", "", fmt.Errorf("解析用户失败")
	}
	return username, password, err

}

// 解析 Authorization 字段，提取用户名和密码
func parseAuthHeader(authHeader string) (string, string, error) {
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
