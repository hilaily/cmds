## ServeSwagger

Render a swagger file in local with http server.

### Install

```bash
go install github.com/hilaily/cmds/serveswagger@latest
```

### Usage

```shell
serveswagger
```

It will find swagger.json or swagger.yaml in current directory automatically and serve it at localhost:8123

find order

`./swagger.json`, `./swagger.yaml`, `./docs/swagger.json`, `./docs/swagger.json` 

-----

```shell
serveswagger -p 8080 -f myswagger.json
```

It will render myswagger.json at localhost:8080

