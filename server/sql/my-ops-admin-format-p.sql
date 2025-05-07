--
-- PostgreSQL database dump
--

-- Dumped from database version 15.12
-- Dumped by pg_dump version 15.12

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: mypg2; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA "mypg2";


SET default_table_access_method = "heap";

--
-- Name: casbin_rule; Type: TABLE; Schema: mypg2; Owner: -
--

CREATE TABLE "mypg2"."casbin_rule" (
    "id" bigint NOT NULL,
    "ptype" character varying(100),
    "v0" character varying(100),
    "v1" character varying(100),
    "v2" character varying(100),
    "v3" character varying(100),
    "v4" character varying(100),
    "v5" character varying(100)
);


--
-- Name: casbin_rule_id_seq; Type: SEQUENCE; Schema: mypg2; Owner: -
--

CREATE SEQUENCE "mypg2"."casbin_rule_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: casbin_rule_id_seq; Type: SEQUENCE OWNED BY; Schema: mypg2; Owner: -
--

ALTER SEQUENCE "mypg2"."casbin_rule_id_seq" OWNED BY "mypg2"."casbin_rule"."id";


--
-- Name: sys_api; Type: TABLE; Schema: mypg2; Owner: -
--

CREATE TABLE "mypg2"."sys_api" (
    "id" bigint NOT NULL,
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "deleted_at" timestamp with time zone,
    "name" "text",
    "uri" "text",
    "method" "text",
    "group" "text"
);


--
-- Name: COLUMN "sys_api"."name"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_api"."name" IS 'api名称';


--
-- Name: COLUMN "sys_api"."uri"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_api"."uri" IS 'api路径';


--
-- Name: COLUMN "sys_api"."method"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_api"."method" IS 'api方法:GET,POST,PUT,DELETE,PATCH';


--
-- Name: COLUMN "sys_api"."group"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_api"."group" IS 'api分组名称';


--
-- Name: sys_api_id_seq; Type: SEQUENCE; Schema: mypg2; Owner: -
--

CREATE SEQUENCE "mypg2"."sys_api_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: sys_api_id_seq; Type: SEQUENCE OWNED BY; Schema: mypg2; Owner: -
--

ALTER SEQUENCE "mypg2"."sys_api_id_seq" OWNED BY "mypg2"."sys_api"."id";


--
-- Name: sys_dept; Type: TABLE; Schema: mypg2; Owner: -
--

CREATE TABLE "mypg2"."sys_dept" (
    "id" bigint NOT NULL,
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "deleted_at" timestamp with time zone,
    "name" "text",
    "code" "text",
    "parent_id" bigint,
    "sort" bigint,
    "status" smallint
);


--
-- Name: COLUMN "sys_dept"."name"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_dept"."name" IS '部门名称';


--
-- Name: COLUMN "sys_dept"."code"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_dept"."code" IS '部门编号';


--
-- Name: COLUMN "sys_dept"."parent_id"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_dept"."parent_id" IS '父节点id';


--
-- Name: COLUMN "sys_dept"."sort"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_dept"."sort" IS '显示顺序';


--
-- Name: COLUMN "sys_dept"."status"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_dept"."status" IS '状态(1-正常 0-禁用)';


--
-- Name: sys_dept_id_seq; Type: SEQUENCE; Schema: mypg2; Owner: -
--

CREATE SEQUENCE "mypg2"."sys_dept_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: sys_dept_id_seq; Type: SEQUENCE OWNED BY; Schema: mypg2; Owner: -
--

ALTER SEQUENCE "mypg2"."sys_dept_id_seq" OWNED BY "mypg2"."sys_dept"."id";


--
-- Name: sys_dict_items; Type: TABLE; Schema: mypg2; Owner: -
--

CREATE TABLE "mypg2"."sys_dict_items" (
    "id" bigint NOT NULL,
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "deleted_at" timestamp with time zone,
    "dict_code" "text",
    "label" "text",
    "sort" bigint,
    "status" bigint,
    "tag_type" "text",
    "value" "text"
);


--
-- Name: COLUMN "sys_dict_items"."dict_code"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_dict_items"."dict_code" IS '字典编码';


--
-- Name: COLUMN "sys_dict_items"."label"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_dict_items"."label" IS '字典项标签';


--
-- Name: COLUMN "sys_dict_items"."sort"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_dict_items"."sort" IS '排序';


--
-- Name: COLUMN "sys_dict_items"."status"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_dict_items"."status" IS '状态（0：禁用，1：启用）';


--
-- Name: COLUMN "sys_dict_items"."tag_type"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_dict_items"."tag_type" IS '字典类型（用于显示样式）';


--
-- Name: COLUMN "sys_dict_items"."value"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_dict_items"."value" IS '字典项值';


--
-- Name: sys_dict_items_id_seq; Type: SEQUENCE; Schema: mypg2; Owner: -
--

CREATE SEQUENCE "mypg2"."sys_dict_items_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: sys_dict_items_id_seq; Type: SEQUENCE OWNED BY; Schema: mypg2; Owner: -
--

ALTER SEQUENCE "mypg2"."sys_dict_items_id_seq" OWNED BY "mypg2"."sys_dict_items"."id";


--
-- Name: sys_dicts; Type: TABLE; Schema: mypg2; Owner: -
--

CREATE TABLE "mypg2"."sys_dicts" (
    "id" bigint NOT NULL,
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "deleted_at" timestamp with time zone,
    "dict_code" "text",
    "name" "text",
    "remark" "text",
    "status" bigint
);


--
-- Name: COLUMN "sys_dicts"."dict_code"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_dicts"."dict_code" IS '字典编码';


