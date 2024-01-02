package account

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	Model     Model  `gorm:"embedded"`
	Username  string `gorm:"unique"`
	Password  string
	Identity  Role
	Question1 string
	Answer1   string
	Question2 string
	Answer2   string
	Question3 string
	Answer3   string
	Two2FA    string
}

type Model struct {
	ID        uint `gorm:"primarykey;autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role int

const (
	Root Role = iota
	User
)

func (acctx *AccountCtx) add(account *Account) error {
	passwd, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	account.Password = string(passwd)
	return acctx.Db.Create(account).Error
}

func (acctx *AccountCtx) update(account *Account) error {
	return acctx.Db.Model(account).Updates(account).Error
}

func (acctx *AccountCtx) delete(account *Account) error {
	return acctx.Db.Delete(&Account{}, account.Model.ID).Error
}

func (acctx *AccountCtx) findOne(account *Account) *Account {
	var res Account
	acctx.Db.Where(&Account{}, account).First(&res)
	return &res
}

func (acctx *AccountCtx) list(account *Account) *[]Account {
	var res []Account
	acctx.Db.Where(&Account{}, account).Find(&res)
	return &res
}
