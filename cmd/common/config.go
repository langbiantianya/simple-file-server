package common

import (
	account "simpleFileServer/cmd/plugins/account"
)

type ServerContext struct {
	WorkHome string
	RootUser string
	Passwd   string
	Acctx    *account.AccountCtx
}
