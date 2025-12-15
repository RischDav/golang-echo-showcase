package main

import (
    "github.com/labstack/echo/v4"
    "golang-echo-showcase/src/user"   
    "golang-echo-showcase/src/kpi"   
)

func setupRoutes(e *echo.Echo, userHandler *user.Handler, kpiHandler *kpi.Handler) {
    e.POST("/users", userHandler.SaveUser)
    e.GET("/users", userHandler.GetUser)
    e.PUT("/users/:id", userHandler.UpdateUserFirstname)
    e.DELETE("/users/:id", userHandler.DeleteUser)
    e.GET("/allkpis", kpiHandler.GetAllKPI)
    e.GET("/kpi/:name", kpiHandler.GetKPI)
    e.POST("/kpi", kpiHandler.SetKPI)
    
}