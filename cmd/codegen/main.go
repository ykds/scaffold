package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "ctl",
		Usage: "代码生成工具",
		Commands: []*cli.Command{
			{
				Name: "gen",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "module",
						Aliases:  []string{"m"},
						Usage:    "模块名",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "target",
						Aliases: []string{"t"},
						Usage:   "输出目录",
						Value:   ".",
					},
				},
				Action: func(c *cli.Context) error {
					return Generate(Config{
						Name:       c.String("module"),
						OutputPath: c.String("target"),
					})
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
