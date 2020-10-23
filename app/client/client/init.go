package client

import (
	"context"

	"github.com/upkit/grpc_learn/libs/conf"
	"github.com/upkit/grpc_learn/libs/log"
	"github.com/upkit/grpc_learn/proto/person/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Client struct {
	c  conf.Config
	cn *grpc.ClientConn
	cc person.PersonServiceClient

	ctx    context.Context
	cancel context.CancelFunc
}

func New(cfg *conf.Config) *Client {
	cli := Client{
		c: *cfg,
	}
	cli.ctx, cli.cancel = context.WithCancel(context.Background())

	var err error
	cli.cn, err = grpc.DialContext(cli.ctx, cfg.GrpcAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("grpc.DialContext failed", zap.String("addr", cfg.GrpcAddr))
	}

	cli.cc = person.NewPersonServiceClient(cli.cn)
	return &cli
}

func (c *Client) Close() {
	c.cancel()
	c.cn.Close()
}
