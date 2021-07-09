/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/7/8 上午10:29
 * @note: http(s) 哑代理
 * @refer: https://mojotv.cn/2018/12/26/how-to-create-a-https-proxy-service-in-100-lines-of-code
 */

package main

import (
	"crypto/tls"
	"github.com/magiclyde/http-proxy/internal"
	"log"
	"net/http"
	"strconv"
)

var (
	BuildDate    string
	BuildVersion string
)

func init() {
	log.Printf("[main] Git commit:%s\n", BuildVersion)
	log.Printf("[main] Build time:%s\n", BuildDate)
}

func main() {
	cfg := internal.NewConfig()
	log.Printf("[main] cfg: %+v", cfg)

	server := &http.Server{
		Addr:         ":" + strconv.Itoa(cfg.Port),
		Handler:      &internal.Proxy{},
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)), // 关闭 http2
	}
	log.Printf("[main] ServeHttp on :%d\n", cfg.Port)

	switch cfg.Proto {
	case "http":
		log.Fatalf("[main] ListenAndServe.err: %+v", server.ListenAndServe())

	case "https":
		log.Fatalf("[main] ListenAndServeTLS.err: %+v", server.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile))

	default:
		log.Fatal("[main] Protocol must be either http or https")
	}

}
