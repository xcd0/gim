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

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Display this app version information",
	}
	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Shows a list of commands or help for one command",
	}

	// https://cli.urfave.org/v2/examples/full-api-example/
	app := &cli.App{
		Name:           CmdName,
		Version:        Version,
		Compiled:       time.Now(),
		DefaultCommand: "add",
		//Authors:  []*cli.Author{&cli.Author{Name: "xcd0", Email: "xcd0@xcd0.com"}},
		//Copyright:   "(c) 2023 xcd0",
		//UsageText:   "contrive - demonstrating the available API",
		//Description: "no really, there is a lot of dooing to be done",
		//ArgsUsage:   "[arrgh]",
		Flags: []cli.Flag{
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
		},
		EnableBashCompletion: true,
		HideHelp:             false,
		HideVersion:          false,
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(cCtx *cli.Context) error {
					fmt.Println("added task: ", cCtx.Args().First())
					return nil
				},
			},
			{
				Name:  "short",
				Usage: "complete a task on the list",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "serve", Aliases: []string{"s"}},
					&cli.BoolFlag{Name: "option", Aliases: []string{"o"}},
					&cli.StringFlag{Name: "message", Aliases: []string{"m"}},
				},
				Action: func(cCtx *cli.Context) error {
					fmt.Println("serve:", cCtx.Bool("serve"))
					fmt.Println("option:", cCtx.Bool("option"))
					fmt.Println("message:", cCtx.String("message"))
					return nil
				},
			},
			{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "options for task templates",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("new task template: ", cCtx.Args().First())
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("removed task template: ", cCtx.Args().First())
							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