--
-- Name: COLUMN "sys_dicts"."name"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_dicts"."name" IS '字典名称';


--
-- Name: COLUMN "sys_dicts"."remark"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_dicts"."remark" IS '备注';


--
-- Name: COLUMN "sys_dicts"."status"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_dicts"."status" IS '显示状态(1:显示';


--
-- Name: sys_dicts_id_seq; Type: SEQUENCE; Schema: mypg2; Owner: -
--

CREATE SEQUENCE "mypg2"."sys_dicts_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: sys_dicts_id_seq; Type: SEQUENCE OWNED BY; Schema: mypg2; Owner: -
--

ALTER SEQUENCE "mypg2"."sys_dicts_id_seq" OWNED BY "mypg2"."sys_dicts"."id";


--
-- Name: sys_menu; Type: TABLE; Schema: mypg2; Owner: -
--

CREATE TABLE "mypg2"."sys_menu" (
    "id" bigint NOT NULL,
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "deleted_at" timestamp with time zone,
    "always_show" bigint,
    "component" "text",
    "icon" "text",
    "keep_alive" bigint,
    "name" "text",
    "parent_id" bigint,
    "perm" "text",
    "redirect" "text",
    "route_name" "text",
    "route_path" "text",
    "sort" bigint,
    "type" bigint,
    "visible" bigint
);


--
-- Name: COLUMN "sys_menu"."always_show"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_menu"."always_show" IS '【目录】只有一个子路由是否始终显示';


--
-- Name: COLUMN "sys_menu"."component"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_menu"."component" IS '组件路径(vue页面完整路径，省略.vue后缀)';


--
-- Name: COLUMN "sys_menu"."icon"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_menu"."icon" IS '菜单图标';


--
-- Name: COLUMN "sys_menu"."keep_alive"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_menu"."keep_alive" IS '【菜单】是否开启页面缓存';


--
-- Name: COLUMN "sys_menu"."name"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_menu"."name" IS '菜单名称';


--
-- Name: COLUMN "sys_menu"."parent_id"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_menu"."parent_id" IS '父菜单ID';


--
-- Name: COLUMN "sys_menu"."perm"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_menu"."perm" IS '权限标识';


--
-- Name: COLUMN "sys_menu"."redirect"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_menu"."redirect" IS '跳转路径';


--
-- Name: COLUMN "sys_menu"."route_name"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_menu"."route_name" IS '路由名称';


--
-- Name: COLUMN "sys_menu"."route_path"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_menu"."route_path" IS '路由路径';


--
-- Name: COLUMN "sys_menu"."sort"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_menu"."sort" IS '排序(数字越小排名越靠前)';


--
-- Name: COLUMN "sys_menu"."type"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_menu"."type" IS '菜单类型（1-菜单 2-目录 3-外链 4-按钮）';


--
-- Name: sys_menu_id_seq; Type: SEQUENCE; Schema: mypg2; Owner: -
--

CREATE SEQUENCE "mypg2"."sys_menu_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: sys_menu_id_seq; Type: SEQUENCE OWNED BY; Schema: mypg2; Owner: -
--

ALTER SEQUENCE "mypg2"."sys_menu_id_seq" OWNED BY "mypg2"."sys_menu"."id";


--
-- Name: sys_menu_param; Type: TABLE; Schema: mypg2; Owner: -
--

CREATE TABLE "mypg2"."sys_menu_param" (
    "id" bigint NOT NULL,
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "deleted_at" timestamp with time zone,
    "sys_menu_id" bigint,
    "key" "text",
    "value" "text"
);


--
-- Name: COLUMN "sys_menu_param"."sys_menu_id"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_menu_param"."sys_menu_id" IS '菜单ID';


--
-- Name: COLUMN "sys_menu_param"."key"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_menu_param"."key" IS '参数key';


--
-- Name: COLUMN "sys_menu_param"."value"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_menu_param"."value" IS '参数value';


--
-- Name: sys_menu_param_id_seq; Type: SEQUENCE; Schema: mypg2; Owner: -
--

CREATE SEQUENCE "mypg2"."sys_menu_param_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: sys_menu_param_id_seq; Type: SEQUENCE OWNED BY; Schema: mypg2; Owner: -
--

ALTER SEQUENCE "mypg2"."sys_menu_param_id_seq" OWNED BY "mypg2"."sys_menu_param"."id";


--
-- Name: sys_role; Type: TABLE; Schema: mypg2; Owner: -
--

CREATE TABLE "mypg2"."sys_role" (
    "id" bigint NOT NULL,
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "deleted_at" timestamp with time zone,
    "code" "text",
    "data_scope" bigint,
    "name" "text",
    "sort" bigint,
    "status" bigint
);


--
-- Name: COLUMN "sys_role"."code"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_role"."code" IS '角色编码';


--
-- Name: COLUMN "sys_role"."data_scope"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_role"."data_scope" IS '数据权限';


--
-- Name: COLUMN "sys_role"."name"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_role"."name" IS '角色名称';


--
-- Name: COLUMN "sys_role"."sort"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_role"."sort" IS '排序';


--
-- Name: COLUMN "sys_role"."status"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_role"."status" IS '角色状态(1-正常；0-停用)';


--
-- Name: sys_role_api; Type: TABLE; Schema: mypg2; Owner: -
--

CREATE TABLE "mypg2"."sys_role_api" (
    "sys_api_id" bigint NOT NULL,
    "sys_role_id" bigint NOT NULL
);


--
-- Name: sys_role_id_seq; Type: SEQUENCE; Schema: mypg2; Owner: -
--

