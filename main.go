package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {

	var defaultTransport http.RoundTripper = &http.Transport{
		Proxy: nil,
		DialContext: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:           1000000,
		MaxIdleConnsPerHost:    1000000,
		MaxConnsPerHost:        1000000,
		IdleConnTimeout:        90 * time.Second,
	}


	url, err := url.Parse("http://*")

	if err != nil {
		println(fmt.Sprintf("Error parsing url : %#v", err))
	}
	proxy := httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.Host = "*"
			req.URL = url
		},
		Transport:      defaultTransport,
		FlushInterval:  -1,
	}

	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	server := http.Server{
		Addr: "0.0.0.0:7777",
	}
	server.ListenAndServe()

}
