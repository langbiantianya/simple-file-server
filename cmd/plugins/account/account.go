package account

import (
	"fmt"
	"simpleFileServer/cmd/plugins/rights"
	"simpleFileServer/cmd/utils"

	"golang.org/x/crypto/bcrypt"
)

func VerifyRights(acctx *AccountCtx, username string, verify func(operations rights.FileOperations) bool) bool {
	res := acctx.findOne(&Account{Username: username})
	return verify(res.Rights)
}

func DefaultVerifyPassword(acctx *AccountCtx, username string, password string) bool {
	return VerifyPassword(acctx, username, func(hashPasswd string) bool {
		return utils.VerifyPassword([]byte(hashPasswd), []byte(password))
	})
}
func VerifyPassword(acctx *AccountCtx, username string, verify func(hashPasswd string) bool) bool {
	res := acctx.findOne(&Account{Username: username})
	return verify(res.Password)
}

//	func Add(acctx *AccountCtx, account *Account) {
//		err := acctx.add(account)
//		if err != nil {
//			log.Default().Println(err)
//		}
//	}
//
//	func Update(acctx *AccountCtx, account *Account) {
//		err := acctx.update(account)
//		if err != nil {
//			log.Default().Println(err)
//		}
//	}
//
//	func Delete(acctx *AccountCtx, account *Account) {
//		err := acctx.delete(account)
//		if err != nil {
//			log.Default().Println(err)
//		}
//	}
func FindOne(acctx *AccountCtx, account *Account) *Account {
	return acctx.findOne(account)
}
func List(acctx *AccountCtx, account *Account) *[]Account {
	return acctx.list(account)
}

// 管理员添加用户
func AccountAdd(ctx *AccountCtx, account Account) error {
	return ctx.add(&account)
}

// 管理员更新用户角色
func UpdateIdentity(ctx *AccountCtx, account Account) error {
	a := &Account{
		Model:    Model{ID: account.Model.ID},
		Username: account.Username,
		Identity: account.Identity,
	}
	return ctx.update(a)
}

// 管理员更新用户读写删权限
func UpdateRights(ctx *AccountCtx, account Account) error {
	a := &Account{
		Model:    Model{ID: account.Model.ID},
		Username: account.Username,
		Rights:   account.Rights,
	}
	return ctx.update(a)
}

// 自己修改密码
func ChangePasswd(ctx *AccountCtx, account Account) error {
	DefaultVerifyPassword(ctx, account.Username, account.Password)
	return updatePasswd(ctx, account)
}

// 忘记密码
func ResetPassword(ctx *AccountCtx, account Account) error {
	vf := false
	res := ctx.findOne(&Account{
		Username: account.Username,
	})
	if res == nil {
		return fmt.Errorf("用户不存在")
	}
	if !vf && account.Answer1 != "" {
		vf = res.Answer1 == account.Answer1
	}
	if !vf && account.Answer2 != "" {
		vf = res.Answer2 == account.Answer2
	}
	if !vf && account.Answer3 != "" {
		vf = res.Answer3 == account.Answer3
	}
	if !vf {
		return fmt.Errorf("全部回答错误请重试")
	}
	return updatePasswd(ctx, account)
}

func updatePasswd(ctx *AccountCtx, account Account) error {
	passwd, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a := &Account{
		Model:    Model{ID: account.Model.ID},
		Username: account.Username,
		Password: string(passwd),
	}
	return ctx.update(a)
}
