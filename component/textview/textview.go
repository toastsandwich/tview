package textview

import (
	"github.com/gdamore/tcell/v2"
	"github.com/toastsandwich/tview/config"
)

func init() {
	config.Menu.SetTextColor(tcell.ColorYellow).SetText("(a) --> add contact\n(q) --> quit")
}
