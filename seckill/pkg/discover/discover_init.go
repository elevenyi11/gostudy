package discover

import (
	"errors"
	"log"

	"github.com/elevenyi11/gostudy/seckill/pkg/loadbalance"
)

var ConsulService DiscoveryClient
var LoadBalance loadbalance.LoadBalance
var Logger *log.Logger
var NoInstanceExistedErr = errors.New("no available client")

func init() {

}
