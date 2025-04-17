package global

import (
	"my-ops-admin/config"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	OPS_DB     *gorm.DB
	OPS_CONFIG config.Server
	OPS_VIPER  *viper.Viper
	OPS_LOG    *zap.Logger
)
