package server

import (
	"context"
	"net"

	"github.com/upkit/grpc_learn/app/server/dao"
	"github.com/upkit/grpc_learn/libs/conf"
	"github.com/upkit/grpc_learn/libs/log"
	"github.com/upkit/grpc_learn/proto/person/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Service struct {
	srv *grpc.Server
	dao *dao.Dao
}

func New(cfg *conf.Config) {
	srv := &Service{
		srv: grpc.NewServer(),
		dao: dao.New(context.Background(), cfg),
	}

	listen, err := net.Listen("tcp", cfg.GrpcAddr)
	if err != nil {
		log.Fatal("net.Listen failed", zap.String("addr", cfg.GrpcAddr))
	}
	person.RegisterPersonServiceServer(srv.srv, srv.dao)

	srv.SignalMonitor()
	srv.srv.Serve(listen)
}

func (srv *Service) Close() {
	srv.srv.GracefulStop()
	srv.dao.Close()
}
