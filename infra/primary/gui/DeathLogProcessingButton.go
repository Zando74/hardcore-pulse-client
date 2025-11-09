package gui

import (
	"deathlog-tracker/app/command"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

type DeathLogProcessingButton struct {
	App                                      fyne.App
	Window                                   fyne.Window
	DeathLogFileFullProcessingCommandHandler command.DeathLogFileFullProcessingCommandHandler
	GamePathBinding                          binding.String
}

func (d *DeathLogProcessingButton) Start() {
	path, err := d.GamePathBinding.Get()
	if err != nil {
		return
	}

	go func() {
		d.DeathLogFileFullProcessingCommandHandler.Handle(command.DeathLogFileFullProcessingCommand{
			WowPath: path,
		})
	}()
	
}

func (d *DeathLogProcessingButton) Stop() {
	log.Println("[DEATHLOG-TRACKER] Stop called")
	d.DeathLogFileFullProcessingCommandHandler.Cancel()
	d.DeathLogFileFullProcessingCommandHandler = command.DeathLogFileFullProcessingCommandHandler{
		WatchFileCommandHandlers: []command.WatchFileCommandHandler{},
		PlayerHashRepository: d.DeathLogFileFullProcessingCommandHandler.PlayerHashRepository,
	}
}