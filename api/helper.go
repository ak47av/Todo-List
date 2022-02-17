package api

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"
	"todo_CLI/JSON"
	"todo_CLI/tasks"
)

func GetTasks(c *gin.Context) {
	tasks := JSON.ReadJSON()
	c.IndentedJSON(http.StatusOK, tasks)
}

func PostTask(c *gin.Context) {
	var newTask tasks.JsonTask

	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	fmt.Println(newTask)
	data := newTask.Data
	priority := newTask.Priority

	newList := JSON.JSONtoTask()
	newList.Insert(data, priority)
	newList.ToJSONFile()
	
	listJSON := JSON.ReadJSON()
	c.IndentedJSON(http.StatusCreated, listJSON)
}

func DeleteTaskByID(c *gin.Context) {
	id := c.Param("id")
	index, _ := strconv.Atoi(id)

	newList := JSON.JSONtoTask()
	newList.Delete(index)
	newList.ToJSONFile()

	listJSON := JSON.ReadJSON()
	c.IndentedJSON(http.StatusCreated, listJSON)
}