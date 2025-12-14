package shared

var KPIs map[string]int

func InitializeKPIs() {
    KPIs = make(map[string]int)
}

func GetKPI(name string) (int, bool) {
    value, exists := KPIs[name]
    return value, exists
}

func SetKPI(name string, value int) {
    KPIs[name] = value
}