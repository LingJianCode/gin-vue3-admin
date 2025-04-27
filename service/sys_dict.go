package service

import (
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/response"
)

func GetDictItem(dictCode string) (dtList []*response.DictItem, err error) {
	err = global.OPS_DB.Model(&models.SysDictItem{}).Where("dict_code = ?", dictCode).Find(&dtList).Error
	if err != nil {
		return
	}
	return
}
