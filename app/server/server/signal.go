package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/upkit/grpc_learn/libs/log"
	"go.uber.org/zap"
)

func (srv *Service) SignalMonitor() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		for s := range c {
			log.Info("received a system signal", zap.String("signal", s.String()))
			switch s {
			case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
				// 退出前的一些扫尾工作
				srv.Close()
				fallthrough
			case syscall.SIGHUP:
			default:
				return
			}
		}
	}()
}
