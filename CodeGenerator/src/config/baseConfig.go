package config

import "github.com/sajiao/goutil/configUtil"

var (
	// 是否是DEBUG模式
	DEBUG bool

	// web服务监听端口
	WebServerAddress string
)

func initBaseConfig(configObj *configUtil.XmlConfig) error {
	debug := configObj.DefaultBool("root/Debug", "", false)
	DEBUG = debug
	// 解析rpcConfig配置
	webServerAddress, err := configObj.String("root/WerbServerAddress", "")
	if err != nil {
		return err
	}
	WebServerAddress = webServerAddress

	return nil
}
