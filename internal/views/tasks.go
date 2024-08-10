package views

import (
	"github.com/rivo/tview"
	"github.com/gdamore/tcell/v2"
	"github.com/shappy0/ntui/internal/widgets"
	"github.com/shappy0/ntui/internal/models"
	"github.com/shappy0/ntui/internal/utils"
)

var (
	TitleTasks = "task"
	TitleEvents = "events"
	TaskLeftInfoKeys = []string{"Name", "Status", "JobId", "Client", "Modified At", "Started At"}
	TaskLeftInfoValues = []string{"", "", "", "", "", ""}
	TaskRightInfoKeys = []string{"Namespace", "Version", "Driver", "Image", "Volumes", "LifeCycle"}
	TaskRightInfoValues = []string{"", "", "", "", "", ""}
)

type Tasks struct {
	*widgets.Flex
	Title 			string
	DetailsView 	*widgets.Flex
	InfoView 		*widgets.Flex
	UsageView 		*widgets.Flex
	EventsTable		*widgets.Table
	Menus 			[]widgets.Item
	RemoveMenus		[]widgets.Item
	Data			models.Allocations
}

var TaskMenus = []widgets.Item{
	widgets.RestartTaskMenu,
	widgets.LogMenu,
}

var RemoveTaskMenus = []widgets.Item{
	widgets.EnterMenu,
	widgets.UpArrowMenu,
	widgets.DownArrowMenu,
}

func NewTasks() *Tasks {
	t := &Tasks{
		Flex:			widgets.NewFlex(),
		Title:			TitleTasks,
		DetailsView: 	widgets.NewFlex(),
		EventsTable: 	widgets.NewTable(TitleEvents),
		Menus:			TaskMenus,
		RemoveMenus:	RemoveTaskMenus,
		Data:			models.Allocations{},
	}
	t.SetTitleX(t.Title, "")
	return t
}

func (t *Tasks) GetTitle() string {
	return t.Title
}

func (t *Tasks) DrawView(data models.Allocations) {
	t.Data = data
	GetTaskData(data)
	t.SetDirection(tview.FlexRow)
	TgTitleName := data.Tasks.Name + "/" + utils.GetID(data.ID)
	t.SetTitleX(t.Title, TgTitleName)
	
	t.TaskDetails(data.Tasks)
	t.AddItem(t.DetailsView, 0, 1, true)

	t.EventsView(data.Events)
	t.AddItem(t.EventsTable, 0, 1, false)
}

func GetTaskData(data models.Allocations) {
	TaskLeftInfoValues[0] = utils.ToCapitalize(data.Name)
	TaskLeftInfoValues[1] = utils.ToCapitalize(data.Status)
	TaskLeftInfoValues[2] = utils.ToCapitalize(data.JobID)
	TaskLeftInfoValues[3] = utils.GetID(data.Client)
	TaskLeftInfoValues[4] = utils.DateTimeToStr(data.Modified)
	for _, event := range data.Events{
		if event.Type == "Started" {
			TaskLeftInfoValues[5] = utils.DateTimeToStr(event.Time)
			break
		}
	}
	TaskRightInfoValues[0] = utils.ToCapitalize(data.Namespace)
	TaskRightInfoValues[1] = utils.IntToStr(data.Version)
	TaskRightInfoValues[2] = utils.ToCapitalize(data.Tasks.Driver)
	TaskRightInfoValues[3] = data.Tasks.Config["image"].(string)
	Volumn, Ok := data.Tasks.Config["volumes"]
	if Ok && len(Volumn.([]interface{})) > 0 {
		TaskRightInfoValues[4] = Volumn.([]interface{})[0].(string)
	}
	TaskRightInfoValues[5] = utils.ToCapitalize("Main")
}

func (t *Tasks) TaskDetails(task models.Tasks) {
	t.DetailsView.SetBorder(false)
	t.SetInfoView()

	resourceUsage := task.Resources
	t.DrawUsageGauges()
	t.SetUsageData(resourceUsage)
}

func (t *Tasks) SetInfoView() {
	//Top Left Section
	t.InfoView = widgets.NewFlex()
	t.InfoView.SetBorder(false)
	//Top Left Left
	InfoLeftTable := widgets.NewMapView()
	InfoLeftTable.SetMapKeys(TaskLeftInfoKeys)
	InfoLeftTable.SetMapValues(TaskLeftInfoValues)
	InfoLeftTable.DrawMapView()

	//Top Left Right
	InfoRightTable := widgets.NewMapView()
	InfoRightTable.SetMapKeys(TaskRightInfoKeys)
	InfoRightTable.SetMapValues(TaskRightInfoValues)
	InfoRightTable.DrawMapView()

	t.InfoView.AddItemX(InfoLeftTable, 0, 1, false)
	t.InfoView.AddItemX(InfoRightTable, 0, 1, false)

	t.DetailsView.AddItemX(t.InfoView, 0, 1, false)
}

func (t *Tasks) DrawUsageGauges() {
	//Top Right Section
	t.UsageView = widgets.NewFlex()
	t.UsageView.SetBorder(false)
	t.UsageView.SetDirection(tview.FlexRow)
	
	//CPU Usage Guage
	cpu_gauge := widgets.NewUtilGauge("CPU:   ")
	t.UsageView.AddItemX(cpu_gauge,  0, 1, false)

	//Memory Usage Guage
	memory_gauge := widgets.NewUtilGauge("Memory:")
	t.UsageView.AddItemX(memory_gauge,  0, 1, false)

	t.DetailsView.AddItemX(t.UsageView, 0, 1, false)
}

func (t *Tasks) SetUsageData(usage models.TaskResource) {
	//CPU Usage Guage
	cpu_gauge := widgets.PrimitiveToGauge(t.UsageView.GetItem(0))

	cpu_gauge.SetValue(usage.CPUPercent)
	//Memory Usage Guage
	memory_gauge := widgets.PrimitiveToGauge(t.UsageView.GetItem(1))
	memory_gauge.SetValue(usage.MemoryPercent)
}

func (t *Tasks) EventsView(events []models.Events) {
	t.EventsTable.Headers = []string{"time", "type", "description"}
	t.EventsTable.SetBorder(false)
	t.EventsTable.ClearTable()
	t.EventsTable.DrawHeaderLeft()
	t.UpdateEventTable(events)
}

func (t *Tasks) UpdateEventTable(events []models.Events) {
	t.EventsTable.ClearTable()
	t.EventsTable.SetSelectable(false, false)
	RowTextColor := tcell.ColorWhite
	for I := 0; I < len(events); I++ {
		t.EventsTable.DrawLeftCell(I + 1, 0, utils.DateTimeDiff(events[I].Time), RowTextColor)
		t.EventsTable.DrawLeftCell(I + 1, 1, events[I].Type, RowTextColor)
		t.EventsTable.DrawLeftCell(I + 1, 2, events[I].DisplayMessage, RowTextColor)
	}
}