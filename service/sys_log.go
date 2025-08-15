package service

import (
	"my-ops-admin/global"
	"my-ops-admin/models"
)

var SysLogServiceApp = new(SysLogService)

type SysLogService struct{}

func (a *SysLogService) CreateSysOperationLog(sysOperationLog models.SysLog) (err error) {
	err = global.OPS_DB.Create(&sysOperationLog).Error
	return err
}
