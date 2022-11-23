package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

/*
命令示例

1. 发送主题为 test 的邮件给当前邮箱
mail2 -s test aa@aa.com
2. 使用管道命令
echo “mail content”| mail2 -s test aa@aa.com
3. 读取文件中的内容发送邮件
mail2 -s test aa@aa.com< file
4. 给多个用户发送邮件
mail2 -s test -c admin@aispider.com  root@aispider.com< file
*/

var (
	// subject  = ""
	attaches = []string{}
)

func main() {
	app := cli.NewApp()
	app.Name = "mail2"
	app.Usage = "Send email with stmp"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "s",
			Usage: "subject for mail",
			Value: "",
		},
		cli.StringFlag{
			Name:  "m",
			Usage: "msg for mail",
			Value: "",
		},
		cli.StringSliceFlag{
			Name:  "a",
			Usage: "attach",
			Value: nil,
		},
		cli.StringFlag{
			Name:  "config",
			Usage: "config file path",
			Value: "",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "init config file",
			Action: func(c *cli.Context) error {
				initConfigFile()
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		if len(os.Args) == 1 {
			return cli.ShowAppHelp(c)
		}
		SendMail(c.String("config"), c.String("s"), c.String("m"), c.Args(), c.StringSlice("a"))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
