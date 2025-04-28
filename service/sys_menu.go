package service

import (
	"errors"
	"fmt"
	"my-ops-admin/global"
	"my-ops-admin/models"
	"my-ops-admin/request"
	"my-ops-admin/response"
	"reflect"

	"gorm.io/gorm"
)

func CreateMenu(menu models.SysMenu) error {
	// 这里感觉写的有问题，应该所有错误都返回
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

func GetMenuRouteTree(userId uint) (menus []*response.MenuRouteRes, err error) {
	var user models.SysUser
	err = global.OPS_DB.Preload("Roles").First(&user, userId).Error
	if err != nil {
		return
	}
	var roleIds []uint
	var rootFlag bool = false
	for _, v := range user.Roles {
		if v.Code == "ROOT" {
			rootFlag = true
			break
		}
		roleIds = append(roleIds, v.ID)
	}
	var menuList []*models.SysMenu
	if rootFlag {
		err = global.OPS_DB.Where("type != ?", models.SYS_MENU_TYPE_BUTTON).Order("sort").Preload("Params").Find(&menuList).Error
	} else {
		err = global.OPS_DB.
			Distinct(`
				"sys_menu"."id",
				"sys_menu"."created_at",
				"sys_menu"."updated_at",
				"sys_menu"."deleted_at",
				"sys_menu"."always_show",
				"sys_menu"."component",
				"sys_menu"."icon",
				"sys_menu"."keep_alive",
				"sys_menu"."name",
				"sys_menu"."parent_id",
				"sys_menu"."perm",
				"sys_menu"."redirect",
				"sys_menu"."route_name",
				"sys_menu"."route_path",
				"sys_menu"."sort",
				"sys_menu"."type",
				"sys_menu"."visible"
				`).
			Joins(`JOIN "sys_role_menu" ON "sys_role_menu"."sys_menu_id" = "sys_menu"."id" AND "sys_role_menu"."sys_role_id" in ?`, roleIds).
			Where(`"sys_menu"."type" != ?`, models.SYS_MENU_TYPE_BUTTON).
			Order(`"sys_menu"."sort"`).
			Preload("Params").
			Find(&menuList).Error
	}
	if err != nil {
		return
	}

	menu := buildMenuRoute(menuList, 0)
	return menu, nil
}

// buildMenuRoute
//
//	@param []*models.SysMenu
//	@param uint
//	@return []*response.MenuRouteRes
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
					AlwaysShow: v.AlwaysShow == 1,
					Hidden:     v.Visible == 0,
					Icon:       v.Icon,
					KeepAlive:  v.KeepAlive == 1,
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

func GetMenuOptions(onlyParent bool) (menuOptions []*response.MenuOption, err error) {
	var menuList []*models.SysMenu
	if onlyParent {
		err = global.OPS_DB.Where("type in ?", []interface{}{models.SYS_MENU_TYPE_CATALOG, models.SYS_MENU_TYPE_MENU}).Order("sort").Find(&menuList).Error
	} else {
		err = global.OPS_DB.Order("sort").Find(&menuList).Error
	}
	if err != nil {
		return
	}
	menuOptions = buildMenuOptions(menuList, 0)
	return
}

func buildMenuOptions(menuList []*models.SysMenu, parentId uint) []*response.MenuOption {
	var nodes []*response.MenuOption
	if reflect.ValueOf(menuList).IsValid() {
		for _, v := range menuList {
			if v.ParentID == parentId {
				menuOption := response.MenuOption{
					Label: v.Name,
					Value: v.ID,
				}
				menuOption.Children = buildMenuOptions(menuList, v.ID)
				nodes = append(nodes, &menuOption)
			}
		}
	}
	return nodes
}

func UpdateMenu(id uint, menuReq request.MenuInfo) error {
	var menu models.SysMenu
	if errors.Is(global.OPS_DB.First(&menu, id).Error, gorm.ErrRecordNotFound) {
		return errors.New("记录不存在")
	}
	menu.AlwaysShow = menuReq.AlwaysShow
	menu.Component = menuReq.Component
	menu.Icon = menuReq.Icon
	menu.KeepAlive = menuReq.KeepAlive
	menu.Name = menuReq.Name
	menu.Params = menuReq.Params
	menu.Perm = menuReq.Perm
	menu.Redirect = menuReq.Redirect
	menu.RouteName = menuReq.RouteName
	menu.RoutePath = menuReq.RoutePath
	menu.Sort = menuReq.Sort
	menu.Visible = menuReq.Visible
	menu.ParentID = menuReq.ParentID
	return global.OPS_DB.Save(&menu).Error
}

// DeleteMenu
// 是否需要对菜单下的子节点进行递归删除？
//
//	@param uint
//	@return error
//	@author lingjian
//	@date 2025-04-28 11:02:41
func DeleteMenu(id uint) error {
	return global.OPS_DB.Delete(&models.SysMenu{}, id).Error
}

func GetMenuForm(id uint) (menu models.SysMenu, err error) {
	if errors.Is(global.OPS_DB.Preload("Params").First(&menu, id).Error, gorm.ErrRecordNotFound) {
		err = errors.New("记录不存在")
		return
	}
	fmt.Println("menu", menu.ParentID)
	return
}
