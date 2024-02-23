package views

import (
	"github.com/shappy0/ntui/internal/widgets"
	"github.com/shappy0/ntui/internal/models"
)

var TitleVersionDiff = "versiondiff"

type VersionDiff struct {
	*widgets.TextView
	Title 		string
	Data		[]string
}

func NewVersionDiff() *VersionDiff {
	v := &VersionDiff{
		TextView:	widgets.NewTextView(TitleVersionDiff),
		Title:		TitleTaskGroups,
	}
	v.SetDynamicColors(true)
	return v
}

func (vd *VersionDiff) GetTitle() string {
	return vd.Title
}

func (vd *VersionDiff) SetData(jobId string, detail map[string]string, diff models.JobVersionDiff) {
	vd.SetTextVTitle(jobId, detail["version"])
	Data := ""
	for I := 0; I < len(diff.TaskGroups); I++ {
		if diff.TaskGroups[I].Type != "None" {
			Data += "[b][blue]TaskGroup:[white] " + diff.TaskGroups[I].Name + "\n"
			Tasks := diff.TaskGroups[I].Tasks
			for J := 0; J < len(Tasks); J++ {
				if Tasks[J].Type != "None" {
					if Tasks[J].Type == "Edited" {
						Data += "[[green]+[white]/[red]-[white]]" + Tasks[J].Name + "\n"
						TaskObjects := Tasks[J].Objects
						for K := 0; K < len(TaskObjects); K++ {
							Data += "-- [blue]" + TaskObjects[K].Name + ":[white]\n"
							TaskFields := TaskObjects[K].Fields
							for M := 0; M < len(TaskFields); M++ {
								if TaskFields[M].Type == "Edited" {
									Data += "---[[green]+[white]/[red]-[white]] " + TaskFields[M].Name + ": " + TaskFields[M].Old + " [green]=>[white] " + TaskFields[M].New + "\n"
								}
							}
						} 

					}
				}
			}
		}
	}
	vd.SetText(Data)
}