CREATE SEQUENCE "mypg2"."sys_role_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: sys_role_id_seq; Type: SEQUENCE OWNED BY; Schema: mypg2; Owner: -
--

ALTER SEQUENCE "mypg2"."sys_role_id_seq" OWNED BY "mypg2"."sys_role"."id";


--
-- Name: sys_role_menu; Type: TABLE; Schema: mypg2; Owner: -
--

CREATE TABLE "mypg2"."sys_role_menu" (
    "sys_menu_id" bigint NOT NULL,
    "sys_role_id" bigint NOT NULL
);


--
-- Name: sys_user; Type: TABLE; Schema: mypg2; Owner: -
--

CREATE TABLE "mypg2"."sys_user" (
    "id" bigint NOT NULL,
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "deleted_at" timestamp with time zone,
    "username" "text",
    "nickname" "text" DEFAULT '系统用户'::"text",
    "gender" smallint,
    "password" "text",
    "dept_id" bigint,
    "avatar" "text" DEFAULT 'https://foruda.gitee.com/images/1723603502796844527/03cdca2a_716974.gif'::"text",
    "mobile" "text",
    "status" smallint DEFAULT 1,
    "email" "text",
    "openid" "text"
);


--
-- Name: COLUMN "sys_user"."username"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_user"."username" IS '用户登录名';


--
-- Name: COLUMN "sys_user"."nickname"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_user"."nickname" IS '用户昵称';


--
-- Name: COLUMN "sys_user"."gender"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_user"."gender" IS '性别';


--
-- Name: COLUMN "sys_user"."password"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_user"."password" IS '用户登录密码';


--
-- Name: COLUMN "sys_user"."dept_id"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_user"."dept_id" IS '用户登录密码';


--
-- Name: COLUMN "sys_user"."avatar"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_user"."avatar" IS '用户头像';


--
-- Name: COLUMN "sys_user"."mobile"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_user"."mobile" IS '用户手机号';


--
-- Name: COLUMN "sys_user"."status"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_user"."status" IS '状态(1-正常 0-禁用)';


--
-- Name: COLUMN "sys_user"."email"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_user"."email" IS '用户邮箱';


--
-- Name: COLUMN "sys_user"."openid"; Type: COMMENT; Schema: mypg2; Owner: -
--

COMMENT ON COLUMN "mypg2"."sys_user"."openid" IS '微信openid';


--
-- Name: sys_user_id_seq; Type: SEQUENCE; Schema: mypg2; Owner: -
--

CREATE SEQUENCE "mypg2"."sys_user_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: sys_user_id_seq; Type: SEQUENCE OWNED BY; Schema: mypg2; Owner: -
--

ALTER SEQUENCE "mypg2"."sys_user_id_seq" OWNED BY "mypg2"."sys_user"."id";


--
-- Name: sys_user_role; Type: TABLE; Schema: mypg2; Owner: -
--

CREATE TABLE "mypg2"."sys_user_role" (
    "sys_user_id" bigint NOT NULL,
    "sys_role_id" bigint NOT NULL
);


--
-- Name: casbin_rule id; Type: DEFAULT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."casbin_rule" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."casbin_rule_id_seq"'::"regclass");


--
-- Name: sys_api id; Type: DEFAULT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_api" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."sys_api_id_seq"'::"regclass");


--
-- Name: sys_dept id; Type: DEFAULT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_dept" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."sys_dept_id_seq"'::"regclass");


--
-- Name: sys_dict_items id; Type: DEFAULT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_dict_items" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."sys_dict_items_id_seq"'::"regclass");


--
-- Name: sys_dicts id; Type: DEFAULT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_dicts" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."sys_dicts_id_seq"'::"regclass");


--
-- Name: sys_menu id; Type: DEFAULT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_menu" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."sys_menu_id_seq"'::"regclass");


--
-- Name: sys_menu_param id; Type: DEFAULT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_menu_param" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."sys_menu_param_id_seq"'::"regclass");


--
-- Name: sys_role id; Type: DEFAULT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_role" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."sys_role_id_seq"'::"regclass");


--
-- Name: sys_user id; Type: DEFAULT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_user" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."sys_user_id_seq"'::"regclass");


--
-- Data for Name: casbin_rule; Type: TABLE DATA; Schema: mypg2; Owner: -
--

COPY "mypg2"."casbin_rule" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") FROM stdin;
\.


--
-- Data for Name: sys_api; Type: TABLE DATA; Schema: mypg2; Owner: -
--

