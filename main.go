package main

import (
	"fmt"
	"sync"
	"github.com/gin-gonic/gin"
	//"database/sql"
)

func Welcome(c *gin.Context) {
fmt.Println("Welcome to Taskmaster API")
}

type Project struct { //create project struct and export as json
Title string `json:"title"`
FixedCost int `json:"fixed_cost"`
}

func (p Project) Check() string {
if p.Title == "" {
return "Invalid: Missing Title"
}
return fmt.Sprintf("Audited: %s with cost %d", p.Title, p.FixedCost)
}

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

func main() {

router := gin.Default()

router.GET("/", Welcome)
router.POST("/bulk-audit", BulkAudit)

router.Run(":8080")
}
