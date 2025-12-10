# Go Backend API

A simple backend API built with **Golang**, **Gin**, and **PostgreSQL**.

## Requirements
- **Go version:** `go1.25.5 windows/amd64`
- Run `go version` to verify:
## Database Setup

The project connects to PostgreSQL using the `InitDB()` function:

```go
func InitDB() {
    dsn := "host=localhost user=postgres password=yourpassword dbname=your_db_name port=yourport sslmode=disable"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    DB = db
    fmt.Println("Database connected")

    // Auto-migrate Product table (optional)
    // db.AutoMigrate(&models.Product{})
}
> **Important:** Make sure to update the DSN values (`host`, `user`, `password`, `dbname`, `port`) before running the project.

