package main

import (
	"errors"
	"fmt"
)

type Task struct {
	Id        int
	Name      string
	Status    string
	Priority  int
	CreatedAt string
	CreatedBy string
	DueDate   string
}

type TaskList struct {
	Tasks []Task
}

func (tl *TaskList) Create(t Task) error {
	tl.Tasks = append(tl.Tasks, t)
	return nil
}

func (tl *TaskList) Update(t Task) error {
	for i, val := range tl.Tasks {
		if val.Id == t.Id {
			tl.Tasks[i] = t
			return nil
		}
	}
	return errors.New("no task with such id")
}

func (tl *TaskList) Get(id int) (Task, error) {
	for i, val := range tl.Tasks {
		if val.Id == id {
			return tl.Tasks[i], nil
		}
	}
	return Task{}, errors.New("no task with such id")
}

func (tl *TaskList) GetAll() []Task {
	return tl.Tasks
}

func (tl *TaskList) Delete(t Task) error {
	for i, val := range tl.Tasks {
		if val.Id == t.Id {
			tl.Tasks = append(tl.Tasks[:i], tl.Tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("no task with such id")
}

func main() {
	tl := new(TaskList)

	hw := Task{
		1,
		"Do hw",
		"To do",
		1,
		"19.09.2021",
		"Bekzod",
		"21.09.2021",
	}
	work := Task{
		2,
		"work",
		"doing",
		2,
		"20.09.2021",
		"Bekzod",
		"22.09.2021",
	}

	tl.Create(hw)
	tl.Create(work)

	fmt.Println(tl.GetAll())

	fmt.Printf("\n\n\n")

	fmt.Println(tl.Get(2))

	fmt.Printf("\n\n\n")

	hw.Priority = 5
	tl.Update(hw)
	fmt.Println(tl.GetAll())

}
