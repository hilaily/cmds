package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hilaily/kit/pathx"
	"github.com/jessevdk/go-flags"
	"github.com/pkg/browser"
)

var (
	//go:embed dist/*
	dist        embed.FS
	baseURL     = "http://127.0.0.1:"
	port        = "8123"
	customIdent = "/myswagger"
)

func main() {
	log.SetFlags(0)

	op := &ops{}
	_, err := flags.ParseArgs(op, os.Args)
	if err != nil {
		log.Println(err.Error())
		return
	}
	file, err := checkFile(op.File)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Printf("serve file: %s, port: %d\n", file, op.Port)
	port = fmt.Sprintf("%d", op.Port)
	baseURL = baseURL + port
	render(file)
}

type ops struct {
	Port int    `short:"p" long:"port" description:"HTTP server port" default:"8123"`
	File string `short:"f" long:"file" description:"A path of swagger file"`
}

func checkFile(file string) (string, error) {
	if file != "" {
		return file, nil
	}
	for _, v := range []string{"./swagger.json", "./swagger.yaml", "./docs/swagger.json", "./docs/swagger.yaml"} {
		if pathx.IsExist(v) {
			return v, nil
		}
	}
	if file == "" {
		return "", fmt.Errorf("you should use it like 'serveswagger -f /opt/swagger.json'")
	}
	return file, nil
}

func render(file string) {
	sub, _ := fs.Sub(dist, "dist")

	files := http.FS(sub)
	u := file
	h := http.FileServer(files)
	if !strings.HasPrefix(file, "http") {
		h = &mfs{
			Handler: http.FileServer(files),
			path:    file,
		}
		u = baseURL + customIdent
	}
	http.Handle("/", h)
	browser.OpenURL(baseURL + "?file=" + u)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

type mfs struct {
	//http.FileSystem
	http.Handler
	path string
}

/*
func (f *mfs) Open(name string) (http.File, error) {
	if name == "/swagger.json" {
		en, _ := os.ReadFile(f.path)
		log.Println(string(en))
		return os.OpenFile(f.path, os.O_RDONLY, 0777)
	}
	return f.FileSystem.Open(name)
}
*/

func (f *mfs) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	name := req.URL.Path
	if name == customIdent {
		en, _ := os.ReadFile(f.path)
		resp.WriteHeader(200)
		resp.Write(en)
		return
	}
	f.Handler.ServeHTTP(resp, req)
}
