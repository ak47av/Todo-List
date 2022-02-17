package JSON

import (
	"todo_CLI/tasks"
	"log"
	"io/ioutil"
	"encoding/json"
)	
func JSONtoTask() *tasks.List {
	taskObjects := ReadJSON()
	newList := &tasks.List{}
	for _, result := range taskObjects {
		err := newList.Insert(result.Data, result.Priority)
		if err != nil {
			log.Fatal(err)
		}
	}
	return newList
}

func ReadJSON() []tasks.JsonTask{
	byteValue, err := ioutil.ReadFile("tasks.json")
	if err != nil {
		log.Fatal(err)
	}
	results := []tasks.JsonTask{}
	json.Unmarshal(byteValue, &results)
	return results
}