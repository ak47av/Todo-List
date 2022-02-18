package main

import (
	"todo_CLI/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/tasks", api.GetTasks)
	router.POST("/tasks", api.PostTask)
	router.DELETE("/tasks/:id", api.DeleteTaskByID)
	router.Run(":8080")
}
