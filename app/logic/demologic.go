package logic

import (
	repository2 "carefreex/app/repository"
	"context"
	"sync"
)

type DemoLogic struct {
}

var (
	demoLogic     *DemoLogic
	demoLogicOnce sync.Once
)

func NewDemoLogic() *DemoLogic {
	if demoLogic == nil {
		demoLogicOnce.Do(func() {
			demoLogic = &DemoLogic{}
		})
	}

	return demoLogic
}

func (l *DemoLogic) Demo(ctx context.Context) (result string, err error) {
	if result, err = repository2.NewDemoRepository().Demo(ctx); err != nil {
		return result, err
	}

	return result, nil
}
