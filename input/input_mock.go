package input

type MockFloatInputGetter struct {
	Responses []float64
	Index     int
}

func (m *MockFloatInputGetter) GetFloatInput(prompt string) float64 {
	value := m.Responses[m.Index]
	m.Index++
	return value
}

type MockIntInputGetter struct {
	Responses []int
	Index     int
}

func (m *MockIntInputGetter) GetIntInput(prompt string) int {
	value := m.Responses[m.Index]
	m.Index++
	return value
}
