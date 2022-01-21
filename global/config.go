package global

import (
	"flag"
	"fmt"
	"github.com/jinzhu/configor"
	"log"
	"os"
	"path"
	"time"
)

var Config = struct {
	Service struct {
		AppName string
		Network string
		Port    string
	}
	Mysql struct {
		Write string
		Read  string
	}
	Log struct {
		Path         string
		RotationTime time.Duration
		Level        string
		WithMaxAge   int
	}
	Registry struct {
		Addr           string
		BasePath       string
		UpdateInterval time.Duration
		Group          string
	}
	Rpc struct {
		WithTimeout int
	}
	RateLimit struct {
		Enable       bool
		FillInterval time.Duration
		Token        int64
	}
}{}

var DefaultConfigPath = "conf/conf.yaml"

func InitConfig() {
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("get work dir failed: err=%v", err)
	}
	defaultFilePath := path.Join(workDir, DefaultConfigPath)

	configFile := flag.String("c", defaultFilePath, "the config file")

	fmt.Printf("config path: %v\n", *configFile)
	if err = configor.Load(&Config, *configFile); err != nil {
		log.Fatalf("config parse error: err=%v", err)
	}
}
