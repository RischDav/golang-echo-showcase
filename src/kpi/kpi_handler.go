package kpi

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"fmt"
)

type Handler struct {
	Service *Service
}

func (h *Handler) GetKPI(c echo.Context) error {
	kpiName := c.Param("name")
	
	value, err := h.Service.GetKPI(c.Request().Context(), kpiName)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	
	return c.String(http.StatusOK, fmt.Sprintf("KPI %s: %d", kpiName, value))
}

func (h *Handler) SetKPI(c echo.Context) error {
	kpiName := c.QueryParam("name")
	kpiTypeStr := c.QueryParam("type")
	valueStr := c.QueryParam("value")
	
	err := h.Service.SetKPI(c.Request().Context(), kpiName, kpiTypeStr, valueStr)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	
	return c.NoContent(http.StatusOK)
}

func (h *Handler) GetAllKPI(c echo.Context) error {
	value, err := h.Service.GetAllKPIs(c.Request().Context())
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}
	
	return c.String(http.StatusOK, fmt.Println("%s", value))
}