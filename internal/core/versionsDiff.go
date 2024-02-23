package core

import (
	"github.com/shappy0/ntui/internal/views"
	"github.com/shappy0/ntui/internal/utils"
	"github.com/shappy0/ntui/internal/models"
)

type VersionDiff struct {
	*views.VersionDiff
	App			*App
	Listener	*utils.Listener
	JobId		string
	Data		map[string]string
}

func NewVersionDiff(app *App) *VersionDiff {
	vd := &VersionDiff{
		VersionDiff:		views.NewVersionDiff(),
		App:				app,
	}
	vd.App.Layout.Body.AddPageX(vd.GetTitle(), vd, true, false)
	return vd
}

func (vd *VersionDiff) Render(jobId string, row map[string]string, diff models.JobVersionDiff) {
	vd.SetData(jobId, row, diff)
}