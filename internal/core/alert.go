package core

type Alert struct {
	App 	*App
}

func NewAlert(app *App) *Alert {
	a := &Alert{
		App:	app,
	}
	return a
}

func (a *Alert) Listen() {
	for {
		select {
		case Msg := <-a.App.Alert.Channel():
			a.App.Layout.QueueUpdateDraw(func() {
				a.App.Layout.Footer.SetAlert(Msg)
			}) 
		}
	}
}