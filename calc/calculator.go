package calc

func (Addition) Calculate(a, b int) int {
	return a + b
}

func (Subtraction) Calculate(a, b int) int {
	return a - b
}

func (Multiplication) Calculate(a, b int) int {
	return a * b
}

func (Division) Calculate(a, b int) int {
	return a / b
}
