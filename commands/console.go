package commands

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/nanobox-io/nanobox/processors"
	"github.com/nanobox-io/nanobox/util/display"
)

var (

	// ConsoleCmd ...
	ConsoleCmd = &cobra.Command{
		Use:   "console",
		Short: "Opens an interactive console inside a production component.",
		Long:  ``,
		Run:   consoleFn,
	}

	// consoleCmdFlags ...
	consoleCmdFlags = struct {
		app string
	}{}
)

//
func init() {
	ConsoleCmd.Flags().StringVarP(&consoleCmdFlags.app, "app", "a", "", "app name or alias")
}

// consoleFn ...
func consoleFn(ccmd *cobra.Command, args []string) {

	// validate we have args required to set the meta we'll need; if we don't have
	// the required args this will os.Exit(1) with an error message
	if len(args) != 1 {
		fmt.Printf(`
Wrong number of arguments (expecting 1 got %v). Run the command again with the
name of the component you wish to console into:

ex: nanobox console [-a appname] <container>

`, len(args))
		return
	}

	// set the meta arguments to be used in the processor and run the processor
	display.CommandErr(processors.Console(consoleCmdFlags.app, args[0]))
}
