package database

import (
	"fmt"
	"net/url"

	"github.com/cde/model"
	"github.com/cde/util/logger"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init 初始化
func Init() *gorm.DB {
	// driver := viper.GetString("database.driver")
	host := viper.GetString("database.host")
	user := viper.GetString("database.user")
	port := viper.GetString("database.port")
	pass := viper.GetString("database.pass")
	dbname := viper.GetString("database.dbname")
	charset := viper.GetString("database.charset")
	loc := viper.GetString("database.loc")
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s", user, pass, host, port, dbname, charset, url.QueryEscape(loc))
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		logger.Debugf("链接数据库错误 %v", err)
	}
	logger.Debug("链接数据库成功")
	db.AutoMigrate(&model.Admin{}) // 自动创建 User 表
	//db.AutoMigrate(&model.Roles{})     // 自动创建 Roles 表
	//db.AutoMigrate(&model.Menu{})      // 自动创建 Menu 表
	//db.AutoMigrate(&model.RoleMenu{})  // 自动创建 RoleMenu 表
	//db.AutoMigrate(&model.Goods{})     // 自动创建 Goods 表
	//db.AutoMigrate(&model.RoleRules{}) // 自动创建 RoleRules 表
	//db.AutoMigrate(&model.Users{})     // 自动创建 Users 表
	//db.AutoMigrate(&model.Category{})  // 自动创建 Category 表
	// database.SingularTable(true)            // 支持单数创建数据表
	DB = db
	return db
}

// 获取db句柄
func GetDB() *gorm.DB {
	return DB
}
