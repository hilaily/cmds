package main

import (
	"log"
	"net"
	"os"
)

// 支持三种使用方式
// ppp socks ip:port
// ppp http ip:port
// ppp s2h from_ip:port to_ip:port
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	args := os.Args[1:]
	if len(args) == 0 {
		showhelp()
		return
	}
	var err error
	typ := args[0]
	switch typ {
	case "socks":
		err = runSocks5(args[1])
	case "http":
		err = runHTTP(args[1], &net.Dialer{})
	case "s2h":
		err = s2h(args[1], args[2])
	case "help":
		showhelp()
	default:
		log.Printf("type is not support, %s\n", typ)
	}
	if err != nil {
		log.Println(err.Error())
	}
}

func showhelp() {
	log.Printf(`
just support
ppp socks ip:port
ppp http ip:port
ppp s2h from_ip:port to_ip:port
`)
}
