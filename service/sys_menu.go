package service

import (
	"errors"
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/response"
	"my-ops-admin/utils"
	"reflect"

	"gorm.io/gorm"
)

func CreateMenu(menu models.SysMenu) error {
	if !errors.Is(global.OPS_DB.Where("name = ?", menu.Name).First(&models.SysMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return global.OPS_DB.Create(&menu).Error
}

func GetMenuTree() (menus []*response.MenuTreeRes, err error) {
	var treeMenu []*response.MenuTreeRes
	err = global.OPS_DB.Model(&models.SysMenu{}).Order("sort").Preload("Params").Find(&treeMenu).Error
	if err != nil {
		return
	}
	menu := getChildrenList(treeMenu, 0)
	return menu, nil
}

func getChildrenList(menus []*response.MenuTreeRes, parentId uint) []*response.MenuTreeRes {
	// 定义子节点目录
	var nodes []*response.MenuTreeRes
	// 判断反射值的有效性
	if reflect.ValueOf(menus).IsValid() {
		// 循环所有菜单
		for _, v := range menus {
			// 操作指定级别菜单 parentId为0是表示一级
			if v.ParentID == parentId {
				// 将子级菜单 不定长压入 children数组
				v.Children = append(v.Children, getChildrenList(menus, v.ID)...)
				// 放入结果数组
				nodes = append(nodes, v)
			}
		}
	}
	return nodes
}

func GetMenuRouteTree() (menus []*response.MenuRouteRes, err error) {
	var menuList []*models.SysMenu
	err = global.OPS_DB.Where("type != ?", 4).Order("sort").Preload("Params").Find(&menuList).Error
	if err != nil {
		return
	}
	menu := buildMenuRoute(menuList, 0)
	return menu, nil
}

func buildMenuRoute(menuList []*models.SysMenu, parentId uint) []*response.MenuRouteRes {
	res := make([]*response.MenuRouteRes, 0)
	for _, v := range menuList {
		if v.ParentID == parentId {
			menuRoute := response.MenuRouteRes{
				Component: v.Component,
				Name:      v.RoutePath,
				Path:      v.RoutePath,
				Redirect:  v.Redirect,
				Meta: response.Meta{
					AlwaysShow: utils.Int64ToBool(v.AlwaysShow),
					Hidden:     utils.Int64ToBool(v.Visible),
					Icon:       v.Icon,
					KeepAlive:  utils.Int64ToBool(v.KeepAlive),
					Title:      v.Name,
					Params:     v.Params,
				},
			}
			menuRoute.Children = buildMenuRoute(menuList, v.ID)
			res = append(res, &menuRoute)
		}
	}
	return res
}
