package main

//import (
//	"fmt"
//	"net"
//	"net/http"
//	"net/http/httputil"
//	"net/url"
//	"time"
//)
//
//func main() {
//
//	var defaultTransport http.RoundTripper = &http.Transport{
//		Proxy: nil,
//		DialContext: (&net.Dialer{
//			Timeout:   10 * time.Second,
//			KeepAlive: 30 * time.Second,
//		}).DialContext,
//		MaxIdleConns:           1000000,
//		MaxIdleConnsPerHost:    1000000,
//		MaxConnsPerHost:        1000000,
//		IdleConnTimeout:        90 * time.Second,
//	}
//
//
//	url, err := url.Parse("http://*")
//
//	if err != nil {
//		println(fmt.Sprintf("Error parsing url : %#v", err))
//	}
//	proxy := httputil.ReverseProxy{
//		Director: func(req *http.Request) {
//			req.Host = "*"
//			req.URL = url
//		},
//		Transport:      defaultTransport,
//		FlushInterval:  -1,
//	}
//
//	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
//		proxy.ServeHTTP(w, r)
//	})
//
//	server := http.Server{
//		Addr: "0.0.0.0:7777",
//	}
//	server.ListenAndServe()
//
//}

import (
	"github.com/valyala/fasthttp"
	"log"
)

var proxyClient = &fasthttp.HostClient{
	Addr: "*",
	// set other options here if required - most notably timeouts.
}

func ReverseProxyHandler(ctx *fasthttp.RequestCtx) {
	req := &ctx.Request
	resp := &ctx.Response
	prepareRequest(req)
	if err := proxyClient.Do(req, resp); err != nil {
		ctx.Logger().Printf("error when proxying the request: %s", err)
	}
	postprocessResponse(resp)
}

func prepareRequest(req *fasthttp.Request) {
	// do not proxy "Connection" header.
	req.Header.Del("Connection")
	req.SetHost("*")
	// strip other unneeded headers.

	// alter other request params before sending them to upstream host
}

func postprocessResponse(resp *fasthttp.Response) {
	// do not proxy "Connection" header
	resp.Header.Del("Connection")
	resp.Header.Set("proxy", "fasthttp")
	// strip other unneeded headers

	// alter other response data if needed
}

func main() {
	if err := fasthttp.ListenAndServe(":8080", ReverseProxyHandler); err != nil {
		log.Fatalf("error in fasthttp server: %s", err)
	}
}