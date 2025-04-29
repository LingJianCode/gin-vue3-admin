package service

import (
	"my-ops-admin/global"
	"my-ops-admin/models"
)

func GetApiOptions() (api []models.SysApi, err error) {
	err = global.OPS_DB.Find(&api).Error
	return
}
