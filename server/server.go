package server

import (
	"fmt"
	"hk4e-proxy/config"
	"hk4e-proxy/pkg/logger"
	"net/http"

	"github.com/elazarl/goproxy"
)

// ProxyServer 代理服务器
type ProxyServer struct {
	proxy *goproxy.ProxyHttpServer
}

func NewProxyServer() *ProxyServer {
	p := new(ProxyServer)
	p.proxy = goproxy.NewProxyHttpServer()

	p.initProxy()
	return p
}

func (p *ProxyServer) initProxy() {
	// 开启详细输出
	// p.proxy.Verbose = true
	// 捕获https
	p.proxy.OnRequest().HandleConnect(goproxy.AlwaysMitm)
	// 设置代理处理函数
	p.proxy.OnRequest().DoFunc(p.redirect)
	go func() {
		addr := fmt.Sprintf(":%v", config.GetConfig().Proxy.ProxyPort)
		logger.Error("listen and serve error, err: %v", http.ListenAndServe(addr, p.proxy))
	}()
	logger.Info("proxy listen, port: %v", config.GetConfig().Proxy.ProxyPort)
}
