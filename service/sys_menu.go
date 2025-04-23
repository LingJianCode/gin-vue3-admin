package service

import (
	"errors"
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/response"

	"gorm.io/gorm"
)

func CreateMenu(menu models.SysMenu) error {
	if !errors.Is(global.OPS_DB.Where("name = ?", menu.Name).First(&models.SysMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name，请修改name")
	}
	return global.OPS_DB.Create(&menu).Error
}

func GetMenusList() (menus []response.GetMenuRes, err error) {
	treeMap, err := getMenuTreeMap()
	if err != nil {
		return nil, err
	}
	menus = treeMap[0]
	for i := 0; i < len(menus); i++ {
		err = getChildrenList(&menus[i], treeMap)
	}
	return menus, err
}

func getMenuTreeMap() (treeMap map[uint][]response.GetMenuRes, err error) {
	var allMenus []models.SysMenu
	treeMap = make(map[uint][]response.GetMenuRes)
	err = global.OPS_DB.Order("sort").Preload("Params").Find(&allMenus).Error
	if err != nil {
		return
	}
	for _, v := range allMenus {
		treeMap[v.ParentID] = append(treeMap[v.ParentID], response.GetMenuRes{
			Component: v.Component,
			Icon:      v.Icon,
			ID:        v.ID,
			Name:      v.Name,
			ParentID:  v.ParentID,
			Perm:      v.Perm,
			Redirect:  v.Redirect,
			RouteName: v.RouteName,
			RoutePath: v.RoutePath,
			Sort:      v.Sort,
			Type:      v.Type,
			Visible:   v.Visible,
		})
	}
	return treeMap, err
}

func getChildrenList(menu *response.GetMenuRes, treeMap map[uint][]response.GetMenuRes) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}
