package account

import (
	"simpleFileServer/cmd/plugins/rights"

	"gorm.io/gorm"
)

type AccountCtx struct {
	Db *gorm.DB
}

func (acctx *AccountCtx) InitDb() {
	acctx.Db.AutoMigrate(&Account{})
	acctx.Db.AutoMigrate(&AccountMatedata{})
}

func (acctx *AccountCtx) InitRoot(username, passwd string) {
	matedata := acctx.findMatedata()
	if matedata == nil || !matedata.Initialized {
		account := &Account{
			Username: username,
			Password: passwd,
			Identity: Root,
			Rights:   rights.R | rights.W | rights.D,
		}
		acctx.add(account)
		acctx.addMatedata()
	}
}
