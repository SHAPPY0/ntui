package core

import (
	"github.com/shappy0/ntui/internal/views"
	"github.com/shappy0/ntui/internal/models"
)

type JobDefinition struct {
	*views.JobDefinition
	App		*App
	Data	map[string]string
}

func NewJobDefinition(app *App) *JobDefinition {
	jd := &JobDefinition{
		JobDefinition:	views.NewJobDefinition(),
		App:			app,
		Data:			make(map[string]string),
	}
	jd.App.Layout.Body.AddPageX(jd.GetTitle(), jd, true, false)
	jd.SetFocusFunc(jd.OnFocus)
	jd.SetBlurFunc(jd.OnBlur)
	return jd
}

func (jd *JobDefinition) OnFocus() {
	jd.App.Layout.Header.Menu.RemoveMenus(jd.RemoveMenus)
	go jd.UpdateTable()
}

func (jd *JobDefinition) OnBlur() {
	// jd.App.Layout.Header.Menu.RemoveMenus(j.Menus)
}

func (jd *JobDefinition) UpdateTable() {
	selectedJob := jd.App.Primitives.Jobs.SelectedValue
	if selectedJob != nil {
		params := &models.NomadParams{
			Region:		jd.App.Config.GetRegion(),
			Namespace:	jd.App.Config.GetNamespace(),
		}
		jd.SetTextVTitle(selectedJob["name"], selectedJob["version"])
		jd.App.Alert.Loader(true)
		data, err := jd.App.NomadClient.Submission(selectedJob["name"], selectedJob["version"], params)
		jd.App.Alert.Loader(false)
		if err != nil {
			jd.App.Logger.Errorf("Error getting job definition %s", err.Error())
			jd.App.Alert.Error("Error getting JobDefinition for " + selectedJob["name"])
		} else {
			jd.Data["source"] = data.Source
			jd.Data["format"] = data.Format
			// jd.Data["variableFlags"] = data.VariableFlags
			jd.Data["variables"] = data.Variables
			jd.SetJDData(jd.Data)
		}
	}
}