COPY "mypg2"."sys_api" ("id", "created_at", "updated_at", "deleted_at", "name", "uri", "method", "group") FROM stdin;
31	2025-05-07 11:20:49.221677+08	2025-05-07 11:20:49.221677+08	\N	字典分页列表	/api/v1/dicts/page	GET	字典管理
30	2025-05-07 11:20:49.220908+08	2025-05-07 11:20:49.220908+08	\N	字典列表	/api/v1/dicts	GET	字典管理
32	2025-05-07 11:20:49.222419+08	2025-05-07 11:20:49.222419+08	\N	新增字典	/api/v1/dicts	POST	字典管理
33	2025-05-07 11:20:49.223195+08	2025-05-07 11:20:49.223195+08	\N	字典表单数据	/api/v1/dicts/:dictId/form	GET	字典管理
34	2025-05-07 11:20:49.223883+08	2025-05-07 11:20:49.223883+08	\N	修改字典	/api/v1/dicts/:dictId	PUT	字典管理
35	2025-05-07 11:20:49.224534+08	2025-05-07 11:20:49.224534+08	\N	删除字典	/api/v1/dicts/:dictId	DELETE	字典管理
36	2025-05-07 11:20:49.225415+08	2025-05-07 11:20:49.225415+08	\N	字典项分页列表	/api/v1/dict/:dictCode/items/page	GET	字典管理
4	2025-05-07 11:20:49.199272+08	2025-05-07 11:20:49.199272+08	\N	创建用户	/api/v1/users	POST	用户管理
5	2025-05-07 11:20:49.200365+08	2025-05-07 11:20:49.200365+08	\N	获取当前登录用户信息	/api/v1/users/me	GET	用户管理
6	2025-05-07 11:20:49.201232+08	2025-05-07 11:20:49.201232+08	\N	用户分页列表	/api/v1/users/page	GET	用户管理
7	2025-05-07 11:20:49.20229+08	2025-05-07 11:20:49.20229+08	\N	用户表单数据	/api/v1/users/:userId/form	GET	用户管理
8	2025-05-07 11:20:49.203299+08	2025-05-07 11:20:49.203299+08	\N	重置用户密码	/api/v1/users/:userId/password/reset	PUT	用户管理
9	2025-05-07 11:20:49.204124+08	2025-05-07 11:20:49.204124+08	\N	修改用户	/api/v1/users/:userId	PUT	用户管理
10	2025-05-07 11:20:49.204975+08	2025-05-07 11:20:49.204975+08	\N	删除用户	/api/v1/users/:userId	DELETE	用户管理
21	2025-05-07 11:20:49.213155+08	2025-05-07 11:20:49.213155+08	\N	新增角色	/api/v1/roles	POST	角色管理
22	2025-05-07 11:20:49.213958+08	2025-05-07 11:20:49.213958+08	\N	角色下拉列表	/api/v1/roles/options	GET	角色管理
23	2025-05-07 11:20:49.214869+08	2025-05-07 11:20:49.214869+08	\N	角色分页列表	/api/v1/roles/page	GET	角色管理
11	2025-05-07 11:20:49.205728+08	2025-05-07 11:20:49.205728+08	\N	部门列表	/api/v1/dept	GET	部门管理
12	2025-05-07 11:20:49.206547+08	2025-05-07 11:20:49.206547+08	\N	新增部门	/api/v1/dept	POST	部门管理
13	2025-05-07 11:20:49.207271+08	2025-05-07 11:20:49.207271+08	\N	部门下拉列表	/api/v1/dept/options	GET	部门管理
24	2025-05-07 11:20:49.215663+08	2025-05-07 11:20:49.215663+08	\N	角色表单数据	/api/v1/roles/:roleId/form	GET	角色管理
25	2025-05-07 11:20:49.216386+08	2025-05-07 11:20:49.216386+08	\N	修改角色	/api/v1/roles/:roleId	PUT	角色管理
37	2025-05-07 11:20:49.226136+08	2025-05-07 11:20:49.226136+08	\N	字典项列表	/api/v1/dict/:dictCode/items	GET	字典管理
27	2025-05-07 11:20:49.218578+08	2025-05-07 11:20:49.218578+08	\N	分配菜单(包括按钮权限)给角色	/api/v1/roles/:roleId/menus	PUT	角色管理
26	2025-05-07 11:20:49.217695+08	2025-05-07 11:20:49.217695+08	\N	获取角色的菜单ID集合	/api/v1/roles/:roleId/menuIds	GET	角色管理
28	2025-05-07 11:20:49.219464+08	2025-05-07 11:20:49.219464+08	\N	获取角色api	/api/v1/roles/:roleId/apiIds	GET	角色管理
29	2025-05-07 11:20:49.220148+08	2025-05-07 11:20:49.220148+08	\N	分配角色api接口权限	/api/v1/roles/:roleId/apiIds	PUT	角色管理
14	2025-05-07 11:20:49.208086+08	2025-05-07 11:20:49.208086+08	\N	菜单列表	/api/v1/menus	GET	菜单管理
15	2025-05-07 11:20:49.208809+08	2025-05-07 11:20:49.208809+08	\N	新增菜单	/api/v1/menus	POST	菜单管理
16	2025-05-07 11:20:49.209559+08	2025-05-07 11:20:49.209559+08	\N	修改菜单	/api/v1/menus/:menuId	PUT	菜单管理
17	2025-05-07 11:20:49.210279+08	2025-05-07 11:20:49.210279+08	\N	删除菜单	/api/v1/menus/:menuId	DELETE	菜单管理
18	2025-05-07 11:20:49.211+08	2025-05-07 11:20:49.211+08	\N	菜单表单数据	/api/v1/menus/:menuId/form	GET	菜单管理
19	2025-05-07 11:20:49.211725+08	2025-05-07 11:20:49.211725+08	\N	菜单表单数据	/api/v1/menus/options	GET	菜单管理
20	2025-05-07 11:20:49.212444+08	2025-05-07 11:20:49.212444+08	\N	菜单路由列表	/api/v1/menus/routes	GET	菜单管理
38	2025-05-07 11:20:49.226823+08	2025-05-07 11:20:49.226823+08	\N	新增字典项	/api/v1/dict/:dictCode/items	POST	字典管理
39	2025-05-07 11:20:49.227577+08	2025-05-07 11:20:49.227577+08	\N	字典项表单数据	/api/v1/dict/:dictCode/items/:itemId/form	GET	字典管理
40	2025-05-07 11:20:49.228404+08	2025-05-07 11:20:49.228404+08	\N	修改字典项	/api/v1/dict/:dictCode/items/:itemId	PUT	字典管理
41	2025-05-07 11:20:49.229031+08	2025-05-07 11:20:49.229031+08	\N	删除字典项	/api/v1/dict/:dictCode/items/:itemIds	DELETE	字典管理
42	2025-05-07 11:20:49.229847+08	2025-05-07 11:20:49.229847+08	\N	获取api下拉列表	/api/v1/apis/options	GET	api管理
\.


