package core

import (
	// "fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/shappy0/ntui/internal/nomad"
	"github.com/shappy0/ntui/internal/utils"
)

type App struct {
	Config 			*Config
	Version 		string
	Layout 			*Layout
	NomadClient 	*nomad.Nomad
	Primitives		PrimitivesX
	Alert 			*utils.Alert
	Logger 			*utils.Logger
}

type PrimitivesX struct {
	Regions			*Regions
	Namespaces		*Namespaces
	Main			*Main
	Jobs			*Jobs
	TaskGroups		*TaskGroups
	Allocations		*Allocations
	Tasks			*Tasks
	Log				*Log
	Versions		*Versions
	VersionDiff 	*VersionDiff
	Nodes 			*Nodes
	Modal 			*Modal
	JobDefinition	*JobDefinition
}

func NewApp(version string, config *Config, logger *utils.Logger) (*App, error) {
	a := &App{
		Version:		version,
		Config:			config,
		Layout:			NewLayout(version),
		Alert:			utils.NewAlert(),
		Logger:			logger,
	}
	NomadClient, Err :=	nomad.New(a.Logger)
	if Err != nil {
		return a, Err
	}
	a.NomadClient = NomadClient
	return a, nil
}

func (app *App) Init() error {
	app.Logger.Info("Initializing ntui app ...")
	app.Primitives = PrimitivesX{
		Regions:		NewRegions(app),
		Namespaces:		NewNamespaces(app),
		Jobs:			NewJobs(app),
		TaskGroups:		NewTaskGroups(app),
		Allocations:	NewAllocations(app),
		Tasks:			NewTasks(app),
		Log:			NewLog(app),
		Versions:		NewVersions(app),
		VersionDiff:	NewVersionDiff(app),
		Nodes:			NewNodes(app),
		Modal:			NewModal(app),
		JobDefinition:	NewJobDefinition(app),
	}
	app.Primitives.Main  = NewMain(app)
	BindAppKeys(app)
	Alert := NewAlert(app)
	go Alert.Listen()
	return nil
}

func BindAppKeys(app *App) {
	app.Layout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case utils.NtuiExitKey.Key:
			app.Logger.Info("Stopping ntui by user ...")
			app.StopX()
			break
		case utils.NtuiEscKey.Key:
			app.Layout.GoBack()
			break
		case utils.NtuiCtrlVKey.Key:
			if app.Layout.GetActivePage() == app.Primitives.Versions.GetTitle() {
				app.Primitives.Versions.ConfirmModal()
			}
			break
		case utils.NtuiCtrlTKey.Key:
			if app.Layout.GetActivePage() == app.Primitives.Allocations.GetTitle() || app.Layout.GetActivePage() == app.Primitives.Tasks.GetTitle() {
				app.Primitives.Allocations.InitRestartModal()
			}
			break
		case utils.NtuiCtrlQKey.Key:
			if app.Layout.GetActivePage() == app.Primitives.Jobs.GetTitle() {
				app.Primitives.Jobs.StopModal()
			}
			break
		case utils.NtuiCtrlSKey.Key:
			if app.Layout.GetActivePage() == app.Primitives.Jobs.GetTitle() {
				app.Primitives.Jobs.StartModal()
			}
			break
		case utils.NtuiRuneKey.Key:
			switch event.Rune() {
			case '1':
				app.Layout.OpenPage(app.Primitives.Nodes.GetTitle(), true)
				break
			case '2':
				app.Layout.OpenPage(app.Primitives.Main.GetTitle(), false)
				break
			case 'v':
				if app.Layout.GetActivePage() == app.Primitives.TaskGroups.GetTitle() {
					app.Layout.OpenPage(app.Primitives.Versions.GetTitle(), true)
				}
				break
			case 'd':
				if app.Layout.GetActivePage() == app.Primitives.Jobs.GetTitle() {
					app.Primitives.Jobs.GoToDefinitions()
				}
				break
			case 'l':
				if app.Layout.GetActivePage() == app.Primitives.Tasks.GetTitle() || app.Layout.GetActivePage() == app.Primitives.Allocations.GetTitle() {
					app.Primitives.Allocations.SetSelectedRow()
					app.Primitives.Log.SetPageSource(app.Layout.GetActivePage())
					app.Layout.OpenPage(app.Primitives.Log.GetTitle(), true)
				}
				break
			case 'e':
				if app.Layout.GetActivePage() == app.Primitives.Log.GetTitle() {
					app.Primitives.Log.FetchStdErrLog()
				}
				break
			case 'o':
				if app.Layout.GetActivePage() == app.Primitives.Log.GetTitle() {
					app.Primitives.Log.FetchStdOutLog()
				}
				break
			case 'a':
				if app.Layout.GetActivePage() == app.Primitives.Log.GetTitle() {
					app.Primitives.Log.SetFollow(true)
				}
				break
			case 'f':
				if app.Layout.GetActivePage() == app.Primitives.Log.GetTitle() {
					app.Primitives.Log.ShowFullScreen(true)
				}
				break
			}
		}
		return event
	})
}

func (app *App) RunX() error {
	if Err := app.Layout.Run(app); Err != nil {
		return Err
	}
	return nil
}

func (app *App) StopX() {
	app.Layout.Stop()
}