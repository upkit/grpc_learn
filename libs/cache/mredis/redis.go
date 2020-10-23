package mredis

import (
	"github.com/go-redis/redis"
)

type Conn struct {
	*redis.Client
}

type Conf struct {
	Addr       string `yaml:"addr"`        // 连接地址
	Pwd        string `yaml:"pwd"`         // 密码
	DB         int    `yaml:"db"`          // db号
	Pool       int    `yaml:"pool"`        // 连接池
	MaxRetries int    `yaml:"max_retries"` // 失败重试次数
}

func New(cfg Conf) *Conn {
	client := redis.NewClient(&redis.Options{
		Addr:       cfg.Addr,
		Password:   cfg.Pwd,
		PoolSize:   cfg.Pool,
		MaxRetries: cfg.MaxRetries,
		DB:         cfg.DB,
	})
	return &Conn{Client: client}
}
