package views

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
	"github.com/shappy0/ntui/internal/widgets"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/utils"
)

var (
	TitleTasks = "tasks"
	TitleEvents = "events"
	TaskLeftInfoKeys = []string{"Status", "JobId", "Namespace", "Client", "Modified At", "Started At", "Driver"}
	TaskLeftInfoValues = []string{"", "", "", "", "", "", ""}
	TaskRightInfoKeys = []string{"Name", "Version", "Image", "Volumes", "LifeCycle"}
	TaskRightInfoValues = []string{"", "", "", "", ""}
)

type Tasks struct {
	*widgets.Flex
	Title 		string
	DetailsView *widgets.Flex
	InfoView 	*widgets.Flex
	UsageView 	*tview.TextView
	EventsTable	*widgets.Table
}

func NewTasks() *Tasks {
	t := &Tasks{
		Flex:		widgets.NewFlex(),
		Title:		TitleTasks,
		DetailsView: widgets.NewFlex(),
		EventsTable: widgets.NewTable(TitleEvents),
	}
	t.SetTitleX(t.Title, "")
	return t
}

func (t *Tasks) GetTitle() string {
	return t.Title
}

func (t *Tasks) DrawView(data models.Allocations) {
	GetTaskData(data)
	TgTitleName := data.Tasks.Name + "/" + utils.GetID(data.ID)
	t.SetTitleX(t.Title, TgTitleName)
	
	t.TaskDetails(data.Tasks)
	t.AddItem(t.DetailsView, 0, 1, true)

	t.EventsView(data.Events)
	t.AddItem(t.EventsTable, 0, 1, false)
}

func GetTaskData(data models.Allocations) {
	TaskLeftInfoValues[0] = utils.ToCapitalize(data.Status)
	TaskLeftInfoValues[1] = utils.ToCapitalize(data.JobID)
	TaskLeftInfoValues[2] = utils.ToCapitalize(data.Namespace)
	TaskLeftInfoValues[3] = utils.GetID(data.Client)
	TaskLeftInfoValues[4] = utils.DateTimeToStr(data.Modified)
	TaskLeftInfoValues[6] = utils.ToCapitalize(data.Tasks.Driver)
	for _, event := range data.Events{
		if event.Type == "Started" {
			TaskLeftInfoValues[5] = utils.DateTimeToStr(event.Time)
			break
		}
	}
	TaskRightInfoValues[0] = utils.ToCapitalize(data.Name)
	TaskRightInfoValues[1] = utils.IntToStr(data.Version)
	TaskRightInfoValues[2] = data.Tasks.Config["image"].(string)
	Volumn, Ok := data.Tasks.Config["volumes"]
	if Ok && len(Volumn.([]interface{})) > 0 {
		TaskRightInfoValues[3] = Volumn.([]interface{})[0].(string)
	}
	TaskRightInfoValues[4] = utils.ToCapitalize("Main")
}

func (t *Tasks) TaskDetails(task models.Tasks) {
	t.DetailsView.SetBorder(false)
	t.DetailsView.SetDirection(tview.FlexRow)
	t.SetInfoView()
	t.SetUsageView()
}

func (t *Tasks) SetInfoView() {
	//Top Section
	t.InfoView = widgets.NewFlex()
	t.InfoView.SetBorder(false)
	//Top Left
	InfoLeftTable := widgets.NewMapView()
	InfoLeftTable.SetMapKeys(TaskLeftInfoKeys)
	InfoLeftTable.SetMapValues(TaskLeftInfoValues)
	InfoLeftTable.DrawMapView()

	//Top Right
	InfoRightTable := widgets.NewMapView()
	InfoRightTable.SetMapKeys(TaskRightInfoKeys)
	InfoRightTable.SetMapValues(TaskRightInfoValues)
	InfoRightTable.DrawMapView()

	t.InfoView.AddItemX(InfoLeftTable, 0, 1, false)
	t.InfoView.AddItemX(InfoRightTable, 0, 1, false)

	t.DetailsView.AddItemX(t.InfoView, 0, 1, false)
}

func (t *Tasks) SetUsageView() {
	//Bottom Section
	t.UsageView = tview.NewTextView()
	t.UsageView.SetText("This place for usage")
	t.DetailsView.AddItemX(t.UsageView, 0, 1, false)
}

func (t *Tasks) EventsView(events []models.Events) {
	t.EventsTable.Headers = []string{"time", "type", "description"}
	t.EventsTable.SetBorder(false)
	t.EventsTable.ClearTable()
	t.EventsTable.DrawHeader()
	t.UpdateEventTable(events)
}

func (t *Tasks) UpdateEventTable(events []models.Events) {
	t.EventsTable.ClearTable()
	t.EventsTable.SetSelectable(false, false)
	RowTextColor := tcell.ColorWhite
	for I := 0; I < len(events); I++ {
		t.EventsTable.DrawCell(I + 1, 0, utils.DateTimeToStr(events[I].Time), RowTextColor)
		t.EventsTable.DrawCell(I + 1, 1, events[I].Type, RowTextColor)
		t.EventsTable.DrawCell(I + 1, 2, events[I].DisplayMessage, RowTextColor)
	}
}