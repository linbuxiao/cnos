package main

import (
	"fmt"
	cnos "github.com/linbuxiao/cons"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name: "cons",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "from",
				Aliases:  []string{"f"},
				Required: true,
			},
			&cli.IntFlag{
				Name:     "to",
				Aliases:  []string{"t"},
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			from := c.Int("from")
			to := c.Int("to")
			base := c.Args().First()
			cn, err := cnos.NewCnos(from, base)
			if err != nil {
				return err
			}
			result, err := cnos.GetValue(cn, to)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
