package conf

import (
	"errors"
	"flag"
	"os"

	"github.com/upkit/grpc_learn/libs/cache/mredis"
	"github.com/upkit/grpc_learn/libs/db/msql"
	"github.com/upkit/grpc_learn/libs/log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Debug    bool        `yaml:"debug"`     // 是否测试环境
	GrpcAddr string      `yaml:"grpc_addr"` // grpc地址
	Mysql    msql.Conf   `yaml:"mysql"`     // mysql配置
	Redis    mredis.Conf `yaml:"redis"`     // redis配置
	Log      log.Conf    `yaml:"log"`       // log配置
}

var (
	conf     Config
	confPath string
)

func init() {
	flag.StringVar(&confPath, "c", "", "config path")
}

func Init() (*Config, error) {
	if confPath == "" {
		return nil, errors.New("not load config file")
	}
	file, err := os.Open(confPath)
	if err != nil {
		return nil, err
	}
	err = yaml.NewDecoder(file).Decode(&conf)
	return &conf, err
}
