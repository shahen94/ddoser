package internal

import (
	"os"

	"github.com/shahen94/ddoser/pkg/domain"
	uCli "github.com/urfave/cli/v2"
)

// Cli is an interface for CLI
type cli struct{}

// Parse parses CLI params
func (c *cli) Parse() (*domain.CliParams, error) {
	params := &domain.CliParams{}

	app := uCli.App{
		Name:  "ddoser",
		Usage: "DDoS tool",
		Flags: []uCli.Flag{
			&uCli.StringFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				Usage:    "URL to attack",
				Required: true,
			},
			&uCli.IntFlag{
				Name:    "concurrent",
				Aliases: []string{"c"},
				Usage:   "Number of concurrent workers",
				Value:   1,
			},
			&uCli.IntFlag{
				Name:    "requests",
				Aliases: []string{"r"},
				Usage:   "Number of requests",
				Value:   100,
			},
			&uCli.IntFlag{
				Name:    "timeout",
				Aliases: []string{"t"},
				Usage:   "Timeout in milliseconds",
				Value:   1000,
			},
			&uCli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "Verbose requests",
				Value:   false,
			},
		},
		Action: func(c *uCli.Context) error {
			params.BaseUrl = c.String("url")
			params.Concurrent = c.Int("concurrent")
			params.RequestCount = c.Int("requests")
			params.Timeout = c.Int("timeout")
			params.Verbose = c.Bool("verbose")

			return nil
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		return nil, domain.ErrParsingCliParams
	}

	return params, nil
}

// NewCli creates a new CLI
func NewCli() domain.Cli {
	return &cli{}
}
