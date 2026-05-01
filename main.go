package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Task struct { //create task struct and export as json
Description string `json:"description"`
Estimate int `json:"estimate"`
}

type Project struct { //create project struct and export as json
Title string `json:"title"`
FixedCost int `json:"fixed_cost"`
}

func Welcome(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Welcome to Taskmaster API")
}
//w: gets response from api | r: http request from api visit
func HealthCheck(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "Taskmaster System: Functional")//print response from func and msg
}

func GetTask(w http.ResponseWriter, r *http.Request) {
t1 := Task{Description: "Fix Bug", Estimate: 5}//struct of data
w.Header().Set("Content-Type", "application/json")  //encode response as json
json.NewEncoder(w).Encode(t1) //tell browser its json data
}

func main() {
http.HandleFunc("/", Welcome)
http.HandleFunc("/health", HealthCheck) //route handler to health
http.HandleFunc("/task", GetTask)//route handler to task

fmt.Println("Taskmaster Server starting on :8080...")

http.ListenAndServe(":8080", nil)
}
