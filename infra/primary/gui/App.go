package gui

import (
	"deathlog-tracker/app/command"
	"deathlog-tracker/app/query"
	"deathlog-tracker/infra/secondary"
	"log"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"fyne.io/systray"
)

type App struct {
	App        fyne.App
	Window     fyne.Window
	GamePathUI GamePathUI
	DeathLogBtn DeathLogProcessingButton
	switchBtn *widget.Button
	Started     binding.Bool
	GamePathBinding binding.String
}

func (a *App) Run() {

	a.ShowMainScreen()

	a.Window.Resize(fyne.NewSize(800, 600))

	go func() {
        systray.Run(onReady, onExit)
    }()


    go func() {
        prevSize := a.Window.Canvas().Size()
        for {
            currentSize := a.Window.Canvas().Size()
            if currentSize.Width < prevSize.Width && currentSize.Height < prevSize.Height {
                fyne.Do(func() {
                    a.Window.Hide()
                })
            }
            prevSize = currentSize
            time.Sleep(100 * time.Millisecond)
        }
    }()

	a.Window.ShowAndRun()
}

func (a *App) ShowMainScreen() {
	gamePathLabel := widget.NewLabelWithData(a.GamePathUI.GamePathBinding)

	btnText := "Start Watching Deathlog cache updates"

	
	a.switchBtn = widget.NewButton(btnText, func() {
		startedVal, _ := a.Started.Get()
		path, _ := a.GamePathBinding.Get()
		if path == "" {
			log.Println("[DEATHLOG-TRACKER] You must set game path before starting")
			return
		}

		if !startedVal {
			a.Started.Set(true)
			a.DeathLogBtn.Start()
			if(a.switchBtn != nil) {
				a.switchBtn.SetText("Stop Watching Deathlog cache updates")
			}
		} else {
			a.Started.Set(false)
			a.DeathLogBtn.Stop()
			if(a.switchBtn != nil) {
				a.switchBtn.SetText("Start Watching Deathlog cache updates")
			}
		}
	})

	content := container.NewBorder(
		container.NewVBox(
			a.GamePathUI.getMainScreenContent(func() {
				startedVal, _ := a.Started.Get()
				if startedVal {
					a.Started.Set(false)
					a.DeathLogBtn.Stop()
					if(a.switchBtn != nil) {
						a.switchBtn.SetText("Start Watching Deathlog cache updates")
					}
				}
			}),
			container.NewHBox(widget.NewLabel("Game Folder: "), gamePathLabel),
			a.switchBtn,
		),      
		NewLogView(),
		nil,      
		nil,      
		nil,
	)

	log.Println("[DEATHLOG-TRACKER] Application started")



	a.Window.SetContent(content)
}

func NewApp() *App {
	a := app.NewWithID("deathlog-tracker")
	w := a.NewWindow("Deathlog Tracker")

	repo, err := secondary.NewJSONGamePathRepository("deathlog-tracker")
	if err != nil {
		panic(err)
	}

	hashRepo, err := secondary.NewPlayerHashRepositoryImpl("deathlog-tracker")
		if err != nil {
			panic(err)
		}

	gamePath := binding.NewString()

	return &App{
		App:    a,
		Window: w,
		GamePathBinding: gamePath,
		GamePathUI: GamePathUI{
			App: a,
			Window: w,
			AskGamePathCommandHandler: command.AskGamePathCommandHandler{
			GamePathRepository: repo,
			},
			GetGamePathCommandHandler: query.GetGamePathQueryHandler{
				GamePathRepository: repo,
			},
			GamePathBinding: gamePath,
		},
		DeathLogBtn: DeathLogProcessingButton{
			App: a,
			Window: w,
			DeathLogFileFullProcessingCommandHandler: command.DeathLogFileFullProcessingCommandHandler{
				WatchFileCommandHandlers: []command.WatchFileCommandHandler{},
				PlayerHashRepository: hashRepo,
			},
			GamePathBinding: gamePath,
		},
		Started: binding.NewBool(),
	}
}

func onReady() {
    systray.SetTitle("Deathlog Tracker")
    systray.SetTooltip("Deathlog Tracker")

    mShow := systray.AddMenuItem("Open", "Open the window")
    mQuit := systray.AddMenuItem("Quit", "Quit the application")

    go func() {
        for {
            select {
            case <-mShow.ClickedCh:
				fyne.Do(func() {
					fyne.CurrentApp().Driver().AllWindows()[0].Show()
				})
            case <-mQuit.ClickedCh:
                systray.Quit()
                return
            }
        }
    }()
}

func onExit() {
	os.Exit(0)
}