--
-- Data for Name: sys_dept; Type: TABLE DATA; Schema: mypg2; Owner: -
--

COPY "mypg2"."sys_dept" ("id", "created_at", "updated_at", "deleted_at", "name", "code", "parent_id", "sort", "status") FROM stdin;
1	2025-04-24 15:42:52.309828+08	2025-04-24 15:42:52.309828+08	\N	研发部	RD001	0	1	0
\.


--
-- Data for Name: sys_dict_items; Type: TABLE DATA; Schema: mypg2; Owner: -
--

COPY "mypg2"."sys_dict_items" ("id", "created_at", "updated_at", "deleted_at", "dict_code", "label", "sort", "status", "tag_type", "value") FROM stdin;
1	2025-05-07 10:58:17+08	2025-05-07 10:58:17+08	\N	gender	男	1	1	primary	1
2	2025-05-07 10:58:17+08	2025-05-07 10:58:17+08	\N	gender	女	2	1	danger	2
3	2025-05-07 10:58:17+08	2025-05-07 10:58:17+08	\N	gender	保密	3	1	info	0
4	2025-05-07 10:58:17+08	2025-05-07 10:58:17+08	\N	notice_type	系统升级	1	1	success	1
5	2025-05-07 10:58:17+08	2025-05-07 10:58:17+08	\N	notice_type	系统维护	2	1	primary	2
6	2025-05-07 10:58:17+08	2025-05-07 10:58:17+08	\N	notice_type	安全警告	3	1	danger	3
7	2025-05-07 10:58:17+08	2025-05-07 10:58:17+08	\N	notice_type	假期通知	4	1	success	4
8	2025-05-07 10:58:17+08	2025-05-07 10:58:17+08	\N	notice_type	公司新闻	5	1	primary	5
9	2025-05-07 10:58:17+08	2025-05-07 10:58:17+08	\N	notice_type	其他	99	1	info	99
10	2025-05-07 10:58:17+08	2025-05-07 10:58:17+08	\N	notice_level	低	1	1	info	L
11	2025-05-07 10:58:17+08	2025-05-07 10:58:17+08	\N	notice_level	中	2	1	warning	M
12	2025-05-07 10:58:17+08	2025-05-07 10:58:17+08	\N	notice_level	高	3	1	danger	H
\.


--
-- Data for Name: sys_dicts; Type: TABLE DATA; Schema: mypg2; Owner: -
--

COPY "mypg2"."sys_dicts" ("id", "created_at", "updated_at", "deleted_at", "dict_code", "name", "remark", "status") FROM stdin;
1	2025-05-07 10:58:17+08	2025-05-07 10:58:17+08	\N	gender	性别	\N	1
2	2025-05-07 10:58:17+08	2025-05-07 10:58:17+08	\N	notice_type	通知类型	\N	1
3	2025-05-07 10:58:17+08	2025-05-07 10:58:17+08	\N	notice_level	通知级别	\N	1
\.


--
-- Data for Name: sys_menu; Type: TABLE DATA; Schema: mypg2; Owner: -
--

