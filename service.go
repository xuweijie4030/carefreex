package main

import (
	"carefreex/app/service"
	"context"
	"github.com/carefreex-io/example/proto"
)

type Service struct {
	DemoService *service.DemoService
}

func NewService() *Service {
	return &Service{
		DemoService: service.NewDemoService(),
	}
}

func (s *Service) Hello (ctx context.Context, request *proto.DemoHelloRequest, response *proto.DemoHelloResponse) (err error) {
	return s.DemoService.Hello(ctx, request, response)
}


