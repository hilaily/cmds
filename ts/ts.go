package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jessevdk/go-flags"
)

var (
	format = "2006-01-02 15:04:05"
)

type args struct {
	Query string `short:"q" description:"timestamp or a time string"`
}

func main() {
	arg := &args{}
	flags.ParseArgs(arg, os.Args)
	res, _ := transfer(arg.Query)
	fmt.Println(res)
}

func transfer(data string) (string, error) {
	if data == "" {
		return fmt.Sprintf("%d", time.Now().Unix()), nil
	}
	num, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		t, err := time.Parse(format, data)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("%d", t.Unix()), nil
	}
	t := time.Unix(num, 0)
	return t.Format(format), nil
}