COPY "mypg2"."sys_menu" ("id", "created_at", "updated_at", "deleted_at", "always_show", "component", "icon", "keep_alive", "name", "parent_id", "perm", "redirect", "route_name", "route_path", "sort", "type", "visible") FROM stdin;
2	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	system/user/index	el-icon-User	1	用户管理	1	\N	\N	User	user	1	1	1
3	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	system/role/index	role	1	角色管理	1	\N	\N	Role	role	2	1	1
4	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	system/menu/index	menu	1	菜单管理	1	\N	\N	SysMenu	menu	3	1	1
6	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	system/dict/index	dict	1	字典管理	1	\N	\N	Dict	dict	5	1	1
126	2025-04-24 15:58:12+08	2025-04-28 14:54:20.333606+08	\N	0	system/notice/index		0	通知公告	1			Notice	notice	9	1	0
21	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	1	demo/multi-level/level1		\N	菜单一级	20	\N		\N	multi-level1	1	1	1
22	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	0	demo/multi-level/children/level2		\N	菜单二级	21	\N	\N	\N	multi-level2	1	1	1
23	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	0	demo/multi-level/children/children/level3-1		1	菜单三级-1	22	\N		\N	multi-level3-1	1	1	1
30	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N		document	\N	平台文档(外链)	26	\N		\N	https://juejin.cn/post/7228990409909108793	2	3	1
32	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	用户编辑	2	sys:user:edit		\N		2	4	1
72	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	角色删除	3	sys:role:delete	\N	\N		4	4	1
37	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	demo/wang-editor		1	富文本编辑器	36	\N		\N	wang-editor	2	1	1
38	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	demo/upload		1	图片上传	36	\N		\N	upload	3	1	1
39	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	demo/icon-selector		1	图标选择器	36	\N		\N	icon-selector	4	1	1
41	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	demo/api/apifox	api	1	Apifox	40	\N		\N	apifox	1	1	1
73	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	菜单新增	4	sys:menu:add	\N	\N		1	4	1
117	2025-04-24 15:58:12+08	2025-04-28 14:53:06.140903+08	\N	0	system/log/index	document	1	系统日志	1			Log	log	6	1	0
74	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	菜单编辑	4	sys:menu:edit	\N	\N		3	4	1
75	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	菜单删除	4	sys:menu:delete	\N	\N		3	4	1
76	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	部门新增	5	sys:dept:add	\N	\N		1	4	1
78	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	部门删除	5	sys:dept:delete	\N	\N		3	4	1
79	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	字典新增	6	sys:dict:add	\N	\N		1	4	1
81	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	字典编辑	6	sys:dict:edit	\N	\N		2	4	1
106	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	用户导入	2	sys:user:import	\N	\N		5	4	1
31	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	用户新增	2	sys:user:add		\N		1	4	1
112	2025-04-24 15:58:12+08	2025-04-28 11:29:54.820632+08	\N	0	demo/route-param	el-icon-StarFilled	1	参数(type=2)	110			zzz	route-param-type2	2	1	1
90	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	demo/websocket		1	Websocket	89	\N		\N	/function/websocket	3	1	1
95	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	demo/dictionary		1	字典组件	36	\N		\N	dict-demo	4	1	1
97	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	demo/icons	el-icon-Notification	1	Icons	89	\N		\N	icon-demo	2	1	1
102	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	demo/internal-doc	document	\N	document	26	\N			internal-doc	1	3	1
107	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	用户导出	2	sys:user:export	\N	\N		6	4	1
33	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	用户删除	2	sys:user:delete		\N		3	4	1
108	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	demo/curd/index		1	增删改查	36	\N		\N	curd	0	1	1
109	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	demo/table-select/index		1	列表选择器	36	\N		\N	table-select	1	1	1
120	2025-04-24 15:58:12+08	2025-04-28 14:53:21.525441+08	\N	0	system/config/index	setting	1	系统配置	1			Config	config	7	1	0
119	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	0	codegen/index	code	1	代码生成	118	\N	\N	Codegen	codegen	1	1	1
122	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	0	\N		1	系统配置新增	120	sys:config:add	\N	\N		2	4	1
123	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	0	\N		1	系统配置修改	120	sys:config:update	\N	\N		3	4	1
125	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	0	\N		1	系统配置刷新	120	sys:config:refresh	\N	\N		5	4	1
127	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	通知查询	126	sys:notice:query	\N	\N		1	4	1
71	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	角色编辑	3	sys:role:edit	\N	\N		3	4	1
128	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	通知新增	126	sys:notice:add	\N	\N		2	4	1
84	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	字典删除	6	sys:dict:delete	\N	\N		3	4	1
121	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	0	\N		1	系统配置查询	120	sys:config:query	\N	\N		1	4	1
1	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	Layout	system	\N	系统管理	0	\N	/system/user		/system	1	2	1
5	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	system/dept/index	tree	1	部门管理	1	\N	\N	Dept	dept	4	1	1
24	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	0	demo/multi-level/children/children/level3-2		1	菜单三级-2	22	\N		\N	multi-level3-2	2	1	1
111	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	0	demo/route-param	el-icon-Star	1	参数(type=1)	110	\N	\N	\N	route-param-type1	1	1	1
133	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	0	\N		1	通知发布	126	sys:notice:publish	\N	\N		5	4	1
134	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	0	\N		1	通知撤回	126	sys:notice:revoke	\N	\N		6	4	1
135	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	0	system/dict/dict-item		1	字典项	1	\N	\N	DictItem	dict-item	6	1	0
136	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	字典项新增	135	sys:dict-item:add	\N	\N		2	4	1
144	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N		document	\N	后端文档	26	\N		\N	https://youlai.blog.csdn.net/article/details/145178880	3	3	1
145	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N		document	\N	移动端文档	26	\N		\N	https://youlai.blog.csdn.net/article/details/143222890	4	3	1
146	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	demo/drag		\N	拖拽组件	36	\N		\N	drag	5	1	1
147	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	demo/text-scroll		\N	滚动文本	36	\N		\N	text-scroll	6	1	1
70	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	角色新增	3	sys:role:add	\N	\N		2	4	1
77	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	部门编辑	5	sys:dept:edit	\N	\N		2	4	1
88	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	重置密码	2	sys:user:reset-password	\N	\N		4	4	1
105	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	0	\N		0	用户查询	2	sys:user:query	\N	\N		0	4	1
124	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	0	\N		1	系统配置删除	120	sys:config:delete	\N	\N		4	4	1
129	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	通知编辑	126	sys:notice:edit	\N	\N		3	4	1
130	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	通知删除	126	sys:notice:delete	\N	\N		4	4	1
137	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	字典项编辑	135	sys:dict-item:edit	\N	\N		3	4	1
139	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	角色查询	3	sys:role:query	\N	\N		1	4	1
140	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	菜单查询	4	sys:menu:query	\N	\N		1	4	1
141	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	部门查询	5	sys:dept:query	\N	\N		1	4	1
142	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	字典查询	6	sys:dict:query	\N	\N		1	4	1
143	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	字典项查询	135	sys:dict-item:query	\N	\N		1	4	1
138	2025-04-24 15:58:12+08	2025-04-24 15:58:12+08	\N	\N	\N		\N	字典项删除	135	sys:dict-item:delete	\N	\N		4	4	1
\.


--
-- Data for Name: sys_menu_param; Type: TABLE DATA; Schema: mypg2; Owner: -
--

COPY "mypg2"."sys_menu_param" ("id", "created_at", "updated_at", "deleted_at", "sys_menu_id", "key", "value") FROM stdin;
\.


--
-- Data for Name: sys_role; Type: TABLE DATA; Schema: mypg2; Owner: -
--

COPY "mypg2"."sys_role" ("id", "created_at", "updated_at", "deleted_at", "code", "data_scope", "name", "sort", "status") FROM stdin;
1	2025-04-25 10:34:00.605905+08	2025-04-28 12:22:27.933967+08	\N	ROOT	0	ROOT	1	1
\.


