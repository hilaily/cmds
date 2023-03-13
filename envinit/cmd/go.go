package cmd

import (
	"github.com/hilaily/cmds/envinit/util"
	"github.com/urfave/cli/v2"
)

var (
	GoCMD = (&goCMD{name: "go"}).cmd()
)

type goCMD struct {
	name string
}

func (g *goCMD) cmd() *cli.Command {
	return &cli.Command{
		Name:  g.name,
		Usage: "install " + g.name,
		Subcommands: []*cli.Command{
			{Name: "install", Action: g.install},
		},
	}
}

func (g *goCMD) install(ctx *cli.Context) error {
	util.BrewInstall(g.name)
	util.BrewInstall("gopls")
	util.BrewInstall("golanglint-ci")
	return nil
}

/*
func (g *goCMD) rawInstall(){
	        self.brew_install("go")
	        cwd = os.getcwd()
	        os.chdir("/usr/local")
	        if os.path.exists("/usr/local/opt/go/libexec"):
	            run("sudo mv go go_old")
	            run("sudo ln -s /usr/local/opt/go/libexec /usr/local/go")
	        os.chdir(cwd)
	        self.brew_install("gopls")
	        self.brew_install("golanglint-ci")
	return nil
}
*/
