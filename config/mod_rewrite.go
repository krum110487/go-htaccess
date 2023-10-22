package config

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

type ModRewrite struct {
	EngineOn bool
}

func replaceVariables(str string, request *http.Request) string {
	now := time.Now()
	variables := map[string]string{
		"%{HTTP_ACCEPT}":           request.Header.Get("Accept"),
		"%{HTTP_COOKIE}":           request.Header.Get("Cookie"),
		"%{HTTP_FORWARDED}":        request.Header.Get("Forwarded"),
		"%{HTTP_HOST}":             request.Host,
		"%{HTTP_PROXY_CONNECTION}": request.Header.Get("Proxy-Connection"),
		"%{HTTP_REFERER}":          request.Header.Get("Referer"),
		"%{HTTP_USER_AGENT}":       request.Header.Get("User-Agent"),
		"%{AUTH_TYPE}":             request.Header.Get("Authorization"),
		"%{CONN_REMOTE_ADDR}":      request.RemoteAddr,
		"%{CONTEXT_PREFIX}":        "%{CONTEXT_PREFIX_NOT_IMPLEMENTED}",
		"%{CONTEXT_DOCUMENT_ROOT}": "%{CONTEXT_DOCUMENT_ROOT_NOT_IMPLEMENTED}",
		"%{IPV6}":                  "%{IPV6_NOT_IMPLEMENTED}",
		"%{PATH_INFO}":             request.URL.Path,
		"%{QUERY_STRING}":          request.URL.RawQuery,
		"%{REMOTE_ADDR}":           request.RemoteAddr,
		"%{REMOTE_HOST}":           "%{REMOTE_HOST_NOT_IMPLEMENTED}",
		"%{REMOTE_IDENT}":          "%{REMOTE_IDENT_NOT_IMPLEMENTED}",
		"%{REMOTE_PORT}":           request.URL.Port(),
		"%{REMOTE_USER}":           "%{REMOTE_USER_NOT_IMPLEMENTED}",
		"%{REQUEST_METHOD}":        request.Method,
		"%{SCRIPT_FILENAME}":       "%{SCRIPT_FILENAME_NOT_IMPLEMENTED}",
		"%{SERVER_ADMIN}":          "%{SERVER_ADMIN_NOT_IMPLEMENTED}",
		"%{SERVER_NAME}":           request.Host,
		"%{SERVER_PORT}":           request.URL.Port(),
		"%{SERVER_PROTOCOL}":       request.Proto,
		"%{SERVER_SOFTWARE}":       "%{SERVER_SOFTWARE_NOT_IMPLEMENTED}",
		"%{TIME_DAY}":              fmt.Sprintf("%02d", now.Day()),
		"%{TIME_HOUR}":             fmt.Sprintf("%02d", now.Hour()),
		"%{TIME_MIN}":              fmt.Sprintf("%02d", now.Minute()),
		"%{TIME_MON}":              fmt.Sprintf("%02d", now.Month()),
		"%{TIME_SEC}":              fmt.Sprintf("%02d", now.Second()),
		"%{TIME_YEAR}":             fmt.Sprintf("%04d", now.Year()),
		"%{TIME_WDAY}":             fmt.Sprintf("%d", now.Weekday()),
		"%{API_VERSION}":           "%{API_VERSION_NOT_IMPLEMENTED}",
		"%{HTTPS}":                 "%{HTTPS_NOT_IMPLEMENTED}",
		"%{IS_SUBREQ}":             "%{IS_SUBREQ_NOT_IMPLEMENTED}",
		"%{REQUEST_FILENAME}":      request.URL.Path,
		"%{REQUEST_SCHEME}":        "%{REQUEST_SCHEME_NOT_IMPLEMENTED}",
		"%{REQUEST_URI}":           request.URL.RequestURI(),
		"%{THE_REQUEST}":           "%{THE_REQUEST_NOT_IMPLEMENTED}",
	}
	if host, _, err := net.SplitHostPort(request.RemoteAddr); err == nil {
		variables["%{REMOTE_HOST}"] = host
	}
	for variable, value := range variables {
		str = strings.ReplaceAll(str, variable, value)
	}
	return str
}
