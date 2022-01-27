package rpc

import (
	"context"
	"github.com/carefreex-io/config"
	"github.com/carefreex-io/example/proto"
	"github.com/carefreex-io/rpcxclient"
	"sync"
	"time"
)

type Client struct {
	XClient *rpcxclient.Client
}

var (
	c    *Client
	once sync.Once
)

func NewClient() (*Client, error) {
	var (
		err        error
		rpcXClient *rpcxclient.Client
	)

	if c == nil {
		once.Do(func() {
			rpcXClient, err = rpcxclient.NewClient(getOptions())
			if err != nil {
				return
			}
			c = &Client{
				XClient: rpcXClient,
			}
		})
	}

	return c, err
}

// 获取初始化rpcXClient客户端属性，可根据实际需求修改
func getOptions() *rpcxclient.Options {
	options := rpcxclient.DefaultOptions
	options.RegistryOption.BasePath = "/carefreex"
	options.RegistryOption.ServerName = "CarefreeX"
	options.RegistryOption.Addr = config.GetStringSlice("Registry.Addr")
	options.RegistryOption.Group = config.GetString("Registry.Group")
	options.Timeout = config.GetDuration("Rpc.Timeout") * time.Second

	return options
}

func (c *Client) Hello(ctx context.Context, request *proto.DemoHelloRequest, response *proto.DemoHelloResponse) (err error) {
	return c.XClient.Call(ctx, "Hello", request, response)
}


