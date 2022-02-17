package main

import (
	//"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"todo_CLI/tasks"
	//"os"
)

func readJSON() (*tasks.List){
	byteValue, err := ioutil.ReadFile("tasks.json")
	if err!=nil {
		log.Fatal(err)
	}
	results := []tasks.JsonTask{}
	json.Unmarshal(byteValue, &results)

	newList := &tasks.List{}
	for _, result := range results {
		err = newList.Insert(result.Data, result.Priority)
		if(err != nil) {
			log.Fatal(err)
		}
	}
	return newList
}

func main() {
	newList := readJSON()
	newList.Display()
}
