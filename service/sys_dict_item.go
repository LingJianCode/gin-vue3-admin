package service

import (
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/request"
	"my-ops-admin/response"
)

func GetDictItem(dictCode string) (dtList []*response.DictItem, err error) {
	err = global.OPS_DB.Model(&models.SysDictItem{}).Where("dict_code = ?", dictCode).Find(&dtList).Error
	if err != nil {
		return
	}
	return
}
func GetDictItemPagenation(dictCode string, dipi request.DictItemPagenationInfo) (diRes response.DictItemPageRes, err error) {
	db := global.OPS_DB.Model(&models.SysDictItem{})
	limit := dipi.PageSize
	offset := dipi.PageSize * (dipi.PageNum - 1)
	if dipi.Keywords != "" {
		db = db.Or("label LIKE ?", "%"+dipi.Keywords+"%")
		db = db.Or("value LIKE ?", "%"+dipi.Keywords+"%")
	}
	err = db.Count(&diRes.Total).Error
	if err != nil {
		return
	}
	err = db.Where("dict_code = ?", dictCode).Limit(limit).Offset(offset).Find(&diRes.List).Error
	return
}
