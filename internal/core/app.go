package core

import (
	"fmt"
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
	Modal 			*Modal
}

func NewApp(config *Config) (*App, error) {
	a := &App{
		Version:		"1.0",
		Config:			config,
		Layout:			NewLayout(),
		Alert:			utils.NewAlert(),
	}
	NomadClient, Err :=	nomad.New()
	if Err != nil {
		fmt.Println("Error")
		return a, Err
	}
	a.NomadClient = NomadClient
	return a, nil
}

func (app *App) Init() error {
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
		Modal:			NewModal(app),
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
			app.StopX()
			break
		case utils.NtuiEscKey.Key:
			app.Layout.GoBack()
			break
		case utils.NtuiCtrlRKey.Key:
			if app.Layout.GetActivePage() == "jobs" {
				app.Layout.OpenPage("main", false)
			}
			break
		case utils.NtuiCtrlVKey.Key:
			if app.Layout.GetActivePage() == "taskgroups" {
				app.Layout.OpenPage("versions", true)
			}
			break
		case utils.NtuiCtrlTKey.Key:
			if app.Layout.GetActivePage() == "allocations" {
				app.Primitives.Allocations.InitRestartModal()
			}
			break
		case utils.NtuiRuneKey.Key:
			switch event.Rune() {
			case 'l':
				if app.Layout.GetActivePage() == "tasks" {
					app.Layout.OpenPage("log", true)
				}
				break
			case 'e':
				if app.Layout.GetActivePage() == "log" {
					app.Primitives.Log.FetchStdErrLog()
				}
				break
			case 'o':
				if app.Layout.GetActivePage() == "log" {
					app.Primitives.Log.FetchStdOutLog()
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