package application

import (
	"github.com/gdamore/tcell/v2"
	"github.com/toastsandwich/tview/config"
)

func init() {
	config.Application.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			config.Application.Stop()
		} else if event.Rune() == 'a' {
			config.Pages.SwitchToPage("Add Contact")
		}
		return event
	})
}
