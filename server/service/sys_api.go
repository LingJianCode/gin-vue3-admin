package service

import (
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/response"
)

var ApiServiceApp = new(SysApiService)

type SysApiService struct{}

func (a *SysApiService) GetApiOptions() (apiOptions []response.ApiOption, err error) {
	var apiList []models.SysApi
	err = global.OPS_DB.Order("id").Find(&apiList).Error
	apiMap := make(map[string][]models.SysApi)
	for _, v := range apiList {
		apiMap[v.Group] = append(apiMap[v.Group], v)
	}
	for key, val := range apiMap {
		api := response.ApiOption{
			Label: key,
			// value为0的不是api接口
			Value: 0,
		}
		for _, v := range val {
			api.Children = append(api.Children, &response.ApiOption{
				Label: v.Name,
				Value: v.ID,
			})
		}
		apiOptions = append(apiOptions, api)
	}
	return
}
