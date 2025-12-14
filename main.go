package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
    "golang-echo-showcase/shared"
	"golang-echo-showcase/src/user"                
    "golang-echo-showcase/src/user/sqlc/output"
	"time"
    "log"
	_ "embed"
)

var ddl string 

var startTime time.Time

func init() {
	startTime = time.Now()
}

func main() {

	database, err := shared.NewDatabase("./shared/user.db")
	if err != nil {
		log.Fatal("Datenbankfehler:", err)
	}

	queries := sqlc.New(database.DB) 

    userHandler := &user.Handler{
        Queries: queries,
    }

	defer database.CloseDatabase()

    shared.InitializeKPIs()

    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    DefineUserRoutes(e, userHandler)

    e.Logger.Fatal(e.Start(":3333"))
}