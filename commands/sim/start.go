package sim

import (
	"github.com/nanobox-io/golang-docker-client"
	"github.com/spf13/cobra"

	"github.com/nanobox-io/nanobox/commands/steps"
	"github.com/nanobox-io/nanobox/models"
	"github.com/nanobox-io/nanobox/processors/app"
	"github.com/nanobox-io/nanobox/processors/env"
	"github.com/nanobox-io/nanobox/processors/provider"
	"github.com/nanobox-io/nanobox/util/config"
	"github.com/nanobox-io/nanobox/util/display"
)

func init() {
	steps.Build("sim start", startCheck, simStart)
}

// simStart ...
func simStart(ccmd *cobra.Command, args []string) {
	envModel, _ := models.FindEnvByID(config.EnvID())
	appModel, _ := models.FindAppBySlug(config.EnvID(), "sim")

	display.CommandErr(env.Setup(envModel))
	display.CommandErr(app.Start(envModel, appModel, "sim"))
}

func startCheck() bool {
	app, _ := models.FindAppBySlug(config.EnvID(), "sim")
	if app.Status != "up" {
		return false
	}
	provider.Init()
	components, _ := app.Components()
	for _, component := range components {
		info, err := docker.ContainerInspect(component.ID)
		if err != nil || !info.State.Running {
			return false
		}
	}
	return true
}
