package db

import (
	"context"
	"github.com/carefreex-io/dbdao/gormdb"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id      int       `gorm:"primaryKey" json:"id"`
	Email   string    `json:"email"`
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	TFloat  float32   `json:"t_float"`
	TDouble float64   `json:"t_double"`
	TTime   time.Time `json:"t_time"`
}

type UserDb struct {
	DB *gorm.DB
}

func NewUserDb(ctx context.Context, arg ...bool) (db *UserDb) {
	db = &UserDb{
		DB: gormdb.Read,
	}
	if len(arg) != 0 && arg[0] {
		db.DB = gormdb.Write
	}
	db.DB.WithContext(ctx)

	return db
}

func (d *UserDb) TableName() string {
	return "user"
}

