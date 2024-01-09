package account

// 用户表的元数据
type AccountMatedata struct {
	Model       Model `gorm:"embedded"`
	Initialized bool
}

