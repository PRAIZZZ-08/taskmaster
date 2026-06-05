package main

import (
	"taskmaster/todo"
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB //sql db for total cost persistence

func Welcome(c *gin.Context) {
	fmt.Println("Welcome to Taskmaster API")
}

type Task struct { //create task struct and export as json
	Title     string `json:"title"`
	FixedCost int    `json:"fixed_cost"`
}

func (t Task) Check() string {
	if len(t.Title) > 100 {
		return "Error: Title too long"
	}

	if t.Title == "" {
		return "Invalid: Missing Title"
	}
	return fmt.Sprintf("Audited: %s with cost %d", t.Title, t.FixedCost)
}

func BulkAudit(c *gin.Context) {
	var tasks []Task
	if err := c.ShouldBindJSON(&tasks); err != nil { //[ShouldBindJSON]: Gin sees the JSON array and fills our slice
		c.JSON(400, gin.H{"error": "Invalid request format"})
		return
	}

	// 1. DATABASE FETCH: Get previous total
	var previousTotal int
	_ = db.QueryRow("SELECT total_sum FROM stats WHERE id = 1").Scan(&previousTotal)

	// 2. CONCURRENCY: Audit and Log simultaneously
	var wg sync.WaitGroup
	results := make(chan string, len(tasks))
	batchTotal := 0

	for _, t := range tasks {
		batchTotal += t.FixedCost
		wg.Add(1)
		// [go]: Launching worker
		go func(item Task) {
			defer wg.Done()

			// TASK: The Activity Logger (DB Exec inside Goroutine)
			// [Exec]: For SQL that doesn't return data
			_, _ = db.Exec("INSERT INTO audit_logs (title, status) VALUES (?, ?)", item.Title, "COMPLETED")

			results <- item.Check()
		}(t)
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

	router := gin.Default() //initialize server

	router.GET("/", Welcome)
	router.POST("/bulk-audit", BulkAudit)

	router.Run(":8080")
}
