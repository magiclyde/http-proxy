/**
 * Created by GoLand.
 * @author: clyde
 * @date: 2021/7/9 下午5:36
 * @note:
 */

package internal

import (
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

type Proxy struct {
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodConnect {
		handleHttp(w, r)
	} else {
		handleTunneling(w, r)
	}
}

func handleHttp(w http.ResponseWriter, r *http.Request) {
	// step1: 收到客户端的请求，复制原来的请求对象
	outReq := new(http.Request)
	*outReq = *r

	// step 2: 新请求发送到服务器端，并接收到服务器端返回的响应
	transport := http.DefaultTransport
	resp, err := transport.RoundTrip(outReq)
	if err != nil {
		log.Printf("[proxy] transport.RoundTrip err: %+v", err)
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	// step 3: 对响应做一些处理，然后返回给客户端
	for key, value := range resp.Header {
		for _, v := range value {
			w.Header().Add(key, v)
		}
	}

	// 返回状态码
	w.WriteHeader(resp.StatusCode)

	// 返回 body, zero copy would be better, see https://github.com/funny/proxy#零拷贝技术 , todo...
	io.Copy(w, resp.Body)
}

func handleTunneling(w http.ResponseWriter, r *http.Request) {
	//设置超时防止大量超时导致服务器资源不大量占用
	dest_conn, err := net.DialTimeout("tcp", r.Host, 10*time.Second)
	if err != nil {
		log.Printf("[proxy] net.DialTimeout err: %+v", err)
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	w.WriteHeader(http.StatusOK)

	//类型转换
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		log.Println("[proxy] Hijacking not supported")
		http.Error(w, "Hijacking not supported", http.StatusInternalServerError)
		return
	}

	//接管连接
	client_conn, _, err := hijacker.Hijack()
	if err != nil {
		log.Printf("[proxy] hijacker.Hijack err: %+v", err)
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	go transfer(dest_conn, client_conn)
	go transfer(client_conn, dest_conn)
}

func transfer(dst io.WriteCloser, src io.ReadCloser) {
	defer dst.Close()
	defer src.Close()
	io.Copy(dst, src)
}
