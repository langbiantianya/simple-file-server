package utils

import "golang.org/x/crypto/bcrypt"

func VerifyPassword(hashPasswd []byte, password []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPasswd), []byte(password))
	return err == nil
}
