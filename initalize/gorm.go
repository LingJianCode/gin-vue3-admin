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
		models.SysUser{},
		models.SysDepartment{},
	)
	if err != nil {
		global.OPS_LOGGER.Error("AutoMigrate table failed", zap.Error(err))
		os.Exit(0)
	}
	global.OPS_LOGGER.Info("AutoMigrate table success")
}
