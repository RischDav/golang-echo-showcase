package main

import (
    "github.com/labstack/echo/v4"
    "golang-echo-showcase/src/user"   
    "golang-echo-showcase/src/kpi"   
)

func DefineUserRoutes(e *echo.Echo, h *user.Handler) {
    e.POST("/users/:firstname/:lastname", h.SaveUser)
    e.GET("/users/:id", h.GetUser)
    e.PUT("/users/:id", h.UpdateUser)
    e.DELETE("/users/:id", h.DeleteUser)
    e.POST("/kpis/:name/:value/:type", kpi.WriteKPI)
    e.GET("/kpis/:name", kpi.GetKPI)
}