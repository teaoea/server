package tools

import (
	"net"
	"net/http"
	"strings"
)

func ClientIp(r *http.Request) string {
	XForwardFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(XForwardFor, ",")[0])
	if ip != "" {
		return ip
	}
	XRealIp := r.Header.Get("X-Real-Ip")
	ip = strings.TrimSpace(XRealIp)
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return "unable to get client ip"
}
