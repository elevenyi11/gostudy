package config

import (
	"github.com/sajiao/Framework/configMgr"
	"github.com/sajiao/goutil/configUtil"
)

var (
	configObj     *configUtil.XmlConfig
	configManager = configMgr.NewConfigManager()
)

func init() {
	configManager.RegisterInitFunc()
}
