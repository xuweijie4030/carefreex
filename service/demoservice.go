package service

import (
	"carefreesky/rpc"
	"context"
	"fmt"
	"github.com/carefreeskyio/example/proto"
)

type DemoService struct {
}

func NewDemoService() *DemoService {
	return &DemoService{}
}

func (s *DemoService) Hello(ctx context.Context, request *proto.DemoHelloRequest, response *proto.DemoHelloResponse) error {
	cli, _ := rpc.NewClient()

	if err := cli.Hello(context.Background(), request, response); err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)

	return nil
}
