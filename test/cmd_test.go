package test

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestPath(t *testing.T) {
	passwd, _ := bcrypt.GenerateFromPassword([]byte("1234567"), bcrypt.DefaultCost)
	pawd := string(passwd)
	println(pawd)
	err := bcrypt.CompareHashAndPassword([]byte("$2a$10$5y7UCc9jKRejXyKcXNI3cuRURMZgatUyok/jTYVhxQbNuJL9jhLvm"), []byte("1234567"))
	println(err == nil)
}

type FileOperations uint

const (
	R FileOperations = 1 << iota // 1
	W                            // 2
	D                            // 4
)

func TestSS(t *testing.T) {
	p := R | W
	if p&(R|W) == (R | W) {
		fmt.Println("同时拥有读取和写入权限")
	} else {
		fmt.Println("没有同时拥有读取和写入权限")
	}
}
