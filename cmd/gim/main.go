package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

const (
	debug   = true
	CmdName = "gim"
	Version = "v0.1.0"
)

func main() {
	if !debug {
		log.SetFlags(log.Ltime | log.Lshortfile)
		log.Println("Debug mode on")
	} else {
		log.SetFlags(0)
	}

	if err := newCliApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func newVimCmd() *cli.Command {
	return &cli.Command{
		Name:    "vim",
		Aliases: []string{"v"},
		Usage:   "Run text editor like vim",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "",
				Usage:   "optional config file `conf.yaml`",
			},
			&cli.StringFlag{
				Name:    "message",
				Aliases: []string{"m", "msg"},
				Value:   "",
				Usage:   "path to commit message `FILE`",
			},
		},
		Action: func(cCtx *cli.Context) error {
			fmt.Println("run test command", cCtx.Args().First())
			return nil
			//confFilePath := ctx.String("config")
			//fileInput := ctx.String("message")
			//return test_command(confFilePath, fileInput)
		},
		Subcommands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a new template",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("new task template: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:    "remove",
				Aliases: []string{"r"},
				Usage:   "remove an existing template",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("removed task template: ", cCtx.Args().First())
					return nil
				},
			},
		},
	}
}

func newCliApp() *cli.App {
	cmds := []*cli.Command{
		newVimCmd(),
	}
	flags := newFlags()
	cli.VersionFlag = &cli.BoolFlag{Name: "version", Aliases: []string{"v"}, Usage: "Display this app version information"}
	// https://cli.urfave.org/v2/examples/full-api-example/
	app := &cli.App{
		Name:     CmdName,
		Version:  Version,
		Compiled: time.Now(),
		//Authors:  []*cli.Author{&cli.Author{Name: "xcd0", Email: "xcd0@xcd0.com"}},
		//Copyright:   "(c) 2023 xcd0",
		//UsageText:   "contrive - demonstrating the available API",
		//Description: "no really, there is a lot of dooing to be done",
		//ArgsUsage:   "[arrgh]",
		Commands:             cmds,
		DefaultCommand:       "vim",
		Flags:                flags,
		EnableBashCompletion: true,
		HideHelp:             false,
		HideVersion:          false,
	}
	return app
}

func newFlags() []cli.Flag {
	return []cli.Flag{
		/*
			&cli.BoolFlag{Name: "fancy"},
			&cli.BoolFlag{Value: true, Name: "fancier"},
			&cli.DurationFlag{Name: "howlong", Aliases: []string{"H"}, Value: time.Second * 3},
			&cli.Float64Flag{Name: "howmuch"},
			&cli.GenericFlag{Name: "wat", Value: &genericType{}},
			&cli.Int64Flag{Name: "longdistance"},
			&cli.Int64SliceFlag{Name: "intervals"},
			&cli.IntFlag{Name: "distance"},
			&cli.IntSliceFlag{Name: "times"},
			&cli.StringFlag{Name: "dance-move", Aliases: []string{"d"}},
			&cli.StringSliceFlag{Name: "names", Aliases: []string{"N"}},
			&cli.UintFlag{Name: "age"},
			&cli.Uint64Flag{Name: "bigage"},
		*/
	}
}
