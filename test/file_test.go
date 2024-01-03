package test

import (
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
