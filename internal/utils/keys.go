package utils

import (
	"github.com/gdamore/tcell/v2"
)

type NtuiKey struct {
	Key 			tcell.Key
	KeyRune 		rune
	KeyLabel 		string
	KeyDescription  string
}

var (
	NtuiExitKey = NtuiKey{
		Key:			tcell.KeyCtrlC,
		KeyLabel:		"Ctrl + c",
		KeyDescription: "Exit Ntui Application",
	}
	NtuiEscKey = NtuiKey{
		Key:			tcell.KeyEcs,
		KeyLabel:		"Backspace",
		KeyDescription: "To go back to previous screen",
	}
	NtuiTabKey = NtuiKey{
		Key:			tcell.KeyTAB,
		KeyLabel:		"Tab Key",
		KeyDescription: "To switch focus",
	}
	NtuiCtrRKey = NtuiKey{
		Key:			tcell.KeyCtrlR,
		KeyLabel:		"Ctrl + r",
		KeyDescription: "To switch to Region Page",
	}
	NtuiCtrlVKey = NtuiKey{
		Key:			tcell.KeyCtrlV,
		KeyLabel:		"Ctrl + v",
		KeyDescription: "To Switch to Job Versions",
	}
	NtuiRuneKey = NtuiKey{
		Key:			tcell.KeyRune,
		KeyLabel:		"L",
		KeyDescription: "To  Open Log",
	}
)

var NtuiKeyBindings = []NtuiKey{
	NtuiExitKey,
}