package command

import (
	"deathlog-tracker/domain/port"
	"deathlog-tracker/infra/secondary"
	"log"
	"path/filepath"
)

type DeathLogFileFullProcessingCommand struct {
	WowPath string
}

type DeathLogFileFullProcessingCommandHandler struct {
	WatchFileCommandHandlers []WatchFileCommandHandler
	PlayerHashRepository port.PlayerHashRepository
}

func (h *DeathLogFileFullProcessingCommandHandler) Handle(command DeathLogFileFullProcessingCommand) error {
	findDeathLogFileCommand := FindDeathLogFileCommand{
		FolderPath: command.WowPath,
	}

	findDeathLogFileCommandHandler := FindDeathLogFileCommandHandler{
		FileFinder: &secondary.FileFinderImpl{},
	}

	filePaths, err := findDeathLogFileCommandHandler.Handle(findDeathLogFileCommand)
	if err != nil {
		panic(err.Error())
	}
	log.Println("[DEATHLOG-TRACKER] Found Deathlog.lua files : ", filePaths)



	for _, file := range filePaths {
		watchFileCommand := WatchFileCommand{
			FolderPath: filepath.Dir(file),
		}

		buildBatchFromFileCommand := BuildBatchFromFileCommand{
			FilePath: file,
		}

		buildBatchFromFileCommandHandler := BuildBatchFromFileCommandHandler{
			FileReader: &secondary.FileReaderImpl{},
		}

		watchFileCommandHandler := WatchFileCommandHandler{
			FileWatcher: &secondary.FileWatcherImpl{},
			HandleOnChange: func() {
				playerBatch, err := buildBatchFromFileCommandHandler.Handle(buildBatchFromFileCommand)
				if err != nil {
					panic(err)
				}
				processAndSendBatchOfDeathCommand := ProcessAndSendBatchOfDeathCommand{
					PlayerBatch: playerBatch,
				}

				processAndSendBatchOfDeathCommandHandler := ProcessAndSendBatchOfDeathCommandHandler{
					PlayerSender: &secondary.PlayerSenderImpl{},
					PlayerHashRepository: h.PlayerHashRepository,
				}

				err = processAndSendBatchOfDeathCommandHandler.Handle(processAndSendBatchOfDeathCommand)
				if err != nil {
					panic(err)
				}
			},
		}
		h.WatchFileCommandHandlers = append(h.WatchFileCommandHandlers, watchFileCommandHandler)
		err = watchFileCommandHandler.Handle(watchFileCommand)
		if err != nil {
			panic(err.Error())
		}
	}

	return nil
}

func (h *DeathLogFileFullProcessingCommandHandler) Cancel() {
	for _, watchFileCommandHandler := range h.WatchFileCommandHandlers {
		watchFileCommandHandler.FileWatcher.Cancel()
	}
}