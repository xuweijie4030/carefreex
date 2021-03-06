package main

import (
	"carefreex/app/daemon"
	"github.com/carefreex-io/config"
	"github.com/carefreex-io/dbdao/gormdb"
	"github.com/carefreex-io/logger"
	"github.com/carefreex-io/rpcxserver"
)

func main() {
	config.InitConfig()

	logger.InitLogger()

	if err := gormdb.InitDB(); err != nil {
		logger.Errorf("mysql.InitDB failed: err=%v", err)
	}

	rpcxserver.DefaultCustomOptions.Service = NewService()
	s := rpcxserver.NewServer()

	daemon.Register(s)

	s.Start()
}
