package main

import "testing"

var (
	contacts = []Contact{
		{
			id:        1,
			firstName: "Bekzod",
			lastName:  "Isayev",
			phone:     "+99897123456",
			email:     "gmail.com",
		},
		{
			id:        2,
			firstName: "Tom",
			lastName:  "Johns",
			phone:     "+138136843",
			email:     "gmail.us",
		},
		{
			id:        3,
			firstName: "Helen",
			lastName:  "Paker",
			phone:     "+13813253243",
			email:     "gmail.us.uk",
		},
		{
			id:        4,
			firstName: "Anvar",
			lastName:  "Jorakulov",
			phone:     "+9989793546",
			email:     "mail.uz",
		},
	}

	updAnvar = Contact{
		id:        4,
		firstName: "Anvar",
		lastName:  "Sonovich",
		phone:     "+9989793546",
		email:     "mail.uz",
	}

	cList = ContactactList{}
)

func TestCreate(t *testing.T) {
	for _, val := range contacts {
		err := cList.Create(val)
		if err != nil {
			t.Error("failed to Create")
		}
	}
}

func TestGet(t *testing.T) {
	for _, val := range cList.Contacts {
		_, err := cList.Get(val.id)
		if err != nil {
			t.Error("failed to Get")
		}
	}
}

func TestGetAll(t *testing.T) {
	_, err := cList.GetAll()
	if err != nil {
		t.Error("failed to GetAll")
	}
}

func TestUpdate(t *testing.T) {
	err := cList.Update(updAnvar)
	if err != nil {
		t.Error("failed to Update")
	}
}

func TestDelete(t *testing.T) {
	err := cList.Delete(updAnvar)
	if err != nil {
		t.Error("failed to Delete")
	}
}
