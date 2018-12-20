package table

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/lvxiaorun/logger"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:123456@/landlord?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		logger.Error(err.Error())
		return
	}
	DB = db
	DB.LogMode(true)
}
