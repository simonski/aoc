package main

type Counter struct {
	data map[string]int
}

func NewCounter() *Counter {
	data := make(map[string]int)
	m := Counter{data: data}
	return &m
}

func (m *Counter) Get(key string, defaultValue int) int {
	value, exists := m.data[key]
	if exists {
		return value
	} else {
		return defaultValue
	}
}

func (m *Counter) Keys() []string {
	klist := make([]string, 0)
	for k, _ := range m.data {
		klist = append(klist, k)
	}
	return klist
}

func (m *Counter) Put(key string, value int) {
	m.data[key] = value
}

func (m *Counter) Increment(key string) int {
	value := m.Get(key, 0)
	value++
	m.Put(key, value)
	return value
}
