package initalize

import (
	"my-ops-admin/global"
	"my-ops-admin/pkg/log"
)

func InitZap() {
	global.OPS_LOG = log.Zap()
}
