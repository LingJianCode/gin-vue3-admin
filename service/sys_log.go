package service

import (
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/request"
	"my-ops-admin/response"
)

var SysLogServiceApp = new(SysLogService)

type SysLogService struct{}

func (a *SysLogService) CreateSysOperationLog(sysOperationLog models.SysLog) (err error) {
	err = global.OPS_DB.Create(&sysOperationLog).Error
	return err
}

func (a *SysLogService) GetDictPagination(lpi request.LogPaginationInfo) (logPageRes response.LogPaginationRes, err error) {
	db := global.OPS_DB.Model(&models.SysLog{})
	limit := lpi.PageSize
	offset := lpi.PageSize * (lpi.PageNum - 1)
	if lpi.UserId != 0 {
		db = db.Where("user_id = ?", lpi.UserId)
	}
	err = db.Count(&logPageRes.Total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&logPageRes.List).Error
	return
}
