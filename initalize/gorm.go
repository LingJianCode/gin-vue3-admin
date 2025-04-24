package initalize

import (
	"fmt"
	"my-ops-admin/global"
	"my-ops-admin/models"
	"os"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() {
	/*
		注意 AutoMigrate 会自动创建数据库外键约束，您可以在初始化时禁用此功能，例如：
		db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		})
	*/
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  global.OPS_CONFIG.PostgreSQL.Dsn(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println(err)
	}
	global.OPS_DB = db
}

func AutoMigrate() {
	db := global.OPS_DB
	err := db.AutoMigrate(

		models.SysUserRole{},
		models.SysRole{},
		models.SysUser{},
		models.SysDept{},
		models.SysMenu{},
		models.SysMenuParameter{},
	)
	if err != nil {
		global.OPS_LOGGER.Error("AutoMigrate table failed", zap.Error(err))
		os.Exit(0)
	}
	global.OPS_LOGGER.Info("AutoMigrate table success")
}
