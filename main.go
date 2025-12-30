package main

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "golang-echo-showcase/shared"
    "golang-echo-showcase/src/user"
    "golang-echo-showcase/src/kpi"                      
    "golang-echo-showcase/src/user/sqlc/output"
    "log"
    "time"
    "os"
)

var starttime time.Time

func init() {
    starttime = time.Now()
}

func main() {
    e := echo.New()
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())
    
    userHandler, userCleanup := setupUserHandler()
    kpiHandler := setupKPIHandler()
    defer userCleanup()

    setupRoutes(e, userHandler, kpiHandler)

    e.Logger.Fatal(e.Start(":3333"))

}

func setupUserHandler() (*user.Handler, func() error) {
    dbPath := os.Getenv("USER_DB_PATH")
    if dbPath == "" {
        // Fallback für lokale Entwicklung
        dbPath = "./shared/user.db" 
    }
    database, err := shared.NewDatabase(dbPath)
    if err != nil {
        log.Fatal("Datenbankfehler:", err)
    }
    queries := sqlc.New(database.DB)
    service := user.NewService(queries)
    return &user.Handler{Service: service}, database.CloseDatabase
}

func setupKPIHandler() *kpi.Handler {
    store := shared.NewKPIStore()
    service := kpi.NewService(store)
    return &kpi.Handler{
        Service: service,  
    }
}

func getUptime(c echo.Context) error {
    uptime := time.Since(starttime)
    return c.String(200, uptime.String())
}


