package sim

import (
	"github.com/nanobox-io/nanobox/models"
	"github.com/nanobox-io/nanobox/processor/platform"
)

// Log ...
type Log struct {
	App models.App
}

//
func (log Log) Run() error {
	// some messaging about the logging??
	platformMistListen := platform.MistListen{App: log.App}
	return platformMistListen.Run()
}
