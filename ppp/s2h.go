package main

import (
	"fmt"
	"log"

	"golang.org/x/net/proxy"
)

func s2h(socksAddr, httpAddr string) error {
	log.Printf("start server, socks5 %s => http %s", socksAddr, httpAddr)
	dialer, err := proxy.SOCKS5("tcp", socksAddr, nil, proxy.Direct)
	if err != nil {
		return fmt.Errorf("can't connect to the proxy: %s", err.Error())
	}
	return runHTTP(httpAddr, dialer)
}