--
-- Data for Name: sys_role_api; Type: TABLE DATA; Schema: mypg2; Owner: -
--

COPY "mypg2"."sys_role_api" ("sys_api_id", "sys_role_id") FROM stdin;
\.


--
-- Data for Name: sys_role_menu; Type: TABLE DATA; Schema: mypg2; Owner: -
--

COPY "mypg2"."sys_role_menu" ("sys_menu_id", "sys_role_id") FROM stdin;
\.


--
-- Data for Name: sys_user; Type: TABLE DATA; Schema: mypg2; Owner: -
--

COPY "mypg2"."sys_user" ("id", "created_at", "updated_at", "deleted_at", "username", "nickname", "gender", "password", "dept_id", "avatar", "mobile", "status", "email", "openid") FROM stdin;
1	2025-04-24 15:43:05.695057+08	2025-05-06 10:50:22.376992+08	\N	admin	admin	1	$2a$10$I8s2NH.Ft6ShFzzUl8tB0uckpTXjWnGp7Ta0r7jIrCNbBtcYeIZ7u	1	https://foruda.gitee.com/images/1723603502796844527/03cdca2a_716974.gif		1	i1cl5v74@outlook.com	17
\.


--
-- Data for Name: sys_user_role; Type: TABLE DATA; Schema: mypg2; Owner: -
--

COPY "mypg2"."sys_user_role" ("sys_user_id", "sys_role_id") FROM stdin;
1	1
\.


--
-- Name: casbin_rule_id_seq; Type: SEQUENCE SET; Schema: mypg2; Owner: -
--

SELECT pg_catalog.setval('"mypg2"."casbin_rule_id_seq"', 1, false);


--
-- Name: sys_api_id_seq; Type: SEQUENCE SET; Schema: mypg2; Owner: -
--

SELECT pg_catalog.setval('"mypg2"."sys_api_id_seq"', 42, true);


--
-- Name: sys_dept_id_seq; Type: SEQUENCE SET; Schema: mypg2; Owner: -
--

SELECT pg_catalog.setval('"mypg2"."sys_dept_id_seq"', 1, false);


--
-- Name: sys_dict_items_id_seq; Type: SEQUENCE SET; Schema: mypg2; Owner: -
--

SELECT pg_catalog.setval('"mypg2"."sys_dict_items_id_seq"', 1, false);


--
-- Name: sys_dicts_id_seq; Type: SEQUENCE SET; Schema: mypg2; Owner: -
--

SELECT pg_catalog.setval('"mypg2"."sys_dicts_id_seq"', 1, false);


--
-- Name: sys_menu_id_seq; Type: SEQUENCE SET; Schema: mypg2; Owner: -
--

SELECT pg_catalog.setval('"mypg2"."sys_menu_id_seq"', 1, false);


--
-- Name: sys_menu_param_id_seq; Type: SEQUENCE SET; Schema: mypg2; Owner: -
--

SELECT pg_catalog.setval('"mypg2"."sys_menu_param_id_seq"', 1, false);


--
-- Name: sys_role_id_seq; Type: SEQUENCE SET; Schema: mypg2; Owner: -
--

SELECT pg_catalog.setval('"mypg2"."sys_role_id_seq"', 1, false);


--
-- Name: sys_user_id_seq; Type: SEQUENCE SET; Schema: mypg2; Owner: -
--

SELECT pg_catalog.setval('"mypg2"."sys_user_id_seq"', 1, false);


--
-- Name: casbin_rule casbin_rule_pkey; Type: CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."casbin_rule"
    ADD CONSTRAINT "casbin_rule_pkey" PRIMARY KEY ("id");


--
-- Name: sys_api sys_api_pkey; Type: CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_api"
    ADD CONSTRAINT "sys_api_pkey" PRIMARY KEY ("id");


--
-- Name: sys_dept sys_dept_pkey; Type: CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_dept"
    ADD CONSTRAINT "sys_dept_pkey" PRIMARY KEY ("id");


--
-- Name: sys_dict_items sys_dict_items_pkey; Type: CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_dict_items"
    ADD CONSTRAINT "sys_dict_items_pkey" PRIMARY KEY ("id");


--
-- Name: sys_dicts sys_dicts_pkey; Type: CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_dicts"
    ADD CONSTRAINT "sys_dicts_pkey" PRIMARY KEY ("id");


--
-- Name: sys_menu_param sys_menu_param_pkey; Type: CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_menu_param"
    ADD CONSTRAINT "sys_menu_param_pkey" PRIMARY KEY ("id");


--
-- Name: sys_menu sys_menu_pkey; Type: CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_menu"
    ADD CONSTRAINT "sys_menu_pkey" PRIMARY KEY ("id");


--
-- Name: sys_role_api sys_role_api_pkey; Type: CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_role_api"
    ADD CONSTRAINT "sys_role_api_pkey" PRIMARY KEY ("sys_api_id", "sys_role_id");


--
-- Name: sys_role_menu sys_role_menu_pkey; Type: CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_role_menu"
    ADD CONSTRAINT "sys_role_menu_pkey" PRIMARY KEY ("sys_menu_id", "sys_role_id");


--
-- Name: sys_role sys_role_pkey; Type: CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_role"
    ADD CONSTRAINT "sys_role_pkey" PRIMARY KEY ("id");


--
-- Name: sys_user sys_user_pkey; Type: CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_user"
    ADD CONSTRAINT "sys_user_pkey" PRIMARY KEY ("id");


--
-- Name: sys_user_role sys_user_role_pkey; Type: CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_user_role"
    ADD CONSTRAINT "sys_user_role_pkey" PRIMARY KEY ("sys_user_id", "sys_role_id");


