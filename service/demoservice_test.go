package service

import (
	"carefreesky/global"
	"context"
	"github.com/carefreeskyio/example/proto"
	"testing"
)

func TestDemoService_Hello(t *testing.T) {
	global.DefaultConfigPath = "../conf/conf.yaml"
	global.InitConfig()

	ctx := context.Background()
	request := proto.DemoHelloRequest{
		Name: "zhangSan",
	}
	response := proto.DemoHelloResponse{}

	NewDemoService().Hello(ctx, &request, &response)
}
