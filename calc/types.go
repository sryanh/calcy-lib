package calc

type Calculator interface{ Calculate(a, b int) int }

type Addition struct {
	calculator Calculator
}
type Subtraction struct {
	calculator Calculator
}
type Multiplication struct {
	calculator Calculator
}
type Division struct {
	calculator Calculator
}
