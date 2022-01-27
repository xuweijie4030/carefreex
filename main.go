package main

import (
	"github.com/carefreex-io/config"
	"github.com/carefreex-io/dbdao/gormdb"
	"github.com/carefreex-io/logger"
	"github.com/carefreex-io/rpcxserver"
)

func main() {
	config.DefaultCustomOptions.Path = "./conf/conf.yaml"
	config.InitConfig()

	logger.InitLogger()

	if err := gormdb.InitDB(); err != nil {
		logger.Errorf("mysql.InitDB failed: err=%v", err)
	}

	rpcxserver.DefaultCustomOptions.Service = NewService()
	s := rpcxserver.NewServer()

	s.Start()
}
