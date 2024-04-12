package netutil

import (
	"net"
	"net/http"
)

const (
	xForwardedFor = "X-Forwarded-For"
	xRealIP       = "X-Real-IP"
)

// RemoteIP 获取请求的远程IP地址
// req: http请求
//
// return: 远程IP地址
func RemoteIP(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(xRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(xForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}
