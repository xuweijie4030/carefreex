package daemon

import (
	"github.com/carefreex-io/rpcxserver"
	"github.com/smallnest/rpcx/server"
)

var RegisterOnStart []func(s *server.Server)

var RegisterOnStop []func(s *server.Server)

func Register(s *rpcxserver.RpcXServer) {
	for _, fn := range RegisterOnStart {
		s.AddOnStartAction(fn)
	}
	for _, fn := range RegisterOnStop {
		s.AddOnStartAction(fn)
	}
}
