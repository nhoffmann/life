package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/gosuri/uilive"
	"github.com/nhoffmann/life/pattern"
	"github.com/nhoffmann/life/rle"
	"github.com/nhoffmann/life/simulator"
	"github.com/urfave/cli/v2"
)

func main() {
	var rows int
	var columns int
	var frequency int64
	var rule string

	app := &cli.App{
		Name:  "life",
		Usage: "A game of life simulator",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "rows",
				Aliases:     []string{"r"},
				Usage:       "Number of rows in the grid",
				Value:       20,
				Destination: &rows,
			},
			&cli.IntFlag{
				Name:        "columns",
				Aliases:     []string{"c"},
				Usage:       "Number of columns in the grid",
				Value:       20,
				Destination: &columns,
			},
			&cli.Int64Flag{
				Name:        "frequency",
				Aliases:     []string{"f"},
				Usage:       "The rate at which interations happen in milliseconds",
				Value:       300,
				Destination: &frequency,
			},
			&cli.StringFlag{
				Name:        "rule",
				Aliases:     []string{"u"},
				Usage:       "String notation of the born and survive rules",
				Value:       "B3/S23",
				Destination: &rule,
			},
		},
		Action: func(c *cli.Context) error {
			s := simulator.New(rows, columns, rule)

			writer := uilive.New()
			writer.Start()

			if c.NArg() > 0 {
				dat, err := ioutil.ReadFile(c.Args().Get(0))
				if err != nil {
					log.Fatal(err)
				}
				newRle, err := rle.Parse(string(dat))
				if err != nil {
					log.Fatal(err)
				}
				s.LoadPattern(newRle.Pattern)
				s.Rule = simulator.ParseRule(newRle.Rule)
			} else {
				s.LoadPattern(pattern.Chaos2)
			}

			for {
				fmt.Fprint(writer, s.String())
				time.Sleep(time.Millisecond * time.Duration(frequency))

				s.Evolute()
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
