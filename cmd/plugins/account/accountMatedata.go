package account

// 用户表的元数据
type AccountMatedata struct {
	Model       Model `gorm:"embedded"`
	Initialized bool
}

func (acctx *AccountCtx) addMatedata() {
	acctx.Db.Create(&AccountMatedata{
		Initialized: true,
	})
}

func (acctx *AccountCtx) findMatedata() *AccountMatedata {
	var res AccountMatedata
	acctx.Db.First(&res)
	return &res
}