--
-- Name: idx_casbin_rule; Type: INDEX; Schema: mypg2; Owner: -
--

CREATE UNIQUE INDEX "idx_casbin_rule" ON "mypg2"."casbin_rule" USING "btree" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5");


--
-- Name: idx_sys_api_deleted_at; Type: INDEX; Schema: mypg2; Owner: -
--

CREATE INDEX "idx_sys_api_deleted_at" ON "mypg2"."sys_api" USING "btree" ("deleted_at");


--
-- Name: idx_sys_dept_code; Type: INDEX; Schema: mypg2; Owner: -
--

CREATE UNIQUE INDEX "idx_sys_dept_code" ON "mypg2"."sys_dept" USING "btree" ("code");


--
-- Name: idx_sys_dept_deleted_at; Type: INDEX; Schema: mypg2; Owner: -
--

CREATE INDEX "idx_sys_dept_deleted_at" ON "mypg2"."sys_dept" USING "btree" ("deleted_at");


--
-- Name: idx_sys_dict_items_deleted_at; Type: INDEX; Schema: mypg2; Owner: -
--

CREATE INDEX "idx_sys_dict_items_deleted_at" ON "mypg2"."sys_dict_items" USING "btree" ("deleted_at");


--
-- Name: idx_sys_dicts_deleted_at; Type: INDEX; Schema: mypg2; Owner: -
--

CREATE INDEX "idx_sys_dicts_deleted_at" ON "mypg2"."sys_dicts" USING "btree" ("deleted_at");


--
-- Name: idx_sys_menu_deleted_at; Type: INDEX; Schema: mypg2; Owner: -
--

CREATE INDEX "idx_sys_menu_deleted_at" ON "mypg2"."sys_menu" USING "btree" ("deleted_at");


--
-- Name: idx_sys_menu_name; Type: INDEX; Schema: mypg2; Owner: -
--

CREATE UNIQUE INDEX "idx_sys_menu_name" ON "mypg2"."sys_menu" USING "btree" ("name");


--
-- Name: idx_sys_menu_param_deleted_at; Type: INDEX; Schema: mypg2; Owner: -
--

CREATE INDEX "idx_sys_menu_param_deleted_at" ON "mypg2"."sys_menu_param" USING "btree" ("deleted_at");


--
-- Name: idx_sys_role_code; Type: INDEX; Schema: mypg2; Owner: -
--

CREATE UNIQUE INDEX "idx_sys_role_code" ON "mypg2"."sys_role" USING "btree" ("code");


--
-- Name: idx_sys_role_deleted_at; Type: INDEX; Schema: mypg2; Owner: -
--

CREATE INDEX "idx_sys_role_deleted_at" ON "mypg2"."sys_role" USING "btree" ("deleted_at");


--
-- Name: idx_sys_user_deleted_at; Type: INDEX; Schema: mypg2; Owner: -
--

CREATE INDEX "idx_sys_user_deleted_at" ON "mypg2"."sys_user" USING "btree" ("deleted_at");


--
-- Name: idx_sys_user_username; Type: INDEX; Schema: mypg2; Owner: -
--

CREATE UNIQUE INDEX "idx_sys_user_username" ON "mypg2"."sys_user" USING "btree" ("username");


--
-- Name: sys_menu_param fk_sys_menu_params; Type: FK CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_menu_param"
    ADD CONSTRAINT "fk_sys_menu_params" FOREIGN KEY ("sys_menu_id") REFERENCES "mypg2"."sys_menu"("id");


--
-- Name: sys_role_api fk_sys_role_api_sys_api; Type: FK CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_role_api"
    ADD CONSTRAINT "fk_sys_role_api_sys_api" FOREIGN KEY ("sys_api_id") REFERENCES "mypg2"."sys_api"("id");


--
-- Name: sys_role_api fk_sys_role_api_sys_role; Type: FK CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_role_api"
    ADD CONSTRAINT "fk_sys_role_api_sys_role" FOREIGN KEY ("sys_role_id") REFERENCES "mypg2"."sys_role"("id");


--
-- Name: sys_role_menu fk_sys_role_menu_sys_menu; Type: FK CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_role_menu"
    ADD CONSTRAINT "fk_sys_role_menu_sys_menu" FOREIGN KEY ("sys_menu_id") REFERENCES "mypg2"."sys_menu"("id");


--
-- Name: sys_role_menu fk_sys_role_menu_sys_role; Type: FK CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_role_menu"
    ADD CONSTRAINT "fk_sys_role_menu_sys_role" FOREIGN KEY ("sys_role_id") REFERENCES "mypg2"."sys_role"("id");


--
-- Name: sys_user fk_sys_user_dept; Type: FK CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_user"
    ADD CONSTRAINT "fk_sys_user_dept" FOREIGN KEY ("dept_id") REFERENCES "mypg2"."sys_dept"("id");


--
-- Name: sys_user_role fk_sys_user_role_sys_role; Type: FK CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_user_role"
    ADD CONSTRAINT "fk_sys_user_role_sys_role" FOREIGN KEY ("sys_role_id") REFERENCES "mypg2"."sys_role"("id");


--
-- Name: sys_user_role fk_sys_user_role_sys_user; Type: FK CONSTRAINT; Schema: mypg2; Owner: -
--

ALTER TABLE ONLY "mypg2"."sys_user_role"
    ADD CONSTRAINT "fk_sys_user_role_sys_user" FOREIGN KEY ("sys_user_id") REFERENCES "mypg2"."sys_user"("id");


--
-- PostgreSQL database dump complete
--

