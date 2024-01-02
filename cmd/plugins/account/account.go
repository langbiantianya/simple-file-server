package account

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(acctx *AccountCtx, account *Account) bool {
	res := acctx.FindOne(&Account{Username: account.Username})
	err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(account.Password))
	if err != nil {
		return false
	} else {
		return true
	}
}

func Add(acctx *AccountCtx, account *Account) {
	err := acctx.Add(account)
	if err != nil {
		log.Default().Println(err)
	}
}
func Update(acctx *AccountCtx, account *Account) {
	err := acctx.Update(account)
	if err != nil {
		log.Default().Println(err)
	}
}
func Delete(acctx *AccountCtx, account *Account) {
	err := acctx.Delete(account)
	if err != nil {
		log.Default().Println(err)
	}
}
func FindOne(acctx *AccountCtx, account *Account) *Account {
	return acctx.FindOne(account)
}
func List(acctx *AccountCtx, account *Account) *[]Account {
	return acctx.List(account)
}
