package command

import "deathlog-tracker/domain/port"

type WatchFileCommand struct {
	FolderPath string
}

type WatchFileCommandHandler struct {
	FileWatcher port.FileWatcher
	HandleOnChange func()
}

func (h *WatchFileCommandHandler) Handle(command WatchFileCommand) error {
	h.FileWatcher.Watch(command.FolderPath, h.HandleOnChange)
	return nil
}