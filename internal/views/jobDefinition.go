package views

import (
	"github.com/shappy0/ntui/internal/widgets"
	"github.com/shappy0/ntui/internal/models"
)

var titleDefinition = "job definition"

type JobDefinition struct {
	*widgets.TextView
	Title		string
	Data		[]models.Namespaces
	RemoveMenus []widgets.Item
}

var removeMenus = []widgets.Item{
	widgets.EnterMenu,
} 

func NewJobDefinition() *JobDefinition {
	jd := &JobDefinition {
		Title:		titleDefinition,
		TextView:	widgets.NewTextView(titleDefinition),
		RemoveMenus:removeMenus,
	}
	return jd
}

func (jd *JobDefinition) GetTitle() string {
	return jd.Title
}

func (jd *JobDefinition) SetJDData(data map[string]string) {
	jd.TextView.ClearX()
	jd.TextView.SetTextX(data["source"])
}