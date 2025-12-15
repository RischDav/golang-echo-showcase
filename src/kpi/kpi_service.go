package kpi

import (
    "context"
    "golang-echo-showcase/shared"
	"fmt"
	"strconv"
)

type Service struct {
    store *shared.KPIStore  
}

func NewService(store *shared.KPIStore) *Service {
    return &Service{store: store}  
}

func (s *Service) GetKPI(ctx context.Context, name string) (int, error) {
    value, exists := s.store.GetKPI(name)  
    if !exists {
        return 0, fmt.Errorf("KPI %s nicht gefunden", name)
    }
    return value, nil
}

func (s *Service) SetKPI(ctx context.Context, name string, kpiTypeStr string, valueStr string) error {
    value, err := strconv.Atoi(valueStr)
    if err != nil {
        return fmt.Errorf("ungültiger Wert: %w", err)
    }
    
    switch kpiTypeStr {
    case "KPICount":
        name += "_COUNT"
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