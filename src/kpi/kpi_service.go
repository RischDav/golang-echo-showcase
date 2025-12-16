package kpi

import (
    "context"
    "golang-echo-showcase/shared"
	"fmt"
	"strconv"
    "time"
    "strings"
)

type Service struct {
    store *shared.KPIStore  
}

func NewService(store *shared.KPIStore) *Service {
    return &Service{store: store}  
}

func (s *Service) GetKPI(ctx context.Context, name string) (float64, error) {
    value, exists := s.store.GetKPI(name)
    if !exists {
        return 0, fmt.Errorf("KPI %s nicht gefunden", name)
    }

    if strings.Contains(name, "_COUNT_") {
        parts := strings.Split(name, "_COUNT_")
        if len(parts) != 2 {
            return 0, fmt.Errorf("ungültiger COUNT-KPI-Name: %s", name)
        }

        startHour, err := strconv.Atoi(parts[1])
        if err != nil {
            return 0, fmt.Errorf("ungültige Stunde im KPI-Name: %w", err)
        }

        nowHour := time.Now().Hour()
        if nowHour < startHour {
            return 0, fmt.Errorf("aktuelle Stunde ist kleiner als Startstunde")
        }

        sum := 0
        hours := 0
        for h := startHour; h <= nowHour; h++ {
            sum += h 
            hours++
        }

        average := float64(sum) / float64(hours)
        return average, nil
    }

    // Normale KPIs
    return float64(value), nil
}


func (s *Service) SetKPI(ctx context.Context, name string, kpiTypeStr string, valueStr string) error {
    value, err := strconv.Atoi(valueStr)
    if err != nil {
        return fmt.Errorf("ungültiger Wert: %w", err)
    }
    
    switch kpiTypeStr {
    case "KPICount":
        hour := time.Now().Hour() 
        name += "_COUNT_" + strconv.Itoa(hour)
    case "KPISave":
        name += "_SAVE"
    default:
        return fmt.Errorf("falscher KPI-Type: %s", kpiTypeStr)
    }
    
    s.store.SetKPI(name, value)
    return nil
}

func (s *Service) GetAllKPIs(ctx context.Context) (map[string]int, error) {
    allKPIs := s.store.GetAllKPIs()
    return allKPIs, nil
}