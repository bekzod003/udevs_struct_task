package main

import (
	"errors"
	"fmt"
)

type Contact struct {
	id        int
	firstName string
	lastName  string
	phone     string
	email     string
}

type ContactactList struct {
	Contacts []Contact
}

func (cl *ContactactList) Create(cn Contact) error {
	cl.Contacts = append(cl.Contacts, cn)
	return nil
}

func (cl *ContactactList) Update(cn Contact) error {
	for i, val := range cl.Contacts {
		if val.id == cn.id {
			cl.Contacts[i] = cn
		}
	}
	return nil
}

func (cl *ContactactList) delete(cn Contact) error {
	for i, val := range cl.Contacts {
		if val.id == cn.id {
			cl.Contacts = append(cl.Contacts[:i], cl.Contacts[i+1:]...)

			return nil
		}
	}
	return errors.New("contact with such id does not exist")
}

func (cl *ContactactList) Get(id int) (Contact, error) {
	for _, val := range cl.Contacts {
		if val.id == id {
			return val, nil
		}
	}
	return Contact{}, errors.New("contact with such id does not exist")
}

func (cl *ContactactList) GetAll() []Contact {
	return cl.Contacts
}

func main() {
	conts := new(ContactactList)

	bekzod := Contact{
		125,
		"Bekzod",
		"Isayev",
		"4565465",
		"mail.ru",
	}

	ng := Contact{
		1456,
		"asdf",
		"aljs",
		"4565465",
		"gmail.com",
	}

	conts.Create(bekzod)
	conts.Create(ng)

	fmt.Println(conts.GetAll())

	bekzod.firstName = "Bek"
	conts.Update(bekzod)
	fmt.Println()
	fmt.Println(conts.GetAll())

	fmt.Println(conts.Get(1456))
}
