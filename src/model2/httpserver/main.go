package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

//接收客户端 request，并将 request 中带的 header 写入 response header
//读取当前系统的环境变量中的 VERSION 配置，并写入 response header
//Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
//当访问 localhost/healthz 时，应返回 200

func requestClient(w http.ResponseWriter, r *http.Request) {
	//02 设置env
	os.Setenv("VERSION", "v1.0.0-cncp")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("os env is %s\n", version)
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Printf("Header key is %s\n,Header value is %s\n", k, vv)
			// 并将 request 中带的 header 写入 response header
			w.Header().Set(k, vv)
		}
		clientIP := getUserIP(r)
		log.Printf("reques IP is %s\n", clientIP)
		log.Printf("request statusCode %d\n", http.StatusOK)
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "working")
}

func getUserIP(r *http.Request) string {
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.HandleFunc("/", requestClient)
	mux.HandleFunc("/healthz", healthz)

	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("start http server failed %s\n", err.Error())
	}
}
