package utils

import {
	"fmt"
	"github.com/gdamore/tcell/v2"
}

const (
	ColorWhite	=	tcell.ColorWhite
	ColorGray 	=	tcell.ColorGray
	ColorRed    =   tcell.ColorRed
	ColorGreen  =	tcell.ColorGreen
)

func ColorizeStatusCell(status string) (tcell.Color, string) {
	CellColor := tcell.ColorWhite
	Status = ToCapitalize(status)
	if Status == "Dead" {
		CellColor = tcell.ColorGray
	} else if Status == "Failed" {
		Status = SetCellTextColor(Status, "red")
	} else if Status == "Running" {
		Status = SetCellTextColor(Status, "green")
	} else if Status == "Lost" {
		Status = SetCellTextColor(Status, "gray")
	} else if Status == "Complete" {
		Status = SetCellTextColor(Status, "olive")
	} else if Status == "Pending" {
		Status = SetCellTextColor(Status, "yellow")
	}
	return CellColor, Status
}

func SetCellTextColor(text, color string) string {
	Text = fmt.Sprintf("[%s]%s", color, text)
	return Text
}