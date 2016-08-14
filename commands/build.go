package commands

import (
	"github.com/spf13/cobra"

	"github.com/nanobox-io/nanobox/commands/registry"
	"github.com/nanobox-io/nanobox/models"
	"github.com/nanobox-io/nanobox/processor"
	"github.com/nanobox-io/nanobox/util/print"
	"github.com/nanobox-io/nanobox/util/config"
	"github.com/nanobox-io/nanobox/validate"
)

var (

	// BuildCmd ...
	BuildCmd = &cobra.Command{
		Use:   "build",
		Short: "Generates a deployable build package.",
		Long: `
Generates a deployable build package that can be
deployed into dev, sim, or production platforms.
		`,
		PreRun: validate.Requires("provider"),
		Run:    buildFn,
	}

	buildNoCompile bool
)

func init() {
	BuildCmd.PersistentFlags().BoolVarP(&buildNoCompile, "no-compile", "", false, "dont compile the build")
}

// buildFn ...
func buildFn(ccmd *cobra.Command, args []string) {
	if buildNoCompile {
		registry.Set("no-compile", true)
	}
	env, _ := models.FindEnvByID(config.EnvID())
	build := processor.Build{Env: env}
	print.OutputCommandErr(build.Run())
}
