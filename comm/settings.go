package comm

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var CfgLoader *viper.Viper

func InitViperCfg() (err error) {
	CfgLoader = viper.New()
	CfgLoader.SetConfigFile("./conf/conf.yaml")
	CfgLoader.AddConfigPath(".")
	err = CfgLoader.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置信息失败...")
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("读取配置文件成功...")
	CfgLoader.WatchConfig()
	CfgLoader.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改了...")
	})

	return nil
}
