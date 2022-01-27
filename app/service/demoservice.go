package service

import (
	"context"
	"github.com/carefreex-io/example/proto"
)

type DemoService struct {
}

func NewDemoService() *DemoService {
	return &DemoService{}
}

func (s *DemoService) Hello(ctx context.Context, request *proto.DemoHelloRequest, response *proto.DemoHelloResponse) (err error) {
	//if response.Msg, err = logic.NewDemoLogic().Demo(ctx); err != nil {
	//	return err
	//}
	response.Msg = request.Name

	return nil
}
