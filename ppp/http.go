package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"

	"golang.org/x/net/proxy"
)

func runHTTP(addr string, dialer proxy.Dialer) error {
	log.Println("start http server: ", addr)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	for {
		client, err := l.Accept()
		if err != nil {
			return err
		}

		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Println("get panic: ", r)
				}
			}()
			handleClientRequest(client, dialer)
		}()
	}
}

func handleClientRequest(client net.Conn, dialer proxy.Dialer) {
	if client == nil {
		return
	}
	defer client.Close()

	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}
	var method, host, address string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	hostPortURL, err := url.Parse(host)
	if err != nil {
		log.Println(err)
		return
	}

	if hostPortURL.Opaque == "443" { //https访问
		address = hostPortURL.Scheme + ":443"
	} else { //http访问
		if strings.Index(hostPortURL.Host, ":") == -1 { //host不带端口， 默认80
			address = hostPortURL.Host + ":80"
		} else {
			address = hostPortURL.Host
		}
	}

	//获得了请求的host和port，就开始拨号吧
	server, err := dialer.Dial("tcp", address)
	if err != nil {
		log.Println(err)
		return
	}
	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		server.Write(b[:n])
	}
	//进行转发
	go io.Copy(server, client)
	io.Copy(client, server)
}
