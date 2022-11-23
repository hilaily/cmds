package main

import (
	"log"

	"github.com/armon/go-socks5"
)

func runSocks5(addrStr string) error {
	// Create a SOCKS5 server
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		return err
	}

	// Create SOCKS5 proxy on localhost port 8000
	log.Println("start socks5 server: ", addrStr)
	if err := server.ListenAndServe("tcp", addrStr); err != nil {
		return err
	}
	return nil
}
