package plugins

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
}

func (acctx *AccountCtx) InitRoot(username, passwd string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		log.Default().Fatalln(err)
	}
	account := Account{
		Username: username,
		Password: string(hash),
		Identity: Root,
	}
	acctx.Db.Create(&account)
}
