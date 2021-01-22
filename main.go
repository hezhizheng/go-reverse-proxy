package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
)

func main() {
	port := flag.String("p", "1874", "本地监听的端口")
	reverseUrl := flag.String("r", "https://hzz.cool", "需要代理的地址")
	flag.Parse()

	remote, err := url.Parse(*reverseUrl)
	if err != nil {
		panic(err)
	}

	proxy := GoReverseProxy(&RProxy{
		remote: remote,
	})

	log.Println("当前代理地址： " + *reverseUrl + " 本地监听： http://127.0.0.1:" + *port)

	serveErr := http.ListenAndServe(":"+*port, proxy)

	if serveErr != nil {
		panic(serveErr)
	}
}