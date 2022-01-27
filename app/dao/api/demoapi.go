package api

import (
	"context"
	"sync"
)

type DemoApi struct {
}

var (
	demoApi     *DemoApi
	demoApiOnce sync.Once
)

func NewDemoApi() *DemoApi {
	if demoApi == nil {
		demoApiOnce.Do(func() {
			demoApi = &DemoApi{}
		})
	}

	return demoApi
}

func (l *DemoApi) Demo(ctx context.Context) (result string, err error) {
	result = "this is api demo"

	return result, nil
}
