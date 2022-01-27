package rpc

import (
	"context"
	"fmt"
	"github.com/carefreex-io/config"
	"github.com/carefreex-io/example/proto"
	"github.com/carefreex-io/logger"
	"testing"
)

func TestClient_Hello(t *testing.T) {
	config.DefaultCustomOptions.Path = "../conf/conf.yaml"
	config.InitConfig()

	logger.InitLogger()

	ctx := context.Background()
	request := proto.DemoHelloRequest{
		Name: "ZhangSan",
	}
	response := proto.DemoHelloResponse{}

	cli, err := NewClient()
	if err != nil {
		fmt.Printf("new cline failed: err=%v\n", err)
	}
	if err = cli.Hello(ctx, &request, &response); err != nil {
		fmt.Printf("call hello failed: err=%v\n", err)
	}
	fmt.Println(response)
}
