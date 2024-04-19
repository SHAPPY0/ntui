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
		KeyDescription: "Control and c",
	}
	NtuiEscKey = NtuiKey{
		Key:			tcell.KeyEsc,
		KeyLabel:		"Backspace",
		KeyDescription: "Backspace",
	}
	NtuiTabKey = NtuiKey{
		Key:			tcell.KeyTAB,
		KeyLabel:		"Tab Key",
		KeyDescription: "Tab Key",
	}
	NtuiCtrlRKey = NtuiKey{
		Key:			tcell.KeyCtrlR,
		KeyLabel:		"Ctrl + r",
		KeyDescription: "Control and r",
	}
	NtuiCtrlVKey = NtuiKey{
		Key:			tcell.KeyCtrlV,
		KeyLabel:		"Ctrl + v",
		KeyDescription: "Control and v",
	}
	NtuiRuneKey = NtuiKey{
		Key:			tcell.KeyRune,
		KeyLabel:		"rune keys",
		KeyDescription: "Rune keys",
	}
	NtuiCtrlTKey = NtuiKey{
		Key:			tcell.KeyCtrlT,
		KeyLabel:		"ctrl + t",
		KeyDescription: "Control and T",
	}
	NtuiCtrlQKey = NtuiKey{
		Key:			tcell.KeyCtrlQ,
		KeyLabel:		"ctrl + q",
		KeyDescription: "Control and Q",
	}
	NtuiCtrlSKey = NtuiKey{
		Key:			tcell.KeyCtrlS,
		KeyLabel:		"ctrl + s",
		KeyDescription: "Control and S",
	}
)

var NtuiKeyBindings = []NtuiKey{
	NtuiExitKey,
}