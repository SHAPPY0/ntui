package views

import (
	"fmt"
	"github.com/shappy0/ntui/internal/widgets"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/utils"
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
		Title:		TitleVersionDiff,
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
			Data += fmt.Sprintf("[b][%s]TaskGroup:[%s] %s\n", utils.ColorT70d5bf, utils.ColorTWhite, diff.TaskGroups[I].Name)
			Tasks := diff.TaskGroups[I].Tasks
			for J := 0; J < len(Tasks); J++ {
				if Tasks[J].Type != "None" {
					if Tasks[J].Type == "Edited" {
						Data += fmt.Sprintf("[[%s]+[%s]/[%s]-[%s]] %s\n", utils.ColorTGreen, utils.ColorTWhite, utils.ColorTRed, utils.ColorTWhite, Tasks[J].Name)
						TaskObjects := Tasks[J].Objects
						for K := 0; K < len(TaskObjects); K++ {
							Data += fmt.Sprintf("  [%s]%s:[%s]\n", utils.ColorT70d5bf, TaskObjects[K].Name, utils.ColorTWhite)
							TaskFields := TaskObjects[K].Fields
							for M := 0; M < len(TaskFields); M++ {
								if TaskFields[M].Type == "Edited" {
									Data += fmt.Sprintf("   [[%s]+[%s/[%s]-[%s]] %s: %s [%s]=>[%s] %s\n", 
										utils.ColorTGreen,
										utils.ColorTWhite,
										utils.ColorTRed,
										utils.ColorTWhite,
										TaskFields[M].Name,
										TaskFields[M].Old,
										utils.ColorTGreen,
										utils.ColorTWhite,
										TaskFields[M].New,
									)
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