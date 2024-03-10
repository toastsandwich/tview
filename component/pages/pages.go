package pages

import (
	"github.com/toastsandwich/tview/config"
)

func init() {
	config.Pages.AddPage("Menu", config.Menu, true, true)
	config.Pages.AddPage("Add Contact", config.AddContact, true, false)
}
