package initalize

import (
	"my-ops-admin/global"
	"my-ops-admin/pkg/logger"
)

func InitZap() {
	global.OPS_LOGGER = logger.Zap()
}
