package main

import (
	"github.com/copperium/fractals/cli"
	"github.com/integrii/flaggy"
)

const version = "0.0.0"

func init() {
	flaggy.SetName("fractals")
	flaggy.SetDescription("generate fractals from the CLI, or run a server to generate fractals from an API")
	flaggy.SetVersion(version)
	flaggy.AttachSubcommand(cli.Subcommand, 1)
	flaggy.Parse()
}

func main() {
	switch {
	case cli.Subcommand.Used:
		cli.Exec()
	default:
		flaggy.ShowHelpAndExit("No subcommand specified")
	}
}
