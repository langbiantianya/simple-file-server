package account

import (
	"log"

	"golang.org/x/crypto/bcrypt"
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
		hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
		if err != nil {
			log.Default().Fatalln(err)
		}
		account := &Account{
			Username: username,
			Password: string(hash),
			Identity: Root,
		}
		acctx.add(account)
		acctx.addMatedata()
	}
}
