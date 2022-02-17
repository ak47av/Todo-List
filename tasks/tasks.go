package tasks

import (
	"encoding/json"
	"todo_CLI/errors"
	"fmt"
	"io/ioutil"
	// "log"
	// "os"
	"reflect"
)

type Task struct {
	priority int
	data     string
	next     *Task
}

type List struct {
	head *Task
}

type JsonTask struct {
	Priority int `json:"priority"`
	Data string `json:"task"`
}

func (L *List) Insert(data string, priority int) (error) {

	if reflect.TypeOf(data).String() != "string" {
		return errors.New("Enter a String")
	} else if priority <= 0 {
		return errors.New("Assign a positive number as a priority")
	}

	temp := &Task{
		priority: priority,
		data:     data,
	}

	if L.head == nil || priority < L.head.priority {
		temp.next = L.head
		L.head = temp
	} else {
		curr := L.head
		for curr.next != nil && curr.next.priority < priority {
			curr = curr.next
		}
		temp.next = curr.next
		curr.next = temp	
	}
	return nil
}

func (L *List) Delete(index int) (error) {
	if(index == 1){
		L.head = L.head.next
	} else if index<=0{
		return errors.New("Enter a positive index to")
	} else {
		curr := L.head
		for i:=2; i< index; i++ {
			if(curr.next != nil){
				curr = curr.next
			}
		}
		curr.next = curr.next.next
	}
	return nil
}

func (L *List) Display() {
	task := L.head
	index := 1
	for task != nil {
		fmt.Printf("%+v : %+v --%+v \n", task.data, task.priority, index)
		task = task.next
		index++
	}
	fmt.Println()
}

func (L *List) ToJSONFile() {
	task := L.head
	newList := []JsonTask{}
	datum := JsonTask{}
	for task != nil {
		datum = JsonTask {
			Priority: task.priority,
			Data: task.data,
		}
		newList = append(newList, datum)
		task = task.next
	}

	file, err := json.MarshalIndent(newList, "", "  ")
	if err != nil {
		panic(err)
	}
	_ = ioutil.WriteFile("tasks.json", file, 0644)
}