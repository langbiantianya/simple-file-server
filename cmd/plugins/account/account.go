package plugins

import (
	"golang.org/x/crypto/bcrypt"
)

func (acctx *AccountCtx) VerifyPassword(account *Account) bool {
	res := acctx.FindOne(&Account{Username: account.Username})
	err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(account.Password))
	if err != nil {
		return false
	} else {
		return true
	}
}
