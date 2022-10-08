// Package setting
// @Author: fzw
// @Create: 2022/10/8
// @Description: 初始化时读取配置文件相关工具
package setting

import "gopkg.in/ini.v1"

var Conf = new(AppConfig)

//
// AppConfig
// @Description:应用配置
//
type AppConfig struct {
	Release      bool `ini:"release"` //是否是上线模式
	Port         int  `ini:"port"`    //端口
	*MySqlConfig `ini:"mysql"`
}

//
// MySqlConfig
// @Description: mysql相关配置
//
type MySqlConfig struct {
	User     string `ini:"user"`     //用户名
	Password string `ini:"password"` //密码
	DB       string `ini:"db"`       //要操作的数据库
	Host     string `ini:"host"`     //host
	Port     string `ini:"port"`     //端口
}

//
// Init
//  @Description: 初始化配置
//  @param file 配置文件路径
//  @return error
//
func Init(file string) error {
	return ini.MapTo(Conf, file)
}
