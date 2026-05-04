package main

import (
	"github.com/gin-gonic/gin"
)

type Task struct { //create task struct and export as json
Description string `json:"description"`
Estimate int `json:"estimate"`
}

type Project struct { //create project struct and export as json
Title string `json:"title"`
FixedCost int `json:"fixed_cost"`
}

func CreateTask(c *gin.Context) {
	var newTask Task //The empty "clipboard" waiting for data
	// [&newTask]: We pass the MEMORY ADDRESS so Gin can write directly to our variable.
	if err := c.ShouldBindJSON(&newTask); err != nil {
		// : StatusBadRequest. The client sent bad data.
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	message := "Received: " + newTask.Description//store the new data in the Struct 
	c.JSON(200, gin.H{"message": message, "status": "accepted"})//return 201 status & message
}

func CreateProject(c *gin.Context) {
var newProject Project//var to store POST request
if err := c.ShouldBindJSON(&newProject); err != nil {//if there is an error binding
c.JSON(400, gin.H{"error": err.Error()})//return 400 status and error state
return
}
message := newProject.Title
c.JSON(201, gin.H{"Project": message, "status": "has been added to the system"})//return 201 status and message
}

func main() {
router := gin.Default()

router.POST("/task", CreateTask)
router.POST("/project", CreateProject)

router.Run()
}
