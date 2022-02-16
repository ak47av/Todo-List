package main

import (
	"fmt"
)

type Task struct {
	priority int
	data     string
	next     *Task
}

type List struct {
	head *Task
}

func (L *List) Insert(data string, priority int) {

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
			temp.next = curr.next
			curr.next = temp
		}
	}
}

func (L *List) Display() {
	task := L.head
	for task != nil {
		fmt.Printf("%+v ", task.data)
		task = task.next
	}
	fmt.Println()
}

func main() {
	list := &List{}
	list.Insert("Get an Apple", 4)
	list.Insert("Get a Banana", 2)
	list.Insert("Get an Orange", 1)
	list.Insert("Get a Kiwi", 3)
	list.Display()

}
