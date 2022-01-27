package cache

import (
	"context"
	"sync"
)

type DemoCache struct {
}

var (
	demoCache     *DemoCache
	demoCacheOnce sync.Once
)

func NewDemoCache() *DemoCache {
	if demoCache == nil {
		demoCacheOnce.Do(func() {
			demoCache = &DemoCache{}
		})
	}

	return demoCache
}

func (l *DemoCache) Demo(ctx context.Context) (result string, err error) {
	result = "this is cache demo, you can link redis/memeCache/mangodb/pika here"

	return result, nil
}
