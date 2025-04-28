package service

import (
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/request"
	"my-ops-admin/response"
)

func GetDictPagenation(dpi request.DictPagenationInfo) (dictPageRes response.DictPagenationRes, err error) {
	db := global.OPS_DB.Model(&models.SysDict{})
	limit := dpi.PageSize
	offset := dpi.PageSize * (dpi.PageNum - 1)
	if dpi.Keywords != "" {
		db = db.Or("name LIKE ?", "%"+dpi.Keywords+"%")
	}
	err = db.Count(&dictPageRes.Total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&dictPageRes.List).Error
	return
}
