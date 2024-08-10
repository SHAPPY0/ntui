package core

import (
	"strings"
	"fmt"
	// "github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
	"github.com/shappy0/ntui/internal/views"
	"github.com/shappy0/ntui/internal/utils"
	"github.com/hashicorp/nomad/api"
	"github.com/shappy0/ntui/internal/models"
)

type RunJob struct {
	*views.RunJob
	App		*App
	Data	map[string]string
	DryRan	bool
}

func NewRunJob(app *App) *RunJob {
	rj := &RunJob{
		RunJob:	views.NewRunJob(),
		App:	app,
		Data:	make(map[string]string),
		DryRan:	false,
	}
	rj.App.Layout.Body.AddPageX(rj.GetTitle(), rj, true, false)
	rj.SetFocusFunc(rj.OnFocus)
	rj.SetBlurFunc(rj.OnBlur)
	rj.BindKeys()
	return rj
}

func (rj *RunJob) BindKeys() {
	rj.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case utils.NtuiCtrlDKey.Key:
			rj.DryRun()
			break
		}
		return event
	})
}

func (rj *RunJob) DryRun() {
	jobHcl := rj.GetTextAreaText()
	if jobHcl != "" {
		job, err := rj.App.NomadClient.DryRun(jobHcl)
		if err != nil {
			err_msg := "Error parsing job HCL: " + err.Error()
			rj.App.Logger.Error(err_msg)
			rj.App.Alert.Error(err_msg)
			rj.App.OpenModal(err_msg, func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == utils.NtuiEscKey.Key {
					return nil
				}
				return event
			})
		} else {
			rj.DryRan = true
			var runOutput strings.Builder
			if job != nil {
				runOutput.WriteString("[::b][green] ✓ Dry Run Success[white][::B]\n\n")
				runOutput.WriteString("[green] +[orange]Job:[white] " + *job.Name + "\n")
				taskGroups := job.TaskGroups
				for i := 0; i < len(taskGroups); i++ {
					runOutput.WriteString(" [green] +[orange]Task Group:[white] " + *(taskGroups[i].Name) + "\n") 
					tasks := taskGroups[i].Tasks
					for j := 0; j < len(tasks); j++ {
						runOutput.WriteString("  [green] +[orange]Task:[white] " + tasks[j].Name + "\n")
					}
				}
			}
			if runOutput.Len() > 0 {
				runOutput.WriteString("\n\nPress [orange]ENTER[white] to submit the job")
			}
			rj.App.OpenModal(runOutput.String(), func(event *tcell.EventKey) *tcell.EventKey {
				if event.Key() == utils.NtuiEscKey.Key {
					return nil
				} else if event.Key() == utils.NtuiEnterKey.Key {
					rj.SubmitJob(job)
					return nil
				}
				return event
			})
		}
	} else {
		rj.App.Logger.Error("Job HCL is empty")
		rj.App.Alert.Error("Job HCL is empty")
	}	
}

func (rj *RunJob) SubmitJob(job *api.Job) {
	jobHcl := rj.GetTextAreaText()
	if rj.DryRan && jobHcl != "" {
		params := &models.NomadParams{
			Region:		rj.App.Config.GetRegion(),
			Namespace:	rj.App.Config.GetNamespace(),
		}
		if err := rj.App.NomadClient.Register(job, params); err != nil {
			rj.App.Alert.Loader(false)
			rj.App.Alert.Error("Job submission failed...")
			rj.App.Logger.Errorf("Job submission failed:: %s", err.Error())
		} else {
			rj.App.Alert.Loader(false)
			msg := fmt.Sprintf("Job %s submitted successfully...", *job.Name)
			rj.App.Alert.Info(msg)
			rj.App.Logger.Info(msg)
			rj.SetContent("")
			rj.App.Layout.OpenPage(rj.App.Primitives.Jobs.GetTitle(), true)
		}
	}
}

func (rj *RunJob) OnFocus() {
}

func (rj *RunJob) OnBlur() {
	rj.Clear()
}