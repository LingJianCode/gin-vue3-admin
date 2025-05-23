package mycasbin

import (
	"fmt"
	"my-ops-admin/global"
	"strings"
	"sync"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	once   sync.Once
	Casbin = new(casbinHandler)
)

// 定义常量提高代码可读性
const (
	PolicyTypePolicy   = "p"
	PolicyTypeGrouping = "g"
)

type casbinHandler struct {
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	db                   *gorm.DB
	logger               *zap.Logger
}

func (c *casbinHandler) InitCasbin() {
	once.Do(func() {
		c.db = global.OPS_DB
		c.logger = global.OPS_LOGGER

		a, err := gormadapter.NewAdapterByDB(c.db)
		if err != nil {
			panic(fmt.Sprintln("casbin数据库初始化失败!", err.Error()))
		}

		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = g(r.sub, p.sub) && keyMatch2(r.obj,p.obj) && r.act == p.act || isSuperAdmin(r.sub)
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			panic(fmt.Sprintln("字符串加载模型失败!", err.Error()))
		}
		c.syncedCachedEnforcer, err = casbin.NewSyncedCachedEnforcer(m, a)
		if err != nil {
			panic(fmt.Sprintln("NewSyncedCachedEnforcer失败!", err.Error()))
		}
		c.syncedCachedEnforcer.AddFunction("isSuperAdmin", func(arguments ...interface{}) (interface{}, error) {
			// 获取用户名，由于使用了g导致传入参数变成了角色
			arg_0 := arguments[0].(string)
			if strings.HasPrefix(arg_0, "role_") {
				// 检查角色是否为超级管理员
				return arg_0 == c.MakeRoleName(global.SUPER_ADMIN_ROLE_ID), nil
			} else {
				// 检查用户名的角色是否为超级管理员
				return c.syncedCachedEnforcer.HasRoleForUser(arg_0, c.MakeRoleName(global.SUPER_ADMIN_ROLE_ID))
			}

		})
		c.syncedCachedEnforcer.SetExpireTime(60 * 60)
		err = c.syncedCachedEnforcer.LoadPolicy()
		if err != nil {
			panic(fmt.Sprintln("LoadPolicy失败!", err.Error()))
		}
	})
}

// Enforce 校验权限
func (c *casbinHandler) Enforce(user, uri, action string) (bool, error) {
	result, err := c.syncedCachedEnforcer.Enforce(user, uri, action)
	if err != nil {
		c.logger.Error("Enforce 执行失败!", zap.String("user", user), zap.String("uri", uri), zap.String("action", action), zap.Error(err))
	}
	return result, err
}

func (c *casbinHandler) EnforceRole(roleId uint, uri, action string) (bool, error) {
	result, err := c.syncedCachedEnforcer.Enforce(c.MakeRoleName(roleId), uri, action)
	if err != nil {
		c.logger.Error("Enforce 执行失败!", zap.String("roleId", c.MakeRoleName(roleId)), zap.String("uri", uri), zap.String("action", action), zap.Error(err))
	}
	return result, err
}

// AddPolicy 添加策略
func (c *casbinHandler) AddPolicy(roleId uint, uri, method string) (bool, error) {
	result, err := c.syncedCachedEnforcer.AddPolicy(c.MakeRoleName(roleId), uri, method)
	if err == nil && result {
		c.syncedCachedEnforcer.LoadPolicy() //清楚缓存并重新加载策略
	}
	if err != nil {
		c.logger.Error("AddPolicy 执行失败!", zap.Uint("roleId", roleId), zap.String("uri", uri), zap.String("method", method), zap.Error(err))
	}
	return result, err
}

// 拼接角色ID，为了防止角色与用户名冲突
func (c *casbinHandler) MakeRoleName(roleId uint) string {
	return fmt.Sprintf("role_%d", roleId)
}

// AddPolicies 批量添加策略
func (c *casbinHandler) AddPolicies(rules [][]string) (bool, error) {
	result, err := c.syncedCachedEnforcer.AddPolicies(rules)
	if err == nil && result {
		c.syncedCachedEnforcer.LoadPolicy() //清楚缓存并重新加载策略
	}
	if err != nil {
		c.logger.Error("AddPolicies 执行失败!", zap.Any("rules", rules), zap.Error(err))
	}
	return result, err
}

