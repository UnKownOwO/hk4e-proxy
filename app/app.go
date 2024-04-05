package app

import (
	"context"
	"hk4e-proxy/config"
	"hk4e-proxy/pkg/logger"
	"hk4e-proxy/server"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context, configFile string) error {
	config.InitConfig(configFile)

	logger.InitLogger("proxy")
	logger.Warn("proxy start")
	defer func() {
		logger.CloseLogger()
	}()

	_ = server.NewProxyServer()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case <-ctx.Done():
			return nil
		case s := <-c:
			logger.Warn("get a signal %s", s.String())
			switch s {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				logger.Warn("proxy exit")
				return nil
			case syscall.SIGHUP:
			default:
				return nil
			}
		}
	}
}
