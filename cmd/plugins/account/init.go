package account

import (
	"simpleFileServer/cmd/common"
	"simpleFileServer/cmd/plugins/rights"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DefualtCtx(c *common.ServerContext) *AccountCtx {
	return &AccountCtx{
		Db: c.Db,
	}
}

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

func InitRoute(r *gin.Engine, ctx *common.ServerContext) {
	account := r.Group("/api/v1/account", func(c *gin.Context) {
		// 初步过滤未登入用户
	})

	// 获取用户列表
	account.GET("/list", func(c *gin.Context) {
		// 过滤非管理
	}, func(c *gin.Context) {})
	// 添加用户
	account.POST("/add", func(c *gin.Context) {
		// 过滤非管理
	}, func(ctx *gin.Context) {

	})
	// 修改权限
	account.PUT("/modifyPermissions", func(c *gin.Context) {

	})
	// 修改角色
	account.PUT("/modifyRole", func(c *gin.Context) {

	})
	// 修改密码
	account.PUT("/changePassword", func(c *gin.Context) {

	})
	// 忘记密码
	account.POST("/forgotPassword", func(c *gin.Context) {

	})
}
