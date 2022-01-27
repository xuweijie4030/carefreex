package rpc

import (
	"context"
	"github.com/carefreex-io/example/proto"
	"github.com/carefreex-io/example/rpc"
	"sync"
)

type DemoRpc struct {
}

var (
	demoRpc     *DemoRpc
	demoRpcOnce sync.Once
)

func NewDemoRpc() *DemoRpc {
	if demoRpc == nil {
		demoRpcOnce.Do(func() {
			demoRpc = &DemoRpc{}
		})
	}

	return demoRpc
}

func (l *DemoRpc) Demo(ctx context.Context) (result string, err error) {
	result = "this is rpc demo"

	cli, err := rpc.NewClient()
	if err != nil {
		return result, err
	}

	request := proto.DemoHelloRequest{
		Name: "test",
	}
	response := proto.DemoHelloResponse{}
	if err = cli.Hello(ctx, &request, &response); err != nil {
		return result, err
	}

	return result, nil
}
