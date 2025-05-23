PGDMP     ,                    }            mypg    15.12    15.12 �    �           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            �           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            �           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            �           1262    16389    mypg    DATABASE     r   CREATE DATABASE "mypg" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'en_US.UTF-8';
    DROP DATABASE "mypg";
                mypg    false                        2615    16391    mypg2    SCHEMA        CREATE SCHEMA "mypg2";
    DROP SCHEMA "mypg2";
                mypg    false                       1259    25975    casbin_rule    TABLE     +  CREATE TABLE "mypg2"."casbin_rule" (
    "id" bigint NOT NULL,
    "ptype" character varying(100),
    "v0" character varying(100),
    "v1" character varying(100),
    "v2" character varying(100),
    "v3" character varying(100),
    "v4" character varying(100),
    "v5" character varying(100)
);
 "   DROP TABLE "mypg2"."casbin_rule";
       mypg2         heap    mypg    false    6                       1259    25974    casbin_rule_id_seq    SEQUENCE     ~   CREATE SEQUENCE "mypg2"."casbin_rule_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 ,   DROP SEQUENCE "mypg2"."casbin_rule_id_seq";
       mypg2          mypg    false    6    276            �           0    0    casbin_rule_id_seq    SEQUENCE OWNED BY     Q   ALTER SEQUENCE "mypg2"."casbin_rule_id_seq" OWNED BY "mypg2"."casbin_rule"."id";
          mypg2          mypg    false    275                       1259    25889    sys_api    TABLE       CREATE TABLE "mypg2"."sys_api" (
    "id" bigint NOT NULL,
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "deleted_at" timestamp with time zone,
    "name" "text",
    "uri" "text",
    "method" "text",
    "group" "text"
);
    DROP TABLE "mypg2"."sys_api";
       mypg2         heap    mypg    false    6            �           0    0    COLUMN "sys_api"."name"    COMMENT     ;   COMMENT ON COLUMN "mypg2"."sys_api"."name" IS 'api名称';
          mypg2          mypg    false    264            �           0    0    COLUMN "sys_api"."uri"    COMMENT     :   COMMENT ON COLUMN "mypg2"."sys_api"."uri" IS 'api路径';
          mypg2          mypg    false    264            �           0    0    COLUMN "sys_api"."method"    COMMENT     W   COMMENT ON COLUMN "mypg2"."sys_api"."method" IS 'api方法:GET,POST,PUT,DELETE,PATCH';
          mypg2          mypg    false    264            �           0    0    COLUMN "sys_api"."group"    COMMENT     B   COMMENT ON COLUMN "mypg2"."sys_api"."group" IS 'api分组名称';
          mypg2          mypg    false    264                       1259    25888    sys_api_id_seq    SEQUENCE     z   CREATE SEQUENCE "mypg2"."sys_api_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE "mypg2"."sys_api_id_seq";
       mypg2          mypg    false    264    6            �           0    0    sys_api_id_seq    SEQUENCE OWNED BY     I   ALTER SEQUENCE "mypg2"."sys_api_id_seq" OWNED BY "mypg2"."sys_api"."id";
          mypg2          mypg    false    263                       1259    25833    sys_dept    TABLE     '  CREATE TABLE "mypg2"."sys_dept" (
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
    DROP TABLE "mypg2"."sys_dept";
       mypg2         heap    mypg    false    6            �           0    0    COLUMN "sys_dept"."name"    COMMENT     ?   COMMENT ON COLUMN "mypg2"."sys_dept"."name" IS '部门名称';
          mypg2          mypg    false    257            �           0    0    COLUMN "sys_dept"."code"    COMMENT     ?   COMMENT ON COLUMN "mypg2"."sys_dept"."code" IS '部门编号';
          mypg2          mypg    false    257            �           0    0    COLUMN "sys_dept"."parent_id"    COMMENT     C   COMMENT ON COLUMN "mypg2"."sys_dept"."parent_id" IS '父节点id';
          mypg2          mypg    false    257            �           0    0    COLUMN "sys_dept"."sort"    COMMENT     ?   COMMENT ON COLUMN "mypg2"."sys_dept"."sort" IS '显示顺序';
          mypg2          mypg    false    257            �           0    0    COLUMN "sys_dept"."status"    COMMENT     N   COMMENT ON COLUMN "mypg2"."sys_dept"."status" IS '状态(1-正常 0-禁用)';
          mypg2          mypg    false    257                        1259    25832    sys_dept_id_seq    SEQUENCE     {   CREATE SEQUENCE "mypg2"."sys_dept_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE "mypg2"."sys_dept_id_seq";
       mypg2          mypg    false    6    257            �           0    0    sys_dept_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE "mypg2"."sys_dept_id_seq" OWNED BY "mypg2"."sys_dept"."id";
          mypg2          mypg    false    256                       1259    25965    sys_dict_items    TABLE     D  CREATE TABLE "mypg2"."sys_dict_items" (
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
 %   DROP TABLE "mypg2"."sys_dict_items";
       mypg2         heap    mypg    false    6            �           0    0 #   COLUMN "sys_dict_items"."dict_code"    COMMENT     J   COMMENT ON COLUMN "mypg2"."sys_dict_items"."dict_code" IS '字典编码';
          mypg2          mypg    false    274            �           0    0    COLUMN "sys_dict_items"."label"    COMMENT     I   COMMENT ON COLUMN "mypg2"."sys_dict_items"."label" IS '字典项标签';
          mypg2          mypg    false    274            �           0    0    COLUMN "sys_dict_items"."sort"    COMMENT     ?   COMMENT ON COLUMN "mypg2"."sys_dict_items"."sort" IS '排序';
          mypg2          mypg    false    274            �           0    0     COLUMN "sys_dict_items"."status"    COMMENT     ^   COMMENT ON COLUMN "mypg2"."sys_dict_items"."status" IS '状态（0：禁用，1：启用）';
          mypg2          mypg    false    274            �           0    0 "   COLUMN "sys_dict_items"."tag_type"    COMMENT     a   COMMENT ON COLUMN "mypg2"."sys_dict_items"."tag_type" IS '字典类型（用于显示样式）';
          mypg2          mypg    false    274            �           0    0    COLUMN "sys_dict_items"."value"    COMMENT     F   COMMENT ON COLUMN "mypg2"."sys_dict_items"."value" IS '字典项值';
          mypg2          mypg    false    274                       1259    25964    sys_dict_items_id_seq    SEQUENCE     �   CREATE SEQUENCE "mypg2"."sys_dict_items_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 /   DROP SEQUENCE "mypg2"."sys_dict_items_id_seq";
       mypg2          mypg    false    274    6            �           0    0    sys_dict_items_id_seq    SEQUENCE OWNED BY     W   ALTER SEQUENCE "mypg2"."sys_dict_items_id_seq" OWNED BY "mypg2"."sys_dict_items"."id";
          mypg2          mypg    false    273                       1259    25955 	   sys_dicts    TABLE       CREATE TABLE "mypg2"."sys_dicts" (
    "id" bigint NOT NULL,
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "deleted_at" timestamp with time zone,
    "dict_code" "text",
    "name" "text",
    "remark" "text",
    "status" bigint
);
     DROP TABLE "mypg2"."sys_dicts";
       mypg2         heap    mypg    false    6            �           0    0    COLUMN "sys_dicts"."dict_code"    COMMENT     E   COMMENT ON COLUMN "mypg2"."sys_dicts"."dict_code" IS '字典编码';
          mypg2          mypg    false    272            �           0    0    COLUMN "sys_dicts"."name"    COMMENT     @   COMMENT ON COLUMN "mypg2"."sys_dicts"."name" IS '字典名称';
          mypg2          mypg    false    272            �           0    0    COLUMN "sys_dicts"."remark"    COMMENT     <   COMMENT ON COLUMN "mypg2"."sys_dicts"."remark" IS '备注';
          mypg2          mypg    false    272            �           0    0    COLUMN "sys_dicts"."status"    COMMENT     K   COMMENT ON COLUMN "mypg2"."sys_dicts"."status" IS '显示状态(1:显示';
          mypg2          mypg    false    272                       1259    25954    sys_dicts_id_seq    SEQUENCE     |   CREATE SEQUENCE "mypg2"."sys_dicts_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 *   DROP SEQUENCE "mypg2"."sys_dicts_id_seq";
       mypg2          mypg    false    6    272            �           0    0    sys_dicts_id_seq    SEQUENCE OWNED BY     M   ALTER SEQUENCE "mypg2"."sys_dicts_id_seq" OWNED BY "mypg2"."sys_dicts"."id";
          mypg2          mypg    false    271                       1259    25914    sys_menu    TABLE     �  CREATE TABLE "mypg2"."sys_menu" (
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
    DROP TABLE "mypg2"."sys_menu";
       mypg2         heap    mypg    false    6            �           0    0    COLUMN "sys_menu"."always_show"    COMMENT     m   COMMENT ON COLUMN "mypg2"."sys_menu"."always_show" IS '【目录】只有一个子路由是否始终显示';
          mypg2          mypg    false    267            �           0    0    COLUMN "sys_menu"."component"    COMMENT     n   COMMENT ON COLUMN "mypg2"."sys_menu"."component" IS '组件路径(vue页面完整路径，省略.vue后缀)';
          mypg2          mypg    false    267            �           0    0    COLUMN "sys_menu"."icon"    COMMENT     ?   COMMENT ON COLUMN "mypg2"."sys_menu"."icon" IS '菜单图标';
          mypg2          mypg    false    267            �           0    0    COLUMN "sys_menu"."keep_alive"    COMMENT     ]   COMMENT ON COLUMN "mypg2"."sys_menu"."keep_alive" IS '【菜单】是否开启页面缓存';
          mypg2          mypg    false    267            �           0    0    COLUMN "sys_menu"."name"    COMMENT     ?   COMMENT ON COLUMN "mypg2"."sys_menu"."name" IS '菜单名称';
          mypg2          mypg    false    267            �           0    0    COLUMN "sys_menu"."parent_id"    COMMENT     C   COMMENT ON COLUMN "mypg2"."sys_menu"."parent_id" IS '父菜单ID';
          mypg2          mypg    false    267            �           0    0    COLUMN "sys_menu"."perm"    COMMENT     ?   COMMENT ON COLUMN "mypg2"."sys_menu"."perm" IS '权限标识';
          mypg2          mypg    false    267            �           0    0    COLUMN "sys_menu"."redirect"    COMMENT     C   COMMENT ON COLUMN "mypg2"."sys_menu"."redirect" IS '跳转路径';
          mypg2          mypg    false    267            �           0    0    COLUMN "sys_menu"."route_name"    COMMENT     E   COMMENT ON COLUMN "mypg2"."sys_menu"."route_name" IS '路由名称';
          mypg2          mypg    false    267            �           0    0    COLUMN "sys_menu"."route_path"    COMMENT     E   COMMENT ON COLUMN "mypg2"."sys_menu"."route_path" IS '路由路径';
          mypg2          mypg    false    267            �           0    0    COLUMN "sys_menu"."sort"    COMMENT     V   COMMENT ON COLUMN "mypg2"."sys_menu"."sort" IS '排序(数字越小排名越靠前)';
          mypg2          mypg    false    267                        0    0    COLUMN "sys_menu"."type"    COMMENT     h   COMMENT ON COLUMN "mypg2"."sys_menu"."type" IS '菜单类型（1-菜单 2-目录 3-外链 4-按钮）';
          mypg2          mypg    false    267            
           1259    25913    sys_menu_id_seq    SEQUENCE     {   CREATE SEQUENCE "mypg2"."sys_menu_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE "mypg2"."sys_menu_id_seq";
       mypg2          mypg    false    267    6                       0    0    sys_menu_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE "mypg2"."sys_menu_id_seq" OWNED BY "mypg2"."sys_menu"."id";
          mypg2          mypg    false    266                       1259    25940    sys_menu_param    TABLE       CREATE TABLE "mypg2"."sys_menu_param" (
    "id" bigint NOT NULL,
    "created_at" timestamp with time zone,
    "updated_at" timestamp with time zone,
    "deleted_at" timestamp with time zone,
    "sys_menu_id" bigint,
    "key" "text",
    "value" "text"
);
 %   DROP TABLE "mypg2"."sys_menu_param";
       mypg2         heap    mypg    false    6                       0    0 %   COLUMN "sys_menu_param"."sys_menu_id"    COMMENT     H   COMMENT ON COLUMN "mypg2"."sys_menu_param"."sys_menu_id" IS '菜单ID';
          mypg2          mypg    false    270                       0    0    COLUMN "sys_menu_param"."key"    COMMENT     A   COMMENT ON COLUMN "mypg2"."sys_menu_param"."key" IS '参数key';
          mypg2          mypg    false    270                       0    0    COLUMN "sys_menu_param"."value"    COMMENT     E   COMMENT ON COLUMN "mypg2"."sys_menu_param"."value" IS '参数value';
          mypg2          mypg    false    270                       1259    25939    sys_menu_param_id_seq    SEQUENCE     �   CREATE SEQUENCE "mypg2"."sys_menu_param_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 /   DROP SEQUENCE "mypg2"."sys_menu_param_id_seq";
       mypg2          mypg    false    6    270                       0    0    sys_menu_param_id_seq    SEQUENCE OWNED BY     W   ALTER SEQUENCE "mypg2"."sys_menu_param_id_seq" OWNED BY "mypg2"."sys_menu_param"."id";
          mypg2          mypg    false    269                       1259    25863    sys_role    TABLE     &  CREATE TABLE "mypg2"."sys_role" (
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
    DROP TABLE "mypg2"."sys_role";
       mypg2         heap    mypg    false    6                       0    0    COLUMN "sys_role"."code"    COMMENT     ?   COMMENT ON COLUMN "mypg2"."sys_role"."code" IS '角色编码';
          mypg2          mypg    false    261                       0    0    COLUMN "sys_role"."data_scope"    COMMENT     E   COMMENT ON COLUMN "mypg2"."sys_role"."data_scope" IS '数据权限';
          mypg2          mypg    false    261                       0    0    COLUMN "sys_role"."name"    COMMENT     ?   COMMENT ON COLUMN "mypg2"."sys_role"."name" IS '角色名称';
          mypg2          mypg    false    261            	           0    0    COLUMN "sys_role"."sort"    COMMENT     9   COMMENT ON COLUMN "mypg2"."sys_role"."sort" IS '排序';
          mypg2          mypg    false    261            
           0    0    COLUMN "sys_role"."status"    COMMENT     V   COMMENT ON COLUMN "mypg2"."sys_role"."status" IS '角色状态(1-正常；0-停用)';
          mypg2          mypg    false    261            	           1259    25898    sys_role_api    TABLE     m   CREATE TABLE "mypg2"."sys_role_api" (
    "sys_api_id" bigint NOT NULL,
    "sys_role_id" bigint NOT NULL
);
 #   DROP TABLE "mypg2"."sys_role_api";
       mypg2         heap    mypg    false    6                       1259    25862    sys_role_id_seq    SEQUENCE     {   CREATE SEQUENCE "mypg2"."sys_role_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE "mypg2"."sys_role_id_seq";
       mypg2          mypg    false    261    6                       0    0    sys_role_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE "mypg2"."sys_role_id_seq" OWNED BY "mypg2"."sys_role"."id";
          mypg2          mypg    false    260                       1259    25924    sys_role_menu    TABLE     o   CREATE TABLE "mypg2"."sys_role_menu" (
    "sys_menu_id" bigint NOT NULL,
    "sys_role_id" bigint NOT NULL
);
 $   DROP TABLE "mypg2"."sys_role_menu";
       mypg2         heap    mypg    false    6                       1259    25844    sys_user    TABLE       CREATE TABLE "mypg2"."sys_user" (
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
    DROP TABLE "mypg2"."sys_user";
       mypg2         heap    mypg    false    6                       0    0    COLUMN "sys_user"."username"    COMMENT     F   COMMENT ON COLUMN "mypg2"."sys_user"."username" IS '用户登录名';
          mypg2          mypg    false    259                       0    0    COLUMN "sys_user"."nickname"    COMMENT     C   COMMENT ON COLUMN "mypg2"."sys_user"."nickname" IS '用户昵称';
          mypg2          mypg    false    259                       0    0    COLUMN "sys_user"."gender"    COMMENT     ;   COMMENT ON COLUMN "mypg2"."sys_user"."gender" IS '性别';
          mypg2          mypg    false    259                       0    0    COLUMN "sys_user"."password"    COMMENT     I   COMMENT ON COLUMN "mypg2"."sys_user"."password" IS '用户登录密码';
          mypg2          mypg    false    259                       0    0    COLUMN "sys_user"."dept_id"    COMMENT     H   COMMENT ON COLUMN "mypg2"."sys_user"."dept_id" IS '用户登录密码';
          mypg2          mypg    false    259                       0    0    COLUMN "sys_user"."avatar"    COMMENT     A   COMMENT ON COLUMN "mypg2"."sys_user"."avatar" IS '用户头像';
          mypg2          mypg    false    259                       0    0    COLUMN "sys_user"."mobile"    COMMENT     D   COMMENT ON COLUMN "mypg2"."sys_user"."mobile" IS '用户手机号';
          mypg2          mypg    false    259                       0    0    COLUMN "sys_user"."status"    COMMENT     N   COMMENT ON COLUMN "mypg2"."sys_user"."status" IS '状态(1-正常 0-禁用)';
          mypg2          mypg    false    259                       0    0    COLUMN "sys_user"."email"    COMMENT     @   COMMENT ON COLUMN "mypg2"."sys_user"."email" IS '用户邮箱';
          mypg2          mypg    false    259                       0    0    COLUMN "sys_user"."openid"    COMMENT     A   COMMENT ON COLUMN "mypg2"."sys_user"."openid" IS '微信openid';
          mypg2          mypg    false    259                       1259    25843    sys_user_id_seq    SEQUENCE     {   CREATE SEQUENCE "mypg2"."sys_user_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 )   DROP SEQUENCE "mypg2"."sys_user_id_seq";
       mypg2          mypg    false    259    6                       0    0    sys_user_id_seq    SEQUENCE OWNED BY     K   ALTER SEQUENCE "mypg2"."sys_user_id_seq" OWNED BY "mypg2"."sys_user"."id";
          mypg2          mypg    false    258                       1259    25873    sys_user_role    TABLE     o   CREATE TABLE "mypg2"."sys_user_role" (
    "sys_user_id" bigint NOT NULL,
    "sys_role_id" bigint NOT NULL
);
 $   DROP TABLE "mypg2"."sys_user_role";
       mypg2         heap    mypg    false    6                       2604    25978    casbin_rule id    DEFAULT     |   ALTER TABLE ONLY "mypg2"."casbin_rule" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."casbin_rule_id_seq"'::"regclass");
 B   ALTER TABLE "mypg2"."casbin_rule" ALTER COLUMN "id" DROP DEFAULT;
       mypg2          mypg    false    276    275    276                       2604    25892 
   sys_api id    DEFAULT     t   ALTER TABLE ONLY "mypg2"."sys_api" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."sys_api_id_seq"'::"regclass");
 >   ALTER TABLE "mypg2"."sys_api" ALTER COLUMN "id" DROP DEFAULT;
       mypg2          mypg    false    264    263    264            �           2604    25836    sys_dept id    DEFAULT     v   ALTER TABLE ONLY "mypg2"."sys_dept" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."sys_dept_id_seq"'::"regclass");
 ?   ALTER TABLE "mypg2"."sys_dept" ALTER COLUMN "id" DROP DEFAULT;
       mypg2          mypg    false    257    256    257                       2604    25968    sys_dict_items id    DEFAULT     �   ALTER TABLE ONLY "mypg2"."sys_dict_items" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."sys_dict_items_id_seq"'::"regclass");
 E   ALTER TABLE "mypg2"."sys_dict_items" ALTER COLUMN "id" DROP DEFAULT;
       mypg2          mypg    false    273    274    274                       2604    25958    sys_dicts id    DEFAULT     x   ALTER TABLE ONLY "mypg2"."sys_dicts" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."sys_dicts_id_seq"'::"regclass");
 @   ALTER TABLE "mypg2"."sys_dicts" ALTER COLUMN "id" DROP DEFAULT;
       mypg2          mypg    false    271    272    272                       2604    25917    sys_menu id    DEFAULT     v   ALTER TABLE ONLY "mypg2"."sys_menu" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."sys_menu_id_seq"'::"regclass");
 ?   ALTER TABLE "mypg2"."sys_menu" ALTER COLUMN "id" DROP DEFAULT;
       mypg2          mypg    false    266    267    267                       2604    25943    sys_menu_param id    DEFAULT     �   ALTER TABLE ONLY "mypg2"."sys_menu_param" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."sys_menu_param_id_seq"'::"regclass");
 E   ALTER TABLE "mypg2"."sys_menu_param" ALTER COLUMN "id" DROP DEFAULT;
       mypg2          mypg    false    269    270    270                        2604    25866    sys_role id    DEFAULT     v   ALTER TABLE ONLY "mypg2"."sys_role" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."sys_role_id_seq"'::"regclass");
 ?   ALTER TABLE "mypg2"."sys_role" ALTER COLUMN "id" DROP DEFAULT;
       mypg2          mypg    false    260    261    261            �           2604    25847    sys_user id    DEFAULT     v   ALTER TABLE ONLY "mypg2"."sys_user" ALTER COLUMN "id" SET DEFAULT "nextval"('"mypg2"."sys_user_id_seq"'::"regclass");
 ?   ALTER TABLE "mypg2"."sys_user" ALTER COLUMN "id" DROP DEFAULT;
       mypg2          mypg    false    258    259    259            �          0    25975    casbin_rule 
   TABLE DATA           [   COPY "mypg2"."casbin_rule" ("id", "ptype", "v0", "v1", "v2", "v3", "v4", "v5") FROM stdin;
    mypg2          mypg    false    276   0�       �          0    25889    sys_api 
   TABLE DATA           v   COPY "mypg2"."sys_api" ("id", "created_at", "updated_at", "deleted_at", "name", "uri", "method", "group") FROM stdin;
    mypg2          mypg    false    264   M�       �          0    25833    sys_dept 
   TABLE DATA           �   COPY "mypg2"."sys_dept" ("id", "created_at", "updated_at", "deleted_at", "name", "code", "parent_id", "sort", "status") FROM stdin;
    mypg2          mypg    false    257   ��       �          0    25965    sys_dict_items 
   TABLE DATA           �   COPY "mypg2"."sys_dict_items" ("id", "created_at", "updated_at", "deleted_at", "dict_code", "label", "sort", "status", "tag_type", "value") FROM stdin;
    mypg2          mypg    false    274   ۥ       �          0    25955 	   sys_dicts 
   TABLE DATA              COPY "mypg2"."sys_dicts" ("id", "created_at", "updated_at", "deleted_at", "dict_code", "name", "remark", "status") FROM stdin;
    mypg2          mypg    false    272   �       �          0    25914    sys_menu 
   TABLE DATA           �   COPY "mypg2"."sys_menu" ("id", "created_at", "updated_at", "deleted_at", "always_show", "component", "icon", "keep_alive", "name", "parent_id", "perm", "redirect", "route_name", "route_path", "sort", "type", "visible") FROM stdin;
    mypg2          mypg    false    267   ��       �          0    25940    sys_menu_param 
   TABLE DATA           z   COPY "mypg2"."sys_menu_param" ("id", "created_at", "updated_at", "deleted_at", "sys_menu_id", "key", "value") FROM stdin;
    mypg2          mypg    false    270   ̮       �          0    25863    sys_role 
   TABLE DATA           �   COPY "mypg2"."sys_role" ("id", "created_at", "updated_at", "deleted_at", "code", "data_scope", "name", "sort", "status") FROM stdin;
    mypg2          mypg    false    261   �       �          0    25898    sys_role_api 
   TABLE DATA           F   COPY "mypg2"."sys_role_api" ("sys_api_id", "sys_role_id") FROM stdin;
    mypg2          mypg    false    265   G�       �          0    25924    sys_role_menu 
   TABLE DATA           H   COPY "mypg2"."sys_role_menu" ("sys_menu_id", "sys_role_id") FROM stdin;
    mypg2          mypg    false    268   d�       �          0    25844    sys_user 
   TABLE DATA           �   COPY "mypg2"."sys_user" ("id", "created_at", "updated_at", "deleted_at", "username", "nickname", "gender", "password", "dept_id", "avatar", "mobile", "status", "email", "openid") FROM stdin;
    mypg2          mypg    false    259   ��       �          0    25873    sys_user_role 
   TABLE DATA           H   COPY "mypg2"."sys_user_role" ("sys_user_id", "sys_role_id") FROM stdin;
    mypg2          mypg    false    262   g�                  0    0    casbin_rule_id_seq    SEQUENCE SET     D   SELECT pg_catalog.setval('"mypg2"."casbin_rule_id_seq"', 1, false);
          mypg2          mypg    false    275                       0    0    sys_api_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('"mypg2"."sys_api_id_seq"', 42, true);
          mypg2          mypg    false    263                       0    0    sys_dept_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('"mypg2"."sys_dept_id_seq"', 1, false);
          mypg2          mypg    false    256                       0    0    sys_dict_items_id_seq    SEQUENCE SET     G   SELECT pg_catalog.setval('"mypg2"."sys_dict_items_id_seq"', 1, false);
          mypg2          mypg    false    273                       0    0    sys_dicts_id_seq    SEQUENCE SET     B   SELECT pg_catalog.setval('"mypg2"."sys_dicts_id_seq"', 1, false);
          mypg2          mypg    false    271                       0    0    sys_menu_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('"mypg2"."sys_menu_id_seq"', 1, false);
          mypg2          mypg    false    266                       0    0    sys_menu_param_id_seq    SEQUENCE SET     G   SELECT pg_catalog.setval('"mypg2"."sys_menu_param_id_seq"', 1, false);
          mypg2          mypg    false    269                       0    0    sys_role_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('"mypg2"."sys_role_id_seq"', 1, false);
          mypg2          mypg    false    260                       0    0    sys_user_id_seq    SEQUENCE SET     A   SELECT pg_catalog.setval('"mypg2"."sys_user_id_seq"', 1, false);
          mypg2          mypg    false    258            *           2606    25982    casbin_rule casbin_rule_pkey 
   CONSTRAINT     a   ALTER TABLE ONLY "mypg2"."casbin_rule"
    ADD CONSTRAINT "casbin_rule_pkey" PRIMARY KEY ("id");
 K   ALTER TABLE ONLY "mypg2"."casbin_rule" DROP CONSTRAINT "casbin_rule_pkey";
       mypg2            mypg    false    276                       2606    25896    sys_api sys_api_pkey 
   CONSTRAINT     Y   ALTER TABLE ONLY "mypg2"."sys_api"
    ADD CONSTRAINT "sys_api_pkey" PRIMARY KEY ("id");
 C   ALTER TABLE ONLY "mypg2"."sys_api" DROP CONSTRAINT "sys_api_pkey";
       mypg2            mypg    false    264            
           2606    25840    sys_dept sys_dept_pkey 
   CONSTRAINT     [   ALTER TABLE ONLY "mypg2"."sys_dept"
    ADD CONSTRAINT "sys_dept_pkey" PRIMARY KEY ("id");
 E   ALTER TABLE ONLY "mypg2"."sys_dept" DROP CONSTRAINT "sys_dept_pkey";
       mypg2            mypg    false    257            (           2606    25972 "   sys_dict_items sys_dict_items_pkey 
   CONSTRAINT     g   ALTER TABLE ONLY "mypg2"."sys_dict_items"
    ADD CONSTRAINT "sys_dict_items_pkey" PRIMARY KEY ("id");
 Q   ALTER TABLE ONLY "mypg2"."sys_dict_items" DROP CONSTRAINT "sys_dict_items_pkey";
       mypg2            mypg    false    274            %           2606    25962    sys_dicts sys_dicts_pkey 
   CONSTRAINT     ]   ALTER TABLE ONLY "mypg2"."sys_dicts"
    ADD CONSTRAINT "sys_dicts_pkey" PRIMARY KEY ("id");
 G   ALTER TABLE ONLY "mypg2"."sys_dicts" DROP CONSTRAINT "sys_dicts_pkey";
       mypg2            mypg    false    272            "           2606    25947 "   sys_menu_param sys_menu_param_pkey 
   CONSTRAINT     g   ALTER TABLE ONLY "mypg2"."sys_menu_param"
    ADD CONSTRAINT "sys_menu_param_pkey" PRIMARY KEY ("id");
 Q   ALTER TABLE ONLY "mypg2"."sys_menu_param" DROP CONSTRAINT "sys_menu_param_pkey";
       mypg2            mypg    false    270                       2606    25921    sys_menu sys_menu_pkey 
   CONSTRAINT     [   ALTER TABLE ONLY "mypg2"."sys_menu"
    ADD CONSTRAINT "sys_menu_pkey" PRIMARY KEY ("id");
 E   ALTER TABLE ONLY "mypg2"."sys_menu" DROP CONSTRAINT "sys_menu_pkey";
       mypg2            mypg    false    267                       2606    25902    sys_role_api sys_role_api_pkey 
   CONSTRAINT     z   ALTER TABLE ONLY "mypg2"."sys_role_api"
    ADD CONSTRAINT "sys_role_api_pkey" PRIMARY KEY ("sys_api_id", "sys_role_id");
 M   ALTER TABLE ONLY "mypg2"."sys_role_api" DROP CONSTRAINT "sys_role_api_pkey";
       mypg2            mypg    false    265    265                       2606    25928     sys_role_menu sys_role_menu_pkey 
   CONSTRAINT     }   ALTER TABLE ONLY "mypg2"."sys_role_menu"
    ADD CONSTRAINT "sys_role_menu_pkey" PRIMARY KEY ("sys_menu_id", "sys_role_id");
 O   ALTER TABLE ONLY "mypg2"."sys_role_menu" DROP CONSTRAINT "sys_role_menu_pkey";
       mypg2            mypg    false    268    268                       2606    25870    sys_role sys_role_pkey 
   CONSTRAINT     [   ALTER TABLE ONLY "mypg2"."sys_role"
    ADD CONSTRAINT "sys_role_pkey" PRIMARY KEY ("id");
 E   ALTER TABLE ONLY "mypg2"."sys_role" DROP CONSTRAINT "sys_role_pkey";
       mypg2            mypg    false    261                       2606    25854    sys_user sys_user_pkey 
   CONSTRAINT     [   ALTER TABLE ONLY "mypg2"."sys_user"
    ADD CONSTRAINT "sys_user_pkey" PRIMARY KEY ("id");
 E   ALTER TABLE ONLY "mypg2"."sys_user" DROP CONSTRAINT "sys_user_pkey";
       mypg2            mypg    false    259                       2606    25877     sys_user_role sys_user_role_pkey 
   CONSTRAINT     }   ALTER TABLE ONLY "mypg2"."sys_user_role"
    ADD CONSTRAINT "sys_user_role_pkey" PRIMARY KEY ("sys_user_id", "sys_role_id");
 O   ALTER TABLE ONLY "mypg2"."sys_user_role" DROP CONSTRAINT "sys_user_role_pkey";
       mypg2            mypg    false    262    262            +           1259    25983    idx_casbin_rule    INDEX     |   CREATE UNIQUE INDEX "idx_casbin_rule" ON "mypg2"."casbin_rule" USING "btree" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5");
 &   DROP INDEX "mypg2"."idx_casbin_rule";
       mypg2            mypg    false    276    276    276    276    276    276    276                       1259    25897    idx_sys_api_deleted_at    INDEX     Y   CREATE INDEX "idx_sys_api_deleted_at" ON "mypg2"."sys_api" USING "btree" ("deleted_at");
 -   DROP INDEX "mypg2"."idx_sys_api_deleted_at";
       mypg2            mypg    false    264                       1259    25841    idx_sys_dept_code    INDEX     V   CREATE UNIQUE INDEX "idx_sys_dept_code" ON "mypg2"."sys_dept" USING "btree" ("code");
 (   DROP INDEX "mypg2"."idx_sys_dept_code";
       mypg2            mypg    false    257                       1259    25842    idx_sys_dept_deleted_at    INDEX     [   CREATE INDEX "idx_sys_dept_deleted_at" ON "mypg2"."sys_dept" USING "btree" ("deleted_at");
 .   DROP INDEX "mypg2"."idx_sys_dept_deleted_at";
       mypg2            mypg    false    257            &           1259    25973    idx_sys_dict_items_deleted_at    INDEX     g   CREATE INDEX "idx_sys_dict_items_deleted_at" ON "mypg2"."sys_dict_items" USING "btree" ("deleted_at");
 4   DROP INDEX "mypg2"."idx_sys_dict_items_deleted_at";
       mypg2            mypg    false    274            #           1259    25963    idx_sys_dicts_deleted_at    INDEX     ]   CREATE INDEX "idx_sys_dicts_deleted_at" ON "mypg2"."sys_dicts" USING "btree" ("deleted_at");
 /   DROP INDEX "mypg2"."idx_sys_dicts_deleted_at";
       mypg2            mypg    false    272                       1259    25922    idx_sys_menu_deleted_at    INDEX     [   CREATE INDEX "idx_sys_menu_deleted_at" ON "mypg2"."sys_menu" USING "btree" ("deleted_at");
 .   DROP INDEX "mypg2"."idx_sys_menu_deleted_at";
       mypg2            mypg    false    267                       1259    25923    idx_sys_menu_name    INDEX     V   CREATE UNIQUE INDEX "idx_sys_menu_name" ON "mypg2"."sys_menu" USING "btree" ("name");
 (   DROP INDEX "mypg2"."idx_sys_menu_name";
       mypg2            mypg    false    267                        1259    25953    idx_sys_menu_param_deleted_at    INDEX     g   CREATE INDEX "idx_sys_menu_param_deleted_at" ON "mypg2"."sys_menu_param" USING "btree" ("deleted_at");
 4   DROP INDEX "mypg2"."idx_sys_menu_param_deleted_at";
       mypg2            mypg    false    270                       1259    25871    idx_sys_role_code    INDEX     V   CREATE UNIQUE INDEX "idx_sys_role_code" ON "mypg2"."sys_role" USING "btree" ("code");
 (   DROP INDEX "mypg2"."idx_sys_role_code";
       mypg2            mypg    false    261                       1259    25872    idx_sys_role_deleted_at    INDEX     [   CREATE INDEX "idx_sys_role_deleted_at" ON "mypg2"."sys_role" USING "btree" ("deleted_at");
 .   DROP INDEX "mypg2"."idx_sys_role_deleted_at";
       mypg2            mypg    false    261                       1259    25861    idx_sys_user_deleted_at    INDEX     [   CREATE INDEX "idx_sys_user_deleted_at" ON "mypg2"."sys_user" USING "btree" ("deleted_at");
 .   DROP INDEX "mypg2"."idx_sys_user_deleted_at";
       mypg2            mypg    false    259                       1259    25860    idx_sys_user_username    INDEX     ^   CREATE UNIQUE INDEX "idx_sys_user_username" ON "mypg2"."sys_user" USING "btree" ("username");
 ,   DROP INDEX "mypg2"."idx_sys_user_username";
       mypg2            mypg    false    259            3           2606    25948 !   sys_menu_param fk_sys_menu_params    FK CONSTRAINT     �   ALTER TABLE ONLY "mypg2"."sys_menu_param"
    ADD CONSTRAINT "fk_sys_menu_params" FOREIGN KEY ("sys_menu_id") REFERENCES "mypg2"."sys_menu"("id");
 P   ALTER TABLE ONLY "mypg2"."sys_menu_param" DROP CONSTRAINT "fk_sys_menu_params";
       mypg2          mypg    false    267    4125    270            /           2606    25903 $   sys_role_api fk_sys_role_api_sys_api    FK CONSTRAINT     �   ALTER TABLE ONLY "mypg2"."sys_role_api"
    ADD CONSTRAINT "fk_sys_role_api_sys_api" FOREIGN KEY ("sys_api_id") REFERENCES "mypg2"."sys_api"("id");
 S   ALTER TABLE ONLY "mypg2"."sys_role_api" DROP CONSTRAINT "fk_sys_role_api_sys_api";
       mypg2          mypg    false    265    264    4119            0           2606    25908 %   sys_role_api fk_sys_role_api_sys_role    FK CONSTRAINT     �   ALTER TABLE ONLY "mypg2"."sys_role_api"
    ADD CONSTRAINT "fk_sys_role_api_sys_role" FOREIGN KEY ("sys_role_id") REFERENCES "mypg2"."sys_role"("id");
 T   ALTER TABLE ONLY "mypg2"."sys_role_api" DROP CONSTRAINT "fk_sys_role_api_sys_role";
       mypg2          mypg    false    265    261    4114            1           2606    25929 '   sys_role_menu fk_sys_role_menu_sys_menu    FK CONSTRAINT     �   ALTER TABLE ONLY "mypg2"."sys_role_menu"
    ADD CONSTRAINT "fk_sys_role_menu_sys_menu" FOREIGN KEY ("sys_menu_id") REFERENCES "mypg2"."sys_menu"("id");
 V   ALTER TABLE ONLY "mypg2"."sys_role_menu" DROP CONSTRAINT "fk_sys_role_menu_sys_menu";
       mypg2          mypg    false    267    4125    268            2           2606    25934 '   sys_role_menu fk_sys_role_menu_sys_role    FK CONSTRAINT     �   ALTER TABLE ONLY "mypg2"."sys_role_menu"
    ADD CONSTRAINT "fk_sys_role_menu_sys_role" FOREIGN KEY ("sys_role_id") REFERENCES "mypg2"."sys_role"("id");
 V   ALTER TABLE ONLY "mypg2"."sys_role_menu" DROP CONSTRAINT "fk_sys_role_menu_sys_role";
       mypg2          mypg    false    4114    268    261            ,           2606    25855    sys_user fk_sys_user_dept    FK CONSTRAINT     �   ALTER TABLE ONLY "mypg2"."sys_user"
    ADD CONSTRAINT "fk_sys_user_dept" FOREIGN KEY ("dept_id") REFERENCES "mypg2"."sys_dept"("id");
 H   ALTER TABLE ONLY "mypg2"."sys_user" DROP CONSTRAINT "fk_sys_user_dept";
       mypg2          mypg    false    4106    259    257            -           2606    25883 '   sys_user_role fk_sys_user_role_sys_role    FK CONSTRAINT     �   ALTER TABLE ONLY "mypg2"."sys_user_role"
    ADD CONSTRAINT "fk_sys_user_role_sys_role" FOREIGN KEY ("sys_role_id") REFERENCES "mypg2"."sys_role"("id");
 V   ALTER TABLE ONLY "mypg2"."sys_user_role" DROP CONSTRAINT "fk_sys_user_role_sys_role";
       mypg2          mypg    false    262    4114    261            .           2606    25878 '   sys_user_role fk_sys_user_role_sys_user    FK CONSTRAINT     �   ALTER TABLE ONLY "mypg2"."sys_user_role"
    ADD CONSTRAINT "fk_sys_user_role_sys_user" FOREIGN KEY ("sys_user_id") REFERENCES "mypg2"."sys_user"("id");
 V   ALTER TABLE ONLY "mypg2"."sys_user_role" DROP CONSTRAINT "fk_sys_user_role_sys_user";
       mypg2          mypg    false    262    4110    259            �      x������ � �      �   (  x���[OAǟ˧�Qctgfgo}bH���o���M[��(�A���^"Z�� a��o���δ��3��L�?9�3g�efm\"�8���.a\&�L��`��"�u�v)��-F�V����ݽ�5[���a�2��ٰj����ͩ���{��Zk�F�i u<���`�P�4�2Z��W��7Z�ݽs����6'T�9����o��;�BT��2]�U���R�����\e����#�~�t�A�� v�cS�*�t���d���:9uk��T������w�=�;p�Z	��f8�>�8�G O�j����_c/nHa=O4��4 ��.�o�2�`� Z}����:����d��?;��/�U�5���x Dab����nMvN�qV�`����(�����r��֗B� 6	x�2~��?��{�o�w_ ��l��Z�X��6��S<	 �.9��B���a���h��N��ЍԼ�0�����*�����`�׫O¼��"p���v��CJ�i}e�����BvgP�Zk�U�򑥲�����9&U��Vg��R��1r<oZ�ir��K���=4�53W�0r
�?�*�Ua�y�4 ��#�iB�[Ӟg�T�SA`"c�u�,Uy�ږŏ��.�-K�'`V`��]���N�+��:�B��>M\l�T����ܙ<>�	�Ǿ�*X�Y�j%������ћ�x�G�f)Y��_&�kW��ۦ����t�1N칆+�T��΍��=��LvZ�Z���tE�m `@]x�H��c��|`���5=ruxN��|�V��32�2v4(U�#C�H5�Ev��剐�<����>2�%�:i����4]��1��:l�U���C�P�2D�33�����_T0dn�+��m�
x��*h �1���z���(� 4v���r�}�����KdVK�P�������S���6��掯y|C��sL�4�::r����c��.�s�B;|�c�����D83�@��מ �pAH��@��훝B�d�p󓪜8��u�-j��r��������{c      �   F   x�3�4202�50�52Q04�21�25�36��0��6�  ���|����_6��r100�4�b�=... �](      �   ,  x�3�4202�5 "sC+S+Csm\�1~��y)�E�ϧl�4���ĢJNC.#��z�t3�Ь�ļt ׈˘l����t}�1д̼�|N.����/�LN�/�,H�|�y������?ߵ��������b��M�a���[�u-{�F\f���t]���/�.{:�Ѐ5�2�����gs�l��|�RN��0��i�&��yڿ�ٴ/���4E
S.KJM��d�4NKKX����24 �̜Բ��'{���l���!���XN�Ey�y霾\�dd'#_����\1z\\\  �3      �   g   x�3�4202�5 "sC+S+Csm\�1~��y)�E���?�X0�2"ݔ��������ʂTΗ���_�|���!�m`NjYj��]'��qqq �eA�      �   .  x��Y[o�F~v~EA��zl�U�CEU	��PT��g=Y������DKC.PH�E$"R�T�6&�������_Bĸ�$��YsΜ�o�j��j��jT4�T��[�}���|�4��!����l��3
v*v�s+?�%)�{�ÅG��G���[�)zo��߄^��̱�Kx������g0�'���F1�2���pL�`.=�n��`���#���Nae`-�rX����mg`�!��A�b"�h����ԍ��Vu]o���خ��8�py|��hu#�{ݹ(�r�(�Ni�5uBC��.R,<�զNhW|;5�0�.�����ѻ�r��ֈ'��}>�����.[�&.��M��8��eh�2$��S�H�G�@XP�EX\��q!���T�$T,�;�Tɛ��h��py~���X�|y���q���~Щ���yۭv�Z��ZS�Z��j���V��!�1�K�_�p�y���{)�!����R�h J�
�ham��9,��B�X��!���7���_2�^�,����F[7I��d�E6=	�hˈM?�h�A��L��>�-��m_�{�&�q�ez[��� ׍�|�;\�_Y��#�b��`�-A,1�ٷ�ϔ7��x��썑�	�˪�)ם✡�9\~=}k'9C�AǴ��
фA��a$�w�FA�z������V�����W6��+��Oy=,���\'L��+O��V7�Y������pҵǑ�,G_��l��S$����5�
-�`
�+4��4w���@�j�U
�g� #fca��� F����hnC${{���	T��p)����D"�*ݣ�
�.�hm�XՖ�6�L���A�+}�7�}&4�om���ǥ_��_g��+����
�|��eE�d���xh�5yր�d�u/��x�c�ՎS�����\�����\�SX�.�LӟU��ί{;�݀U�9�����9H�M��5���'����fF���`�n�}�t*��)]���ZJIY"���*�s�~�ߥ�ҌX?V?��Gb' q��ģ���V|0��y����W7��!FPml�J�Мt0�#�����f�z���@�ѵ�zU�Z�0P��C6N�q�p�n/i��[��h�?A�f�4i�GH" *<��=8y0L�����l������p�6�R��6�L�ԡ�Q���dI�U�;���N#w����}|���-3۫�&��E��<YHO�88��rH+E�(岿�T!������8�{J�T>�a2A8D�#��	,�U�()�r��*%�Y��<'H�"-�df"�%�W*?
�Kn�N���,8�P2��|�DT�J�@�?N�ɀx�B��L�c�2L-g�\`�[.4�M��Q���h��Zb���nRf� bj0IȬ ��>H��ǉǧ}Kw����b�&;GI��f�hûϣG����E�BRS&Ź�T��J��o"=��oS#Փ��;=#��Y&A�yJ�'�Aa
z�!�U�����Fn��]��J�1�������r�.k��`��д����:j�Z-���RC��S�؉�o�k�52kT!dlmr���%�٣K�X���?���|�n�8��x&�]h���y�`�?QQ
f6��;�:2O@�#�2�YNe0
g�_x���r}����k K�C�&/�%�O<3��CZ��񔀵c0ՊU��K�bZR��i�E�h��1*滖���&��B
,�-V��=K��{9�-�T8y��\j�H�媚��00-D*�����!	�]�T�>�dH��.D*5�8��\e�R�S�,f\ 9̢8[����I��      �      x������ � �      �   N   x�3�4202�50�52U04�26�20�330�40�6�@�Z(Y��������9H6Ə3��?�� Brr��qqq �iO      �      x������ � �      �      x������ � �      �   �   x�-��n�0Eg�ت�����DJ���TM�H�kHB1�!__"E����9�!�Z�\�\edj@q�+P�ٓ��F�A�)�<�=��V]s{�`1�X@\d#�o|��y{��YX���n������0Х6��:�����i��!��I�����S�;�%MgO���TC� )י�
)��U΢�!�s��tds��U$W~
���G�	�<��K6Bc      �      x�3�4����� ]     