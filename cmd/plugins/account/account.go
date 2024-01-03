package account

import (
	"log"
)

func VerifyPassword(acctx *AccountCtx, username string, Verify func(hashPasswd string) bool) bool {
	res := acctx.findOne(&Account{Username: username})
	return Verify(res.Password)
}

func Add(acctx *AccountCtx, account *Account) {
	err := acctx.add(account)
	if err != nil {
		log.Default().Println(err)
	}
}
func Update(acctx *AccountCtx, account *Account) {
	err := acctx.update(account)
	if err != nil {
		log.Default().Println(err)
	}
}
func Delete(acctx *AccountCtx, account *Account) {
	err := acctx.delete(account)
	if err != nil {
		log.Default().Println(err)
	}
}
func FindOne(acctx *AccountCtx, account *Account) *Account {
	return acctx.findOne(account)
}
func List(acctx *AccountCtx, account *Account) *[]Account {
	return acctx.list(account)
}
