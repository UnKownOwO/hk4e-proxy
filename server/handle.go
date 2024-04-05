package server

import (
	"hk4e-proxy/config"
	"hk4e-proxy/pkg/logger"
	"net/http"
	"net/url"
	"strings"

	"github.com/elazarl/goproxy"
)

// redirect 重定向
func (p *ProxyServer) redirect(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	// 域名符合则重定向
	if strings.Contains(req.URL.Host, "yuanshen.com") ||
		strings.Contains(req.URL.Host, "mihoyo.com") ||
		strings.Contains(req.URL.Host, "hoyoverse.com") {

		// 将要重定向的url
		redirectUrlStr := ""
		if strings.Contains(req.URL.Path, "query_security_file") ||
			strings.Contains(req.URL.Path, "query_region_list") ||
			strings.Contains(req.URL.Path, "query_cur_region") {
			// dispatch
			redirectUrlStr = config.GetConfig().Proxy.RedirectDispatch
		} else {
			// sdk
			redirectUrlStr = config.GetConfig().Proxy.RedirectSdk
		}
		redirectUrl, err := url.Parse(redirectUrlStr)
		if err != nil {
			logger.Error("url parse error, err: %v", redirectUrl)
			return req, nil
		}
		logger.Info("redirect, method: %v, %v -> %v", req.Method, req.URL.String(), redirectUrlStr)
		req.URL.Scheme = redirectUrl.Scheme
		req.URL.Host = redirectUrl.Host
	}
	return req, nil
}
