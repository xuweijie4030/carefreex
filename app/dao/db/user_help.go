package db

import (
	"context"
	"github.com/carefreex-io/logger"
)

func (d *UserDb) GetUserById(ctx context.Context, uid int) (result User, err error) {
	if res := d.DB.Find(&result, uid); res.Error != nil {
		logger.ErrorfX(ctx, "get user by id failed: err=%v", err)
		return result, err
	}

	return result, nil
}
