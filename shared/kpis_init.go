package shared


type KPIStore struct {
    kpis  map[string]int
}

func NewKPIStore() *KPIStore {
    return &KPIStore{
        kpis: make(map[string]int),
    }
}

func (s *KPIStore) GetKPI(name string) (int, bool) {
    value, exists := s.kpis[name]
    return value, exists
}

func (s *KPIStore) SetKPI(name string, value int) {
    s.kpis[name] = value
}

func (s *KPIStore) GetAllKPIs() map[string]int {
    return s.kpis
}