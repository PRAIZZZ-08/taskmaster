package main

import (
	"fmt"
	"sync"
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

func (t Task) Check() string {
if t.Description == "" {
return "Invalid: Missing Description"
}
return fmt.Sprintf("Audited %s with cost %d", t.Description, t.Estimate)
}

func (p Project) Check() string {
if p.Title == "" {
return "Invalid: Missing Title"
}
return fmt.Sprintf("Audited: %s with cost %d", p.Title, p.FixedCost)
}
a
func BulkAudit(c *gin.Context) {
	var projects []Project
	if err := c.ShouldBindJSON(&projects); err != nil {//[ShouldBindJSON]: Gin sees the JSON array and fills our slice
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	var wg sync.WaitGroup
	// [results]: A channel to catch reports from our parallel workers
	results := make(chan string, len(projects))
	
	totalCost :=  0 //running tally for cost

	for _, p := range projects {
	totalCost += p.FixedCost
		wg.Add(1)
		// [go]: Launching a Goroutine for every task in the list
		go func(item Project) {
			defer wg.Done()
			results <- item.Check()
		}(p)
	}

	// Wait and close in the background
	wg.Wait()
	close(results)

	// [finalReport]: Extracting results from the channel into a slice
	var finalReport []string
	for r := range results {
		finalReport = append(finalReport, r)
	}

	// [JSON]: Sending the full report back to the user
	c.JSON(200, gin.H{
		"total_cost": totalCost,
		"processed_count": len(projects),
		"reports":         finalReport,
	})
}


func CreateTask(c *gin.Context) {
var newTask Task//var to store POST request
if err := c.ShouldBindJSON(&newTask); err != nil {//if there is an error binding
c.JSON(400, gin.H{"error": err.Error()})//return 400 status and error state
return
}
message := newTask.Description
c.JSON(201, gin.H{"Description": message, "status": "has been added to the system"})//return 201 status and message
}

func main() {

router := gin.Default()

router.POST("/task", CreateTask)
router.POST("/bulk-audit", BulkAudit)

router.Run(":8080")
}
