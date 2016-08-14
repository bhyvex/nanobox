package sim

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/nanobox-io/nanobox/models"
	"github.com/nanobox-io/nanobox/processor/sim"
	"github.com/nanobox-io/nanobox/util/print"
	"github.com/nanobox-io/nanobox/validate"
	"github.com/nanobox-io/nanobox/util/config"
)

// ConsoleCmd ...
var ConsoleCmd = &cobra.Command{
	Use:    "console",
	Short:  "Opens an interactive console inside your sim platform.",
	Long:   ``,
	PreRun: validate.Requires("provider", "provider_up", "built", "sim_deployed"),
	Run:    consoleFn,
}

// consoleFn ...
func consoleFn(ccmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("you need to provide a container to console into")

		return
	}

	component, _ := models.FindComponentBySlug(config.EnvID()+"_sim", args[0])

	simConsole := sim.Console{Component: component}

	print.OutputCommandErr(simConsole.Run())
}
