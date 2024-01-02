package common

import (
	"os"
	"simpleFileServer/cmd/plugins"
	account "simpleFileServer/cmd/plugins/account"
	"strconv"
)

func InitContext() *ServerContext {
	MultipleUser, err := strconv.ParseBool("MULTIPLE_USERS")
	if err != nil {
		MultipleUser = false
	}

	ctx := &ServerContext{
		WorkHome:     os.Getenv("WORK_HOME"),
		RootUser:     os.Getenv("ROOT_USER"),
		Passwd:       os.Getenv("PASSWD"),
		MultipleUser: MultipleUser,
	}
	if ctx.WorkHome == "" {
		ctx.WorkHome = "./"
	}
	if ctx.Passwd == "" {
		ctx.Passwd = "123456"
	}
	if ctx.RootUser == "" {
		ctx.RootUser = "root"
	}
	if MultipleUser {
		ctx.Acctx = &account.AccountCtx{
			Db: plugins.InitSqlite(),
		}
	}
	return ctx
}
