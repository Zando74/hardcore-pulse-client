package command

import "deathlog-tracker/domain/port"

type FindDeathLogFileCommand struct {
	FolderPath string
}

type FindDeathLogFileCommandHandler struct {
	FileFinder port.FileFinder
}

func (h *FindDeathLogFileCommandHandler) Handle(command FindDeathLogFileCommand) ([]string, error) {

	files, err := h.FileFinder.Find(command.FolderPath)
	if err != nil {
		return nil, err
	}
	
	return files, nil
}