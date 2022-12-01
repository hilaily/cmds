package main

func main() {

}

/*
	1. in pkg/mod
		rm go/pkg/mod/github.com/test/pkgname@v1.0.1/
    2. in pkg/mod/cache/download
		rm go/pkg/mod/cache/download/github.com/pkg/testname/@v/v0.0.0-20221121103753-fdd8a8e680aa.lock
    3. in pkg/mod/cache/vcs
*/
