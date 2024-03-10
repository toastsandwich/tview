package main

import (
	"github.com/toastsandwich/tview/config"
)

func main() {
	if err := config.Application.SetRoot(config.Pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
