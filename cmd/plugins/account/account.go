package account

import (
	"log"
	"simpleFileServer/cmd/plugins/rights"
)

func VerifyRights(acctx *AccountCtx, username string, verify func(operations rights.FileOperations) bool) bool {
	res := acctx.findOne(&Account{Username: username})
	return verify(res.Rights)
}
func VerifyPassword(acctx *AccountCtx, username string, verify func(hashPasswd string) bool) bool {
	res := acctx.findOne(&Account{Username: username})
	return verify(res.Password)
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
