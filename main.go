package main

import (
	"os"
	"path/filepath"

	"github.com/codegangsta/cli"
	"github.com/gobuild/gobuild2/cmd/pack"
	"github.com/gobuild/gobuild2/cmd/runinit"
	"github.com/gobuild/gobuild2/cmd/slave"
	"github.com/gobuild/gobuild2/cmd/web"
	"github.com/gobuild/log"
)

const VERSION = "0.0.1.0607"

var app = cli.NewApp()

func init() {
	cwd, _ := os.Getwd()
	program := filepath.Base(cwd)

	app.Name = "gobuild"
	app.Usage = "[COMMANDS]"
	app.Version = VERSION
	app.Commands = append(app.Commands,
		cli.Command{
			Name:   "slave",
			Usage:  "start gobuild compile slave",
			Action: slave.Action,
			Flags: []cli.Flag{
				cli.StringFlag{"webaddr,w", "localhost:8010", "gobuild2 web address"},
			},
		},
		cli.Command{
			Name:   "init",
			Usage:  "initial gobuild.yml file",
			Action: runinit.Action,
		},
		cli.Command{
			Name:   "pack",
			Usage:  "build and pack file into tgz or zip",
			Action: pack.Action,
			Flags: []cli.Flag{
				cli.StringFlag{"os", os.Getenv("GOOS"), "operation system"},
				cli.StringFlag{"arch", os.Getenv("GOARCH"), "arch"},
				cli.StringFlag{"depth", "3", "depth of file to walk"},
				cli.StringFlag{"output,o", program + ".zip", "target file"},
				cli.StringFlag{"gom", "go", "go package manage program"},
				cli.BoolFlag{"nobuild", "donot call go build when pack"},
				cli.StringSliceFlag{"add,a", &cli.StringSlice{}, "add file"},
			},
		},
		cli.Command{
			Name:   "web",
			Usage:  "start gobuild web server",
			Action: web.Action,
			Flags: []cli.Flag{
				cli.StringFlag{"conf,f", "conf/app.ini", "config file"},
			},
		},
	)
}

func main() {
	log.SetOutputLevel(log.Ldebug)
	app.Run(os.Args)
}
