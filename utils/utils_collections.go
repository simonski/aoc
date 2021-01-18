package utils

type IntMap struct {
	data map[int]int
}

func NewIntMap() *IntMap {
	data := make(map[int]int)
	m := IntMap{data: data}
	return &m
}

func (m *IntMap) Get(key int, defaultValue int) int {
	value, exists := m.data[key]
	if exists {
		return value
	} else {
		return defaultValue
	}
}

func (m *IntMap) Put(key int, value int) {
	m.data[key] = value
}

func (m *IntMap) Increment(key int) int {
	value := m.Get(key, 0)
	value++
	m.Put(key, value)
	return value
}
