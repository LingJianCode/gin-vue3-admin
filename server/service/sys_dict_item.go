package service

import (
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/request"
	"my-ops-admin/response"
)

var DictItemServiceApp = new(SysDictItemService)

type SysDictItemService struct{}

func (a *SysDictItemService) GetDictItemList(dictCode string) (dtList []*response.DictItem, err error) {
	err = global.OPS_DB.Model(&models.SysDictItem{}).Where("dict_code = ?", dictCode).Find(&dtList).Error
	if err != nil {
		return
	}
	return
}
func (a *SysDictItemService) GetDictItemListPagination(dictCode string, dipi request.DictItemPaginationInfo) (diRes response.DictItemPaginationRes, err error) {
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

func (a *SysDictItemService) GetDictItemForm(dictCode string, itemId uint) (item models.SysDictItem, err error) {
	err = global.OPS_DB.Model(&models.SysDictItem{}).Where("dict_code = ?", dictCode).First(&item, itemId).Error
	if err != nil {
		return
	}
	return
}
