package kpi

import (
	"golang-echo-showcase/shared"
	"net/http"
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
)

type KPIType string

const (

	// Wird pro Stunde aggregiert (eingegebene Zeit + jetzige Zeit / Stunden gesamt)
    KPICount KPIType = "COUNT"
	
	// Speichert einfach ein wert
    KPISave KPIType = "SAVE"
)

func GetKPI(c echo.Context) error {
	kpiName := c.Param("name")
	kpiValue, exists := shared.KPIs[kpiName]
	if !exists {
		return c.String(http.StatusNotFound, "KPI nicht gefunden: "+kpiName)
	}
	return c.String(http.StatusOK, fmt.Sprintf("KPI %s: %d", kpiName, kpiValue))
}

func WriteKPI(c echo.Context) error {
	kpiName := c.Param("name") 
    kpiTypeStr := c.Param("type")
    valueStr := c.Param("value")
    valueInt, err := strconv.Atoi(valueStr)
	if err != nil {
		return c.String(http.StatusBadRequest, "Ungültiger Wert für Value")
	}

	switch kpiTypeStr {
		case "KPICount":
			kpiName += "_COUNT"
		case "KPISave":
			kpiName += "_SAVE"
		default:
			return fmt.Errorf("Falscher KPI-type")	
	}
	shared.KPIs[kpiName] = valueInt
	return nil
}