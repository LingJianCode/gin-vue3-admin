package config

import "fmt"

type PostgreSQL struct {
	Host         string `mapstructure:"host" json:"host" yaml:"host"`                               // 数据库地址
	Port         int32  `mapstructure:"port" json:"port" yaml:"port"`                               // 数据库端口
	Dbname       string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`                         // 数据库名
	Schema       string `mapstructure:"schema" json:"schema" yaml:"schema"`                         // 模式
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // 数据库账号
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
}

func (p *PostgreSQL) Dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai search_path=%s", p.Host, p.Username, p.Password, p.Dbname, p.Port, p.Schema)
}
