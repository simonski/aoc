package utils

type Counter struct {
	Data map[string]int
}

func NewCounter() *Counter {
	data := make(map[string]int)
	m := Counter{Data: data}
	return &m
}

func (m *Counter) Get(key string, defaultValue int) int {
	value, exists := m.Data[key]
	if exists {
		return value
	} else {
		return defaultValue
	}
}

func (m *Counter) Keys() []string {
	klist := make([]string, 0)
	for k, _ := range m.Data {
		klist = append(klist, k)
	}
	return klist
}

func (m *Counter) Put(key string, value int) {
	m.Data[key] = value
}

func (m *Counter) Increment(key string) int {
	value := m.Get(key, 0)
	value++
	m.Put(key, value)
	return value
}
