package mysql

import (
	"fmt"

	"github.com/luckyss6/go-kit/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func New(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.Mysql.Username,
		cfg.Mysql.Password,
		cfg.Mysql.Host,
		cfg.Mysql.Port,
		cfg.Mysql.DBname)
	fmt.Printf("dsn: %v\n", dsn)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // 数据源名称
		DefaultStringSize:         256,   // 字符串字段的默认大小
		DisableDatetimePrecision:  true,  // 禁用日期时间精度，MySQL 5.6 之前不支持
		DontSupportRenameIndex:    true,  // 重命名索引时删除并创建，MySQL 5.7 之前不支持重命名索引，MariaDB
		DontSupportRenameColumn:   true,  // 重命名列时 `change`，MySQL 8 之前不支持重命名列，MariaDB
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("can't connect to database: %v", err))
	}
	return db
}
