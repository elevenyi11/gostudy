package conf

import (
	"os"

	"github.com/go-kit/kit/log"
	"github.com/openzipkin/zipkin-go"
	"github.com/spf13/viper"
)

const (
	KConfigType = "CONFIG_TYPE"
)

var Zipkintracer *zipkin.Tracer
var Logger log.Logger

func initDefault() {
	viper.SetDefault(KConfigType, "yaml")
}

func init() {
	Logger = log.NewLogfmtLogger(os.Stderr)
	Logger = log.With(Logger, "ts", log.DefaultTimestampUTC)
	Logger = log.With(Logger, "caller", log.DefaultCaller)
	viper.AutomaticEnv()
	initDefault()

}

func LoadRemoteConfig() (err error) {
	//serviceInstance, err := discover.DiscoveryService(bootstrap.ConfigServerConfig.Id)
	return nil
}
