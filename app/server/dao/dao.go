package dao

import (
	"context"

	"github.com/upkit/grpc_learn/libs/cache/mredis"
	"github.com/upkit/grpc_learn/libs/conf"
	"github.com/upkit/grpc_learn/libs/db/msql"
	"github.com/upkit/grpc_learn/libs/xerror"
)

type Dao struct {
	c  conf.Config
	dc *msql.DB
	cc *mredis.Conn

	ctx    context.Context
	cancel context.CancelFunc
}

func New(ctx context.Context, c *conf.Config) *Dao {
	d := Dao{
		c:  *c,
		dc: msql.NewDB(c.Mysql),
		cc: mredis.New(c.Redis),
	}
	d.ctx, d.cancel = context.WithCancel(ctx)

	d.Ping(ctx)
	return &d
}

func (d *Dao) Ping(ctx context.Context) error {
	var err error
	xerror.CheckError(&err, func() error {
		return d.dc.Ping()
	})
	xerror.CheckError(&err, func() error {
		return d.cc.Ping().Err()
	})
	return err
}

func (d *Dao) Close() {
	d.cancel()
	_ = d.dc.Close()
	_ = d.cc.Close()
}
