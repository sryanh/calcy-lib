package calc

type Calculator interface{ Calculate(a, b int) int }

type Addition struct {
	calculator Calculator
}

func (Addition) Calculate(a, b int) int {
	return a + b
}
