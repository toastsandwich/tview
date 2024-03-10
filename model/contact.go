package model

type Contact struct {
	Firstname   string
	Lastname    string
	Email       string
	Phonenumber string
	State       string
	Business    bool
}

var Contacts []Contact
