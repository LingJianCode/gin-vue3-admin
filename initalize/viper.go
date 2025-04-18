package initalize

import (
	"my-ops-admin/global"

	"github.com/spf13/viper"
)

func InitViper() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("未找到配置文件")
		} else {
			panic(err)
		}
	}
	if err := v.Unmarshal(&global.OPS_CONFIG); err != nil {
		panic(err)
	}
}