// 批量添加角色的策略
func (c *casbinHandler) AddRolePolicies(roleId uint, apiUriList, apiMethodList []string) (bool, error) {
	rules := make([][]string, 0)
	for idx, uri := range apiUriList {
		rules = append(rules, []string{c.MakeRoleName(roleId), uri, apiMethodList[idx]})
	}
	result, err := c.AddPolicies(rules)
	if err == nil && result {
		c.syncedCachedEnforcer.LoadPolicy() //清楚缓存并重新加载策略
	}
	if err != nil {
		c.logger.Error("AddRolePolicies 执行失败!", zap.Any("roleId", roleId), zap.Error(err))
	}
	return result, err
}

// DeleteRole 删除角色对应的用户和权限
func (c *casbinHandler) DeleteRole(roleId uint) (bool, error) {
	result, err := c.syncedCachedEnforcer.DeleteRole(c.MakeRoleName(roleId))
	if err == nil && result {
		c.syncedCachedEnforcer.LoadPolicy() //清楚缓存并重新加载策略
	}
	if err != nil {
		c.logger.Error("DeleteRole 执行失败!", zap.Uint("roleId", roleId), zap.Error(err))
	}
	return result, err
}

// DeleteRolePolicy 删除角色下的权限
func (c *casbinHandler) DeleteRolePolicy(roleId uint) (bool, error) {
	result, err := c.syncedCachedEnforcer.RemoveFilteredNamedPolicy(PolicyTypePolicy, 0, c.MakeRoleName(roleId))
	if err == nil && result {
		c.syncedCachedEnforcer.LoadPolicy() //清楚缓存并重新加载策略
	}
	if err != nil {
		c.logger.Error("DeleteRolePolicy 执行失败!", zap.Uint("roleId", roleId), zap.Error(err))
	}
	return result, err
}

// DeleteRoleUser 删除添加用户
func (c *casbinHandler) DeleteRoleUser(roleId uint) (bool, error) {
	result, err := c.syncedCachedEnforcer.RemoveFilteredNamedGroupingPolicy(PolicyTypeGrouping, 1, c.MakeRoleName(roleId))
	if err == nil && result {
		c.syncedCachedEnforcer.LoadPolicy() //清楚缓存并重新加载策略
	}
	if err != nil {
		c.logger.Error("DeleteRoleUser 执行失败!", zap.Uint("roleId", roleId), zap.Error(err))
	}
	return result, err
}

// AddUserRole 添加角色和用户对应关系
func (c *casbinHandler) AddUserRole(user string, roleId uint) (bool, error) {
	result, err := c.syncedCachedEnforcer.AddGroupingPolicy(user, c.MakeRoleName(roleId))
	if err == nil && result {
		c.syncedCachedEnforcer.LoadPolicy() //清楚缓存并重新加载策略
	}
	if err != nil {
		c.logger.Error("AddUserRole 执行失败!", zap.String("user", user), zap.Uint("roleId", roleId), zap.Error(err))
	}
	return result, err
}

// AddUserRoles 批量添加角色和用户对应关联
func (c *casbinHandler) AddUserRoles(usernames []string, roleIds []uint) (bool, error) {
	rules := make([][]string, 0)
	for _, u := range usernames {
		for _, r := range roleIds {
			rules = append(rules, []string{u, c.MakeRoleName(r)})
		}
	}
	result, err := c.syncedCachedEnforcer.AddGroupingPolicies(rules)
	if err == nil && result {
		c.syncedCachedEnforcer.LoadPolicy() //清楚缓存并重新加载策略
	}
	if err != nil {
		c.logger.Error("AddUserRoles 执行失败!", zap.Any("usernames", usernames), zap.Any("roleIds", roleIds), zap.Error(err))
	}
	return result, err
}

// DeleteUserRole 删除用户的角色信息
func (c *casbinHandler) DeleteUserRole(user string) (bool, error) {
	ok, err := c.syncedCachedEnforcer.HasNamedGroupingPolicy(PolicyTypeGrouping, user)
	if err != nil {
		c.logger.Error("DeleteUserRole HasNamedGroupingPolicy 执行失败!", zap.String("user", user), zap.Error(err))
		return false, err
	}
	if !ok {
		return true, nil
	}
	result, err := c.syncedCachedEnforcer.RemoveFilteredNamedGroupingPolicy(PolicyTypeGrouping, 0, user)
	if err == nil && result {
		c.syncedCachedEnforcer.LoadPolicy() //清楚缓存并重新加载策略
	}
	if err != nil {
		c.logger.Error("DeleteUserRole 执行失败!", zap.String("user", user), zap.Error(err))
	}
	return result, err
}
