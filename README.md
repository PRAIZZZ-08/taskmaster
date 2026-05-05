# taskmaster
TaskMaster API v1.1.0
Overview
TaskMaster is an expressive and efficient RESTful API built in Go
. It provides a concurrent engine for auditing  and projects, ensuring "code adaptability" and "modular program construction"
.
Core Mechanics
The system utilizes Go's concurrency primitives goroutines and channels to process bulk data requests simultaneously
.
Goroutines: Launch parallel workers for every audit
.
Channels: Safely communicate results between workers and the main API handler
.
WaitGroups: Synchronize the "Foreman" to ensure all workers finish before the report is sent
.
Architecture
  [ CLIENT ] ──(POST /bulk-audit)──► [ GIN ROUTER ]
                                           │
          ┌────────────────────────────────┴────────────────────────┐
          ▼                                                         ▼
  [ SQL DATABASE ] ◄─────── [ CONCURRENCY ENGINE ] ────────► [ JSON RESPONSE ]
  (Persistent Stats)        (Goroutine Workers)             (Final Report)
Features
Bulk Validation: Send arrays of Projects for high-speed processing
.
Data Persistence: Tracks cumulative grand totals in a relational database using the database/sql package
.
Security: Uses prepared statements and placeholders to mitigate SQL injection risks
.
Standardized Output: All data is marshalled into clean, lowercase JSON for web compatibility
.
Getting Started
Prerequisites
Go 1.25.0+
.
SQLite3 drivers.
Installation
go mod tidy # Installs dependencies like Gin [11]
go build    # Compiles into machine code [2]
Running Tests
Automated unit tests ensure the stability of the audit logic:
go test -v  # Runs all _test.go files in the module [1, 11]