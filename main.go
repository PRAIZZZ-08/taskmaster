package main

import (
	"fmt"
	"sync"
	"github.com/gin-gonic/gin"
	"database/sql"
	"log"
)

var db *sql.DB//sql db for total cost persistence

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
	
	// 1. DATABASE FETCH: Get previous total
	var previousTotal int
	_ = db.QueryRow("SELECT total_sum FROM stats WHERE id = 1").Scan(&previousTotal)

	// 2. CONCURRENCY: Audit and Log simultaneously
	var wg sync.WaitGroup
	results := make(chan string, len(projects))
	batchTotal := 0

	for _, p := range projects {
		batchTotal += p.FixedCost
		wg.Add(1)
		// [go]: Launching worker
		go func(item Project) {
			defer wg.Done()
			
			// TASK: The Activity Logger (DB Exec inside Goroutine)
			// [Exec]: For SQL that doesn't return data 
			_, _ = db.Exec("INSERT INTO audit_logs (title, status) VALUES (?, ?)", item.Title, "COMPLETED")
			
			results <- item.Check()
		}(p)
	}

	wg.Wait()
	close(results)

	// 3. DATABASE UPDATE: Save the new grand total
	newGrandTotal := previousTotal + batchTotal
	_, _ = db.Exec("UPDATE stats SET total_sum = ? WHERE id = 1", newGrandTotal)

	// 4. COLLECT REPORTS
	var finalReport []string
	for r := range results {
		finalReport = append(finalReport, r)
	}

	// [c.JSON]: Unified response
	c.JSON(200, gin.H{
		"batch_total":     batchTotal,
		"new_grand_total": newGrandTotal,
		"reports":         finalReport,
	})
}

func main() {

var err error
// [sql.Open]: Initialize database handle [6]
db, err = sql.Open("sqlite3", "taskmaster.db")
if err != nil {
	log.Fatal(err)
}
	
router := gin.Default()//initialize server

router.GET("/", Welcome)
router.POST("/bulk-audit", BulkAudit)

router.Run(":8080")
}
