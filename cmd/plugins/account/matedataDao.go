package account

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
