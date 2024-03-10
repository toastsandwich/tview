package form

import (
	"github.com/toastsandwich/tview/config"
	"github.com/toastsandwich/tview/model"
	"github.com/toastsandwich/tview/utils"
)

func init() {
	contact := model.Contact{}
	config.AddContact.AddInputField("First Name", "", 20, nil, func(fn string) {
		contact.Firstname = fn
	})
	config.AddContact.AddInputField("Last Name", "", 20, nil, func(ln string) {
		contact.Lastname = ln
	})
	config.AddContact.AddInputField("Email", "", 20, nil, func(e string) {
		contact.Email = e
	})
	config.AddContact.AddInputField("Phone Number", "", 10, nil, func(pn string) {
		contact.Phonenumber = pn
	})
	config.AddContact.AddDropDown("State", utils.States, 0, func(st string, index int) {
		contact.State = st
	})
	config.AddContact.AddCheckbox("Business", false, func(b bool) {
		contact.Business = b
	})
	config.AddContact.AddButton("Save", func() {
		model.Contacts = append(model.Contacts, contact)
		config.Pages.SwitchToPage("Menu")
	})
}
