package kpi

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"golang-echo-showcase/shared/response"
)

type Handler struct {
	Service *Service
}

func successResponse(data interface{}, message string, start time.Time) *response.ApiResponse {
	return &response.ApiResponse{
		Success:        true,
		ProcessingTime: time.Since(start).String(),
		Data:           data,
		Message:        message,
	}
}

func errorResponse(message string, start time.Time) *response.ApiResponse {
	return &response.ApiResponse{
		Success:        false,
		ProcessingTime: time.Since(start).String(),
		Error:          message,
	}
}

func (h *Handler) GetKPI(c echo.Context) error {
	start := time.Now()

	kpiName := c.Param("name")

	value, err := h.Service.GetKPI(c.Request().Context(), kpiName)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			errorResponse(err.Error(), start),
		)
	}

	data := map[string]interface{}{
		"name":  kpiName,
		"value": value,
	}

	return c.JSON(
		http.StatusOK,
		successResponse(data, "KPI erfolgreich abgerufen", start),
	)
}

func (h *Handler) SetKPI(c echo.Context) error {
	start := time.Now()

	kpiName := c.QueryParam("name")
	kpiTypeStr := c.QueryParam("type")
	valueStr := c.QueryParam("value")

	err := h.Service.SetKPI(c.Request().Context(), kpiName, kpiTypeStr, valueStr)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			errorResponse(err.Error(), start),
		)
	}

	return c.JSON(
		http.StatusOK,
		successResponse(nil, "KPI erfolgreich gesetzt", start),
	)
}

func (h *Handler) GetAllKPI(c echo.Context) error {
	start := time.Now()

	value, err := h.Service.GetAllKPIs(c.Request().Context())
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			errorResponse(err.Error(), start),
		)
	}

	return c.JSON(
		http.StatusOK,
		successResponse(value, "Alle KPIs erfolgreich abgerufen", start),
	)
}
