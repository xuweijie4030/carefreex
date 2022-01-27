package service

import (
	"carefreex/app/dao/db"
	"context"
	"fmt"
	"github.com/carefreex-io/config"
	"github.com/carefreex-io/dbdao/gormdb"
	"github.com/carefreex-io/logger"
	"testing"
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

func TestDemoService_Hello(t *testing.T) {
	config.DefaultCustomOptions.Path = "../../conf/conf.yaml"
	config.InitConfig()

	logger.AddCtxFields([]string{"test"})
	logger.InitLogger()

	ctx := context.Background()
	ctx = context.WithValue(ctx, "test", "testTest")
	if err := gormdb.InitDB(); err != nil {
		logger.ErrorfX(ctx, "mysql.InitDB failed: err=%v", err)
	}

	//user := db.User{
	//	Email: "xxx@xx.com",
	//	TTime: time.Now(),
	//}
	//
	//if res := db.NewUserDb(ctx, true).DB.Create(&user); res.Error != nil {
	//	logger.Fatalf("create user failed: err=%v", res.Error)
	//}

	user, err := db.NewUserDb(ctx).GetUserById(ctx, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user)
	//
	//request := proto.DemoHelloRequest{
	//	Name: "zhangSan",
	//}
	//response := proto.DemoHelloResponse{}
	//
	//NewDemoService().Hello(ctx, &request, &response)
}
