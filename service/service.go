package service

import (
	"context"
	"github.com/carefreeskyio/example/proto"
)

type Service struct {
	DemoService *DemoService
}

func NewService() *Service {
	return &Service{
		DemoService: NewDemoService(),
	}
}

func (s *Service) Hello (ctx context.Context, request *proto.DemoHelloRequest, response *proto.DemoHelloResponse) (err error) {
	return s.DemoService.Hello(ctx, request, response)
}


