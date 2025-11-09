package gui

import (
	"deathlog-tracker/app/command"
	"deathlog-tracker/app/query"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/sqweek/dialog"
)

type GamePathUI struct {
	App      fyne.App
	Window   fyne.Window
	AskGamePathCommandHandler command.AskGamePathCommandHandler
	GetGamePathCommandHandler query.GetGamePathQueryHandler
	GamePathBinding binding.String
}

func (g *GamePathUI) chooseFolder() {
	folderPath, err := dialog.Directory().Title("Select the game directory").Browse()
    if err != nil {
        return
    }
    g.AskGamePathCommandHandler.Handle(command.AskGamePathCommand{
        GetUserPath: func() string {
            return folderPath
        },
    })
	g.GamePathBinding.Set(folderPath)
}

func (g *GamePathUI) getMainScreenContent(cancelWatchHandler func()) fyne.CanvasObject {
	gamePath, err := g.GetGamePathCommandHandler.Handle(query.GetGamePathQuery{})
	if err != nil {
		panic(err)
	}
	indication := widget.NewLabel("Please select the game folder to ensure accurate scanning.\nThe next line should end with: \\_classic_era_ : ")

	g.GamePathBinding.Set(gamePath.Path)
	locationLabel := widget.NewLabel("Selected Game Folder : ")
	label := widget.NewLabelWithData(g.GamePathBinding)
	row := container.NewHBox(locationLabel, label)
	refreshBtn := widget.NewButton("Change game location", func() {
		cancelWatchHandler()
		g.chooseFolder()
	})

	return container.NewVBox(indication, row, refreshBtn)
}