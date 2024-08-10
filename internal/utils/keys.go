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
	NtuiEnterKey = NtuiKey{
		Key: 			tcell.KeyEnter,
		KeyLabel:		"enter",
		KeyDescription:	"Enter Key",
	}
	NtuiEscKey = NtuiKey{
		Key:			tcell.KeyEsc,
		KeyLabel:		"Esc",
		KeyDescription: "Esc key",
	}
	NtuiTabKey = NtuiKey{
		Key:			tcell.KeyTAB,
		KeyLabel:		"Tab",
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
	NtuiCtrlJKey = NtuiKey{
		Key:			tcell.KeyCtrlJ,
		KeyLabel:		"ctrl + j",
		KeyDescription: "Control and J",
	}
	NtuiCtrlDKey = NtuiKey{
		Key:			tcell.KeyCtrlD,
		KeyLabel:		"ctrl + d",
		KeyDescription: "Control and D",
	}
)

var NtuiKeyBindings = []NtuiKey{
	NtuiExitKey,
}