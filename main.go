package main

import (
	"carefreesky/global"
	"carefreesky/service"
	"github.com/carefreeskyio/logger"
	"github.com/carefreeskyio/rpcxserver"
	"github.com/smallnest/rpcx/util"
	"log"
	"strings"
)

func main() {
	global.InitConfig()

	logOption := logger.Option{
		Path:         global.Config.Log.Path,
		RotationTime: global.Config.Log.RotationTime,
		Level:        global.Config.Log.Level,
		WithMaxAge:   global.Config.Log.WithMaxAge,
	}
	logger.InitLogger(&logOption)

	addr, err := util.ExternalIPV4()
	if err != nil {
		log.Fatalf("get ipv4 failed: err=%v\n", err)
	}

	options := rpcxserver.Options{
		Server: rpcxserver.ServerOption{
			Name:    global.Config.Service.AppName,
			Addr:    addr,
			Network: global.Config.Service.Network,
			Port:    global.Config.Service.Port,
		},
		Registry: rpcxserver.RegistryOption{
			BasePath:       global.Config.Registry.BasePath,
			UpdateInterval: global.Config.Registry.UpdateInterval,
			Addr:           strings.Split(global.Config.Registry.Addr, " "),
			Group:          global.Config.Registry.Group,
		},
		RateLimit: rpcxserver.RateLimitOption{
			Enable:       global.Config.RateLimit.Enable,
			FillInterval: global.Config.RateLimit.FillInterval,
			Capacity:     global.Config.RateLimit.Token,
		},
		Service: service.NewService(),
	}
	s := rpcxserver.NewServer(&options)

	s.Start(global.Config.Service.Network, addr+":"+global.Config.Service.Port)
}
