# PPP

## Install

```shell
go install github.com/hilaily/cmds/ppp@latest
```

## Usage

1 Start a http http proxy server 

```bash
ppp http {ip:port}
```

2 Start a socks5 proxy server  

```shell
ppp socks {ip:port}  
```

3 Transfer a socks5 proxy to http proxy 

```shell
ppp s2h {socks5_ip:port} {http_ip:port}  
```

