package common

import (
	"gorm.io/gorm"
)

type ServerContext struct {
	WorkHome     string
	RootUser     string
	Passwd       string
	MultipleUser bool
	Db           *gorm.DB
}
