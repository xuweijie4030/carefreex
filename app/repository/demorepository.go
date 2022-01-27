package repository

import (
	api2 "carefreex/app/dao/api"
	cache2 "carefreex/app/dao/cache"
	rpc2 "carefreex/app/dao/rpc"
	"context"
	"sync"
)

type DemoRepository struct {
}

var (
	demoRepository     *DemoRepository
	demoRepositoryOnce sync.Once
)

func NewDemoRepository() *DemoRepository {
	if demoRepository == nil {
		demoRepositoryOnce.Do(func() {
			demoRepository = &DemoRepository{}
		})
	}

	return demoRepository
}

func (l *DemoRepository) Demo(ctx context.Context) (result string, err error) {
	if result, err = api2.NewDemoApi().Demo(ctx); err != nil {
		return result, err
	}

	if result, err = cache2.NewDemoCache().Demo(ctx); err != nil {
		return result, err
	}

	if result, err = rpc2.NewDemoRpc().Demo(ctx); err != nil {
		return result, err
	}

	return result, nil
}
