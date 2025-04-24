/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.50.37-pg
 Source Server Type    : PostgreSQL
 Source Server Version : 150012 (150012)
 Source Host           : 192.168.50.37:5432
 Source Catalog        : mypg
 Source Schema         : mypg

 Target Server Type    : PostgreSQL
 Target Server Version : 150012 (150012)
 File Encoding         : 65001

 Date: 24/04/2025 16:02:37
*/


-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS "mypg"."sys_menu";
CREATE TABLE "mypg"."sys_menu" (
  "id" int8 NOT NULL DEFAULT nextval('sys_menu_id_seq'::regclass),
  "created_at" timestamptz(6),
  "updated_at" timestamptz(6),
  "deleted_at" timestamptz(6),
  "always_show" int8,
  "component" text COLLATE "pg_catalog"."default",
  "icon" text COLLATE "pg_catalog"."default",
  "keep_alive" int8,
  "name" text COLLATE "pg_catalog"."default",
  "parent_id" int8,
  "perm" text COLLATE "pg_catalog"."default",
  "redirect" text COLLATE "pg_catalog"."default",
  "route_name" text COLLATE "pg_catalog"."default",
  "route_path" text COLLATE "pg_catalog"."default",
  "sort" int8,
  "type" int8,
  "visible" int8
)
;
COMMENT ON COLUMN "mypg"."sys_menu"."always_show" IS '【目录】只有一个子路由是否始终显示';
COMMENT ON COLUMN "mypg"."sys_menu"."component" IS '组件路径(vue页面完整路径，省略.vue后缀)';
COMMENT ON COLUMN "mypg"."sys_menu"."icon" IS '菜单图标';
COMMENT ON COLUMN "mypg"."sys_menu"."keep_alive" IS '【菜单】是否开启页面缓存';
COMMENT ON COLUMN "mypg"."sys_menu"."name" IS '菜单名称';
COMMENT ON COLUMN "mypg"."sys_menu"."parent_id" IS '父菜单ID';
COMMENT ON COLUMN "mypg"."sys_menu"."perm" IS '权限标识';
COMMENT ON COLUMN "mypg"."sys_menu"."redirect" IS '跳转路径';
COMMENT ON COLUMN "mypg"."sys_menu"."route_name" IS '路由名称';
COMMENT ON COLUMN "mypg"."sys_menu"."route_path" IS '路由路径';
COMMENT ON COLUMN "mypg"."sys_menu"."sort" IS '排序(数字越小排名越靠前)';
COMMENT ON COLUMN "mypg"."sys_menu"."type" IS '菜单类型（1-菜单 2-目录 3-外链 4-按钮）';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO "mypg"."sys_menu" VALUES (2, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'system/user/index', 'el-icon-User', 1, '用户管理', 1, NULL, NULL, 'User', 'user', 1, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (3, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'system/role/index', 'role', 1, '角色管理', 1, NULL, NULL, 'Role', 'role', 2, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (4, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'system/menu/index', 'menu', 1, '菜单管理', 1, NULL, NULL, 'SysMenu', 'menu', 3, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (6, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'system/dict/index', 'dict', 1, '字典管理', 1, NULL, NULL, 'Dict', 'dict', 5, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (20, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 1, 'Layout', 'cascader', NULL, '多级菜单', 0, NULL, '', NULL, '/multi-level', 9, 2, 1);
INSERT INTO "mypg"."sys_menu" VALUES (21, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 1, 'demo/multi-level/level1', '', NULL, '菜单一级', 20, NULL, '', NULL, 'multi-level1', 1, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (22, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, 'demo/multi-level/children/level2', '', NULL, '菜单二级', 21, NULL, NULL, NULL, 'multi-level2', 1, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (23, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, 'demo/multi-level/children/children/level3-1', '', 1, '菜单三级-1', 22, NULL, '', NULL, 'multi-level3-1', 1, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (26, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'Layout', 'document', NULL, '平台文档', 0, NULL, 'https://juejin.cn/post/7228990409909108793', '', '/doc', 8, 2, 1);
INSERT INTO "mypg"."sys_menu" VALUES (30, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, '', 'document', NULL, '平台文档(外链)', 26, NULL, '', NULL, 'https://juejin.cn/post/7228990409909108793', 2, 3, 1);
INSERT INTO "mypg"."sys_menu" VALUES (31, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '用户新增', 2, 'sys:user:add', '', NULL, '', 1, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (32, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '用户编辑', 2, 'sys:user:edit', '', NULL, '', 2, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (33, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '用户删除', 2, 'sys:user:delete', '', NULL, '', 3, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (36, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'Layout', 'menu', NULL, '组件封装', 0, NULL, '', NULL, '/component', 10, 2, 1);
INSERT INTO "mypg"."sys_menu" VALUES (37, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'demo/wang-editor', '', 1, '富文本编辑器', 36, NULL, '', NULL, 'wang-editor', 2, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (38, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'demo/upload', '', 1, '图片上传', 36, NULL, '', NULL, 'upload', 3, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (39, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'demo/icon-selector', '', 1, '图标选择器', 36, NULL, '', NULL, 'icon-selector', 4, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (40, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 1, 'Layout', 'api', NULL, '接口文档', 0, NULL, '', NULL, '/api', 7, 2, 1);
INSERT INTO "mypg"."sys_menu" VALUES (41, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'demo/api/apifox', 'api', 1, 'Apifox', 40, NULL, '', NULL, 'apifox', 1, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (71, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '角色编辑', 3, 'sys:role:edit', NULL, NULL, '', 3, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (72, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '角色删除', 3, 'sys:role:delete', NULL, NULL, '', 4, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (73, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '菜单新增', 4, 'sys:menu:add', NULL, NULL, '', 1, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (74, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '菜单编辑', 4, 'sys:menu:edit', NULL, NULL, '', 3, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (75, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '菜单删除', 4, 'sys:menu:delete', NULL, NULL, '', 3, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (76, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '部门新增', 5, 'sys:dept:add', NULL, NULL, '', 1, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (78, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '部门删除', 5, 'sys:dept:delete', NULL, NULL, '', 3, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (79, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '字典新增', 6, 'sys:dict:add', NULL, NULL, '', 1, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (81, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '字典编辑', 6, 'sys:dict:edit', NULL, NULL, '', 2, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (84, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '字典删除', 6, 'sys:dict:delete', NULL, NULL, '', 3, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (89, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'Layout', 'menu', NULL, '功能演示', 0, NULL, '', NULL, '/function', 12, 2, 1);
INSERT INTO "mypg"."sys_menu" VALUES (90, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'demo/websocket', '', 1, 'Websocket', 89, NULL, '', NULL, '/function/websocket', 3, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (91, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'demo/other', '', NULL, '敬请期待...', 89, NULL, '', NULL, 'other/:id', 4, 2, 1);
INSERT INTO "mypg"."sys_menu" VALUES (95, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'demo/dictionary', '', 1, '字典组件', 36, NULL, '', NULL, 'dict-demo', 4, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (97, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'demo/icons', 'el-icon-Notification', 1, 'Icons', 89, NULL, '', NULL, 'icon-demo', 2, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (102, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'demo/internal-doc', 'document', NULL, 'document', 26, NULL, '', '', 'internal-doc', 1, 3, 1);
INSERT INTO "mypg"."sys_menu" VALUES (106, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '用户导入', 2, 'sys:user:import', NULL, NULL, '', 5, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (107, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '用户导出', 2, 'sys:user:export', NULL, NULL, '', 6, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (108, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'demo/curd/index', '', 1, '增删改查', 36, NULL, '', NULL, 'curd', 0, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (109, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'demo/table-select/index', '', 1, '列表选择器', 36, NULL, '', NULL, 'table-select', 1, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (110, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 1, 'Layout', 'el-icon-ElementPlus', 1, '路由参数', 0, NULL, NULL, NULL, '/route-param', 11, 2, 1);
INSERT INTO "mypg"."sys_menu" VALUES (112, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, 'demo/route-param', 'el-icon-StarFilled', 1, '参数(type=2)', 110, NULL, NULL, NULL, 'route-param-type2', 2, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (117, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, 'system/log/index', 'document', 1, '系统日志', 1, NULL, NULL, 'Log', 'log', 6, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (118, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, 'Layout', 'menu', 1, '系统工具', 0, NULL, NULL, NULL, '/codegen', 2, 2, 1);
INSERT INTO "mypg"."sys_menu" VALUES (119, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, 'codegen/index', 'code', 1, '代码生成', 118, NULL, NULL, 'Codegen', 'codegen', 1, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (120, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, 'system/config/index', 'setting', 1, '系统配置', 1, NULL, NULL, 'Config', 'config', 7, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (121, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, NULL, '', 1, '系统配置查询', 120, 'sys:config:query', NULL, NULL, '', 1, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (122, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, NULL, '', 1, '系统配置新增', 120, 'sys:config:add', NULL, NULL, '', 2, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (123, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, NULL, '', 1, '系统配置修改', 120, 'sys:config:update', NULL, NULL, '', 3, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (125, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, NULL, '', 1, '系统配置刷新', 120, 'sys:config:refresh', NULL, NULL, '', 5, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (126, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'system/notice/index', '', NULL, '通知公告', 1, NULL, NULL, 'Notice', 'notice', 9, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (127, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '通知查询', 126, 'sys:notice:query', NULL, NULL, '', 1, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (128, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '通知新增', 126, 'sys:notice:add', NULL, NULL, '', 2, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (1, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'Layout', 'system', NULL, '系统管理', 0, NULL, '/system/user', '', '/system', 1, 2, 1);
INSERT INTO "mypg"."sys_menu" VALUES (5, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'system/dept/index', 'tree', 1, '部门管理', 1, NULL, NULL, 'Dept', 'dept', 4, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (24, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, 'demo/multi-level/children/children/level3-2', '', 1, '菜单三级-2', 22, NULL, '', NULL, 'multi-level3-2', 2, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (70, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '角色新增', 3, 'sys:role:add', NULL, NULL, '', 2, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (77, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '部门编辑', 5, 'sys:dept:edit', NULL, NULL, '', 2, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (88, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '重置密码', 2, 'sys:user:reset-password', NULL, NULL, '', 4, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (105, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, NULL, '', 0, '用户查询', 2, 'sys:user:query', NULL, NULL, '', 0, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (111, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, 'demo/route-param', 'el-icon-Star', 1, '参数(type=1)', 110, NULL, NULL, NULL, 'route-param-type1', 1, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (124, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, NULL, '', 1, '系统配置删除', 120, 'sys:config:delete', NULL, NULL, '', 4, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (129, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '通知编辑', 126, 'sys:notice:edit', NULL, NULL, '', 3, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (130, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '通知删除', 126, 'sys:notice:delete', NULL, NULL, '', 4, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (133, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, NULL, '', 1, '通知发布', 126, 'sys:notice:publish', NULL, NULL, '', 5, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (134, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, NULL, '', 1, '通知撤回', 126, 'sys:notice:revoke', NULL, NULL, '', 6, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (135, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, 0, 'system/dict/dict-item', '', 1, '字典项', 1, NULL, NULL, 'DictItem', 'dict-item', 6, 1, 0);
INSERT INTO "mypg"."sys_menu" VALUES (136, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '字典项新增', 135, 'sys:dict-item:add', NULL, NULL, '', 2, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (137, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '字典项编辑', 135, 'sys:dict-item:edit', NULL, NULL, '', 3, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (138, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '字典项删除', 135, 'sys:dict-item:delete', NULL, NULL, '', 4, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (139, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '角色查询', 3, 'sys:role:query', NULL, NULL, '', 1, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (140, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '菜单查询', 4, 'sys:menu:query', NULL, NULL, '', 1, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (141, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '部门查询', 5, 'sys:dept:query', NULL, NULL, '', 1, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (142, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '字典查询', 6, 'sys:dict:query', NULL, NULL, '', 1, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (143, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, NULL, '', NULL, '字典项查询', 135, 'sys:dict-item:query', NULL, NULL, '', 1, 4, 1);
INSERT INTO "mypg"."sys_menu" VALUES (144, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, '', 'document', NULL, '后端文档', 26, NULL, '', NULL, 'https://youlai.blog.csdn.net/article/details/145178880', 3, 3, 1);
INSERT INTO "mypg"."sys_menu" VALUES (145, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, '', 'document', NULL, '移动端文档', 26, NULL, '', NULL, 'https://youlai.blog.csdn.net/article/details/143222890', 4, 3, 1);
INSERT INTO "mypg"."sys_menu" VALUES (146, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'demo/drag', '', NULL, '拖拽组件', 36, NULL, '', NULL, 'drag', 5, 1, 1);
INSERT INTO "mypg"."sys_menu" VALUES (147, '2025-04-24 15:58:12+08', '2025-04-24 15:58:12+08', NULL, NULL, 'demo/text-scroll', '', NULL, '滚动文本', 36, NULL, '', NULL, 'text-scroll', 6, 1, 1);

-- ----------------------------
-- Indexes structure for table sys_menu
-- ----------------------------
CREATE INDEX "idx_sys_menu_deleted_at" ON "mypg"."sys_menu" USING btree (
  "deleted_at" "pg_catalog"."timestamptz_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table sys_menu
-- ----------------------------
ALTER TABLE "mypg"."sys_menu" ADD CONSTRAINT "sys_menu_pkey" PRIMARY KEY ("id");
