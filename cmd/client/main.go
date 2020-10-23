package main

import (
	"flag"

	"github.com/upkit/grpc_learn/app/client/client"
	"github.com/upkit/grpc_learn/libs/conf"
	"github.com/upkit/grpc_learn/libs/log"
)

func main() {
	flag.Parse()
	cfg, err := conf.Init()
	if err != nil {
		panic(err)
	}

	log.Init(cfg.Log)
	defer log.Close()

	log.Info("grpc client started")
	client.New(cfg)
}
