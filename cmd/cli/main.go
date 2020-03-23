package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gosuri/uilive"
	"github.com/nhoffmann/life/simulator"
	"github.com/urfave/cli/v2"
)

func main() {
	var rows int
	var columns int
	var frequency int64

	app := &cli.App{
		Name:  "life",
		Usage: "A game of life simulator",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "rows",
				Aliases:     []string{"r"},
				Usage:       "Number of rows in the grid",
				Value:       9,
				Destination: &rows,
			},
			&cli.IntFlag{
				Name:        "columns",
				Aliases:     []string{"c"},
				Usage:       "Number of columns in the grid",
				Value:       9,
				Destination: &columns,
			},
			&cli.Int64Flag{
				Name:        "frequency",
				Aliases:     []string{"f"},
				Usage:       "The rate at which interations happen in milliseconds",
				Value:       500,
				Destination: &frequency,
			},
		},
		Action: func(c *cli.Context) error {
			s := simulator.New(rows, columns)

			writer := uilive.New()
			writer.Start()

			glider := [][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
			}

			s.LoadPattern(glider, 2, 2)

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
