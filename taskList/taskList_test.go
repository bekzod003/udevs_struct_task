package main

import "testing"

var (
	testTasks = []Task{
		{
			Id:        1,
			Name:      "Hw",
			Status:    "doing",
			Priority:  1,
			CreatedAt: "6.10.2020",
			CreatedBy: "Bekzod",
			DueDate:   "6.11.2020",
		},
		{
			Id:        2,
			Name:      "watch movie",
			Status:    "to do",
			Priority:  4,
			CreatedAt: "9.10.2020",
			CreatedBy: "Bekzod",
			DueDate:   "6.12.2020",
		},
		{
			Id:        3,
			Name:      "read book",
			Status:    "done",
			Priority:  2,
			CreatedAt: "7.10.2020",
			CreatedBy: "Bekzod",
			DueDate:   "9.10.2020",
		},
		{
			Id:        4,
			Name:      "walk",
			Status:    "to do",
			Priority:  3,
			CreatedAt: "12.10.2020",
			CreatedBy: "Bekzod",
			DueDate:   "14.10.2020",
		},
	}

	testStruct = TaskList{}

	updatedTask = Task{
		// field status edited
		Id:        4,
		Name:      "walk",
		Status:    "done",
		Priority:  3,
		CreatedAt: "12.10.2020",
		CreatedBy: "Bekzod",
		DueDate:   "14.10.2020",
	}
)

func TestCreate(t *testing.T) {
	for _, val := range testTasks {
		createErr := testStruct.Create(val)

		if createErr != nil {
			t.Error("failed to create")
		}
	}
}

func TestGet(t *testing.T) {
	for _, val := range testStruct.Tasks {
		_, err := testStruct.Get(val.Id)
		if err != nil {
			t.Error("failed to Get")
		}
	}
}

func TestGetAll(t *testing.T) {
	_, err := testStruct.GetAll()
	if err != nil {
		t.Error("failed to getAll")
	}
}

func TestUpdate(t *testing.T) {
	err := testStruct.Update(updatedTask)

	if err != nil {
		t.Error("failed to Update")
	}
}

func TestDelete(t *testing.T) {
	err := testStruct.Delete(updatedTask)

	if err != nil {
		t.Error("failed to Delete")
	}
